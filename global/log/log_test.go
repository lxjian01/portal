package log

import (
	"bn-cmdb/config"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	logConfig := &config.LogConfig{
		Dir:       "/data/cmdb/",
		Name:      "go.log",
		Format:    "1", //no use right now
		RetainDay: 7,
		Level:     "Info",
	}

	Init(logConfig)
	run := make(chan struct{})

	for i := 1; i <= 1; i++ {
		go func() {
			for {
				Info("2342342 - ", i)
				time.Sleep(time.Duration(1000) * time.Millisecond)
			}
		}()
	}

	<-run
}