package agents

import (
	"time"

	"github.com/raintank/fakemetrics_ng/datagen"
	"github.com/raintank/fakemetrics_ng/out"
	"github.com/raintank/fakemetrics_ng/timer"
)

type Agent struct {
	timer   timer.Timer
	datagen datagen.Datagen
	out     out.Out
	offset  int
}

func (a *Agent) Run() {
	time.Sleep(time.Duration(a.offset))

	a.out.Start()
	tick := a.timer.GetTicker()
	for range tick {
		go a.doTick()
	}
}

func (a *Agent) doTick() {
	metrics := a.datagen.GetData(a.timer.GetTimestamp())
	for _, m := range metrics {
		a.out.Put(m)
	}
}
