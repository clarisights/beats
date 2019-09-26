package currentops

import (
	"github.com/elastic/beats/libbeat/common"
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

func eventMapping(key string, data common.MapStr) (common.MapStr, error) {
	return nil, nil
}

var schemaNone = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}

var schemacommand = s.Schema{
	"truncated":  c.Str("$truncated", s.Optional),
	"comment":    c.Str("comment", s.Optional),
	"db":         c.Str("$db", s.Optional),
	"get_more":   c.Int("getMore", s.Optional),
	"collection": c.Str("collection", s.Optional),
	"batch_size": c.Int("batchSize", s.Optional),
	"term":       c.Int("term", s.Optional),
	"find":       c.Str("find", s.Optional),
	// "filter": c.DictOptional(c.Dict("filter", s.Schema{
	// 	"ts": c.DictOptional(c.Dict("ts", s.Schema{
	// 		"gte": c.Str("$gte", s.Optional),
	// 	})),
	// })),
	"filter":       c.Str("filter", s.Optional),
	"tailable":     c.Bool("tailable", s.Optional),
	"oplog_replay": c.Bool("oplogReplay", s.Optional),
	"await_data":   c.Bool("awaitData", s.Optional),
	"aggregate":    c.Int("aggregate", s.Optional),
	"from_mongos":  c.Bool("fromMongos", s.Optional),
	"cursor": c.DictOptional(c.Dict("cursor", s.Schema{
		"batch_size": c.Int("batchSize", s.Optional),
	})),
}

var schemaCommand = s.Schema{
	"shard":         c.Str("shard", s.Optional),
	"host":          c.Str("host", s.Optional),
	"desc":          c.Str("desc", s.Optional),
	"connection_id": c.Int("connectionId", s.Optional),
	"client_s":      c.Str("client_s", s.Optional),
	"app_name":      c.Str("appName", s.Optional),
	// TODO: Add client metadata
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"secs_running":      c.Int("secs_running", s.Optional),
	"microsecs_running": c.Int("microsecs_running", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"command":           c.DictOptional(c.Dict("command", schemacommand)),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}

var schemaGetMore = s.Schema{
	"shard":         c.Str("shard", s.Optional),
	"host":          c.Str("host", s.Optional),
	"desc":          c.Str("desc", s.Optional),
	"connection_id": c.Int("connectionId", s.Optional),
	"client_s":      c.Str("client_s", s.Optional),
	// TODO: Add client metadata
	"active": c.Bool("active"),
	// TODO: Parse this as time
	"current_op_time":     c.Str("currentOpTime", s.Optional),
	"opid":                c.Str("opid", s.Optional),
	"secs_running":        c.Int("secs_running", s.Optional),
	"microsecs_running":   c.Int("microsecs_running", s.Optional),
	"op":                  c.Str("op", s.Optional),
	"ns":                  c.Str("ns", s.Optional),
	"command":             c.DictOptional(c.Dict("command", schemacommand)),
	"originating_command": c.DictOptional(c.Dict("originatingCommand", schemacommand)),
	"plan_summary":        c.Str("planSummary", s.Optional),
	"num_yields":          c.Int("numYields", s.Optional),
	"waiting_for_locks":   c.Bool("waitingForLock", s.Optional),
}

var schemaUpdate = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}
var schemaInsert = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}
var schemaQuery = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}
var schemaRemove = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}
var schemaKillCursors = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host", s.Optional),
	"desc":              c.Str("desc", s.Optional),
	"active":            c.Bool("active", s.Optional),
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Str("opid", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}
