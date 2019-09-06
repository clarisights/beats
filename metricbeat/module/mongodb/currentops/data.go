package currentops

import (
	s "github.com/elastic/beats/libbeat/common/schema"
	c "github.com/elastic/beats/libbeat/common/schema/mapstriface"
)

var schema = s.Schema{
	"shard":             c.Str("shard", s.Optional),
	"host":              c.Str("host"),
	"desc":              c.Str("desc"),
	"connection_id":     c.Int("connectionId"),
	"client_s":          c.Str("client_s"),
	"app_name":          c.Str("appName"),
	"active":            c.Bool("active"),
	"current_op_time":   c.Str("currentOpTime"),
	"opid":              c.Str("opid"),
	"secs_running":      c.Int("secs_running"),
	"microsecs_running": c.Int("microsecs_running"),
	"op":                c.Str("op"),
	"ns":                c.Str("ns"),
}
