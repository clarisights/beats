package currentops

import (
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/mb"
	"github.com/elastic/beats/metricbeat/module/mongodb"
	"github.com/pkg/errors"
	"gopkg.in/mgo.v2/bson"
)

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

	ops, found := result["inprog"]
	if !found {
		err = errors.Wrap(err, "failed to retrieve inprog ops")
		reporter.Error(err)
		m.Logger().Error(err)
		return err
	}

	// parse the response and transform it to a list of entries
	for _, op := range ops.([]map[string]interface{}) {
		entry, err := schema.Apply(op)
		if err != nil {
			err = errors.Wrap(err, "failed to apply schema")
			reporter.Error(err)
			m.Logger().Error(err)
			return err
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
