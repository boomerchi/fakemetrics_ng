package simple

import (
	"flag"
	"fmt"

	"gopkg.in/raintank/schema.v1"

	mod "github.com/OOM-Killer/fakemetrics_ng/data_gen/module"
	gc "github.com/rakyll/globalconf"
)

var (
	keyCount  int
	keyPrefix string
)

var Module *mod.ModuleT = &mod.ModuleT{
	"simple",
	func(id int) mod.DataGen { return &Simple{id} },
	RegisterFlagSet,
}

type Simple struct {
	agentId int
}

func RegisterFlagSet() {
	flags := flag.NewFlagSet("simple", flag.ExitOnError)
	flags.IntVar(&keyCount, "key-count", 100, "number of keys to generate")
	flags.StringVar(&keyPrefix, "key-prefix", "some.key.", "prefix for keys")
	gc.Register("simple", flags)
}

func (s *Simple) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, keyCount)

	for i := 0; i < keyCount; i++ {
		metrics[i] = &schema.MetricData{
			Name:   fmt.Sprintf(keyPrefix+"%d.%d", s.agentId, i),
			Metric: fmt.Sprintf(keyPrefix+"%d.%d", s.agentId, i),
			OrgId:  i,
			Value:  0,
			Unit:   "ms",
			Mtype:  "gauge",
			Tags:   []string{"some_tag", "ok", "k:2"},
			Time:   ts,
		}
	}
	return metrics
}
