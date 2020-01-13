package currentops

import (
	"fmt"

	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/mongodb"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

// Connection info for redis sentinel
var redisSentinels = []string{"redis-1:16379", "redis-2:16379", "redis-3:16379"}

// We store full queries in cronjob redis
const redisSentinelMasterName = "production-redis"
const redisFullQueryDBNumber = 31

// Redis Sentinel client
var redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
	MasterName:    redisSentinelMasterName,
	SentinelAddrs: redisSentinels,
	DB:            redisFullQueryDBNumber,
})

var logger = logp.NewLogger("mongodb.dbstats")

// init registers the MetricSet with the central registry as soon as the program
// starts. The New function will be called later to instantiate an instance of
// the MetricSet for each host defined in the module's configuration. After the
// MetricSet has been created then Fetch will begin to be called periodically.
func init() {
	mb.Registry.MustAddMetricSet("mongodb", "currentops", New,
		mb.WithHostParser(mongodb.ParseURL),
		mb.DefaultMetricSet(),
	)
}

// MetricSet holds any configuration or state information. It must implement
// the mb.MetricSet interface. And this is best achieved by embedding
// mb.BaseMetricSet because it implements all of the required mb.MetricSet
// interface methods except for Fetch.
type MetricSet struct {
	*mongodb.MetricSet
}

// New creates a new instance of the MetricSet. New is responsible for unpacking
// any MetricSet specific configuration options if there are any.
func New(base mb.BaseMetricSet) (mb.MetricSet, error) {
	ms, err := mongodb.NewMetricSet(base)
	if err != nil {
		return nil, err
	}
	return &MetricSet{ms}, nil
}

// Fetch methods implements the data gathering and data conversion to the right
// format. It publishes the event which is then forwarded to the output. In case
// of an error set the Error field of mb.Event or simply call report.Error().
func (m *MetricSet) Fetch(reporter mb.ReporterV2) error {
	// instantiate direct connections to each of the configured Mongo hosts
	mongoSession, err := mongodb.NewDirectSession(m.DialInfo)
	if err != nil {
		return errors.Wrap(err, "error creating new Session")
	}
	defer mongoSession.Close()

	// run command "db.runCommand({currentOp: 1})" on "admin" db
	var result map[string]interface{}
	cmd := bson.D{
		{
			Name:  "currentOp",
			Value: 1,
		},
	}
	err = mongoSession.DB("admin").Run(cmd, &result)
	if err != nil {
		err = errors.Wrap(err, "failed to retrieve currentOp")
		reporter.Error(err)
		m.Logger().Error(err)
		return err
	}

	logp.Debug("mongodb", "Result: %+v", result)
	ops, found := result["inprog"]
	if !found {
		err = errors.Wrap(err, "failed to retrieve inprog ops")
		reporter.Error(err)
		m.Logger().Error(err)
		return err
	}

	logp.Debug("mongodb", "Type of ops: %T", ops)
	logp.Debug("mongodb", "Type of ops[0]: %T", ops.([]interface{})[0])
	// parse the response and transform it to a list of entries
	for _, op := range ops.([]interface{}) {
		op := op.(map[string]interface{})
		logp.Debug("mongodb", "op: %+v", op)
		opInterface, found := op["op"]
		if !found {
			err = fmt.Errorf("operation not specified")
			reporter.Error(err)
			m.Logger().Error(err)
			continue
		}

		var schema s.Schema
		operation := fmt.Sprintf("%v", opInterface)
		switch operation {
		case "none":
			schema = schemaNone
		case "update":
			schema = schemaUpdate
		case "insert":
			schema = schemaInsert
		case "query":
			schema = schemaQuery
		case "command":
			schema = schemaCommand
		case "getmore":
			schema = schemaGetMore
		case "remove":
			schema = schemaRemove
		case "killcursors":
			schema = schemaKillCursors
		default:
			err = fmt.Errorf("unknown operation: %s", operation)
			reporter.Error(err)
			m.Logger().Error(err)
			continue
		}
		entry, err := schema.Apply(op)
		if err != nil {
			err = errors.Wrap(err, "failed to apply schema")
			reporter.Error(err)
			m.Logger().Error(err)
			continue
		}
		err = appendFullMongoQuery(entry)
		if err != nil {
			err = errors.Wrap(err, "full query not found in the redis")
			reporter.Error(err)
			m.Logger().Error(err)
		}

		// report each entry as an event
		reported := reporter.Event(mb.Event{MetricSetFields: entry})
		if !reported {
			err = errors.Wrap(err, "failed reporting event")
			m.Logger().Error(err)
			reporter.Error(err)
		}
	}
	return nil
}

func appendFullMongoQuery(entry common.MapStr) error {
	queryCommand, _ := entry["command"]
	logp.Debug("MongoDB", "queryCommand: %+v", queryCommand)
	if queryCommand != nil {
		queryMap := queryCommand.(common.MapStr)
		queryComment := queryMap["comment"]

		if queryComment != nil {
			stringComment := queryComment.(string)
			logp.Debug("MongoDB", "stringComment %s", stringComment)

			fullQuery, err := getQueryFromComment(stringComment)
			if err != nil {
				return err
			}
			logp.Debug("MongoDB", "fullQuery %s", fullQuery)
			queryMap["full_query"] = fullQuery
		}
	}
	return nil
}

func getQueryFromComment(queryComment string) (string, error) {
	return valueFromRedis(queryComment)
}

func valueFromRedis(key string) (string, error) {
	return redisClient.Get(key).Result()
}
