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
	// TODO: Add command
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
	"current_op_time":   c.Str("currentOpTime", s.Optional),
	"opid":              c.Int("opid", s.Optional),
	"secs_running":      c.Int("secs_running", s.Optional),
	"microsecs_running": c.Int("microsecs_running", s.Optional),
	"op":                c.Str("op", s.Optional),
	"ns":                c.Str("ns", s.Optional),
	// TODO: Add command
	// TODO: Add originating command
	"plan_summary":      c.Str("planSummary", s.Optional),
	"num_yields":        c.Int("numYields", s.Optional),
	"waiting_for_locks": c.Bool("waitingForLock", s.Optional),
}

var schemaUpdate = s.Schema{}
var schemaInsert = s.Schema{}
var schemaQuery = s.Schema{}
var schemaRemove = s.Schema{}
var schemaKillCursors = s.Schema{}
