package datagen

import (
	"flag"
	"fmt"
	gc "github.com/rakyll/globalconf"
	"gopkg.in/raintank/schema.v1"
)

type Incrementer struct {
	id       int
	maxValue float64
	minValue float64
	curValue float64
	metName  string
}

var (
	maxValue     float64
	minValue     float64
	incKeyPrefix string
)

func init() {
	modules["incrementer"] = incNew
	regFlags = append(regFlags, incRegFlags)
}

func incNew(id int) Datagen {
	return &Incrementer{id, maxValue, minValue, 0, fmt.Sprintf(incKeyPrefix+"%d", id)}
}

func incRegFlags() {
	flags := flag.NewFlagSet("incrementer", flag.ExitOnError)
	flags.Float64Var(&maxValue, "max-value", 100, "integer at which we reset to 0")
	flags.Float64Var(&minValue, "min-value", 0, "integer at which we start")
	flags.StringVar(&incKeyPrefix, "key-prefix", "fake.incrementer", "prefix for keys")
	gc.Register("incrementer", flags)
}

func (i *Incrementer) GetData(ts int64) []*schema.MetricData {
	metrics := make([]*schema.MetricData, 1)
	metrics[0] = &schema.MetricData{
		Name:   i.metName,
		Metric: i.metName,
		OrgId:  1,
		Value:  i.curValue,
		Unit:   "ms",
		Mtype:  "gauge",
		Tags:   []string{},
		Time:   ts,
	}

	i.curValue++
	if i.curValue >= i.maxValue {
		i.curValue = i.minValue
	}

	return metrics
}
