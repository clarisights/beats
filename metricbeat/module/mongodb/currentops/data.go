package currentops

import (
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

var schemaNone = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host"),
	"desc":              c.Str("desc"),
	"active":            c.Bool("active"),
	"current_op_time":   c.Str("currentOpTime"),
	"opid":              c.Str("opid"),
	"op":                c.Str("op"),
	"ns":                c.Str("ns"),
	"num_yields":        c.Int("numYields"),
	"waiting_for_locks": c.Bool("waitingForLock"),
}

var schemaCommand = s.Schema{
	"shard":         c.Str("shard", s.Optional),
	"host":          c.Str("host"),
	"desc":          c.Str("desc"),
	"connection_id": c.Int("connectionId"),
	"client_s":      c.Str("client_s"),
	"app_name":      c.Str("appName"),
	// TODO: Add client metadata
	"active":            c.Bool("active"),
	"current_op_time":   c.Str("currentOpTime"),
	"opid":              c.Str("opid"),
	"secs_running":      c.Int("secs_running"),
	"microsecs_running": c.Int("microsecs_running"),
	"op":                c.Str("op"),
	"ns":                c.Str("ns"),
	// TODO: Add command
	"num_yields":        c.Int("numYields"),
	"waiting_for_locks": c.Bool("waitingForLock"),
}

var schemaGetMore = s.Schema{
	"shard":         c.Str("shard", s.Optional),
	"host":          c.Str("host"),
	"desc":          c.Str("desc"),
	"connection_id": c.Int("connectionId"),
	"client_s":      c.Str("client_s"),
	// TODO: Add client metadata
	"active": c.Bool("active"),
	// TODO: Parse this as time
	"current_op_time":   c.Str("currentOpTime"),
	"opid":              c.Str("opid"),
	"secs_running":      c.Int("secs_running"),
	"microsecs_running": c.Int("microsecs_running"),
	"op":                c.Str("op"),
	"ns":                c.Str("ns"),
	// TODO: Add command
	// TODO: Add originating command
	"plan_summary":      c.Str("planSummary"),
	"num_yields":        c.Int("numYields"),
	"waiting_for_locks": c.Bool("waitingForLock"),
}

var schemaUpdate = s.Schema{}
var schemaInsert = s.Schema{}
var schemaQuery = s.Schema{}
var schemaRemove = s.Schema{}
var schemaKillCursors = s.Schema{}
