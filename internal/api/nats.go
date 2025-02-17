// Copyright (C) NHR@FAU, University Erlangen-Nuremberg.
// All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package api

import (
	"sync"

	"github.com/ClusterCockpit/cc-backend/internal/config"
	"github.com/ClusterCockpit/cc-backend/pkg/log"
	lp "github.com/ClusterCockpit/cc-lib/ccMessage"
	"github.com/ClusterCockpit/cc-lib/sinks"
)

type NatsClient struct {
	SinkManager sinks.SinkManager
	SinkChannel chan lp.CCMessage
}

var (
	initOnce sync.Once
	ni       *NatsClient
)

func Init(wg *sync.WaitGroup) {
	initOnce.Do(func() {
		ni = &NatsClient{}
		var err error

		if len(config.Keys.SinkConfigFile) == 0 {
			log.Error("Sink configuration file must be set")
			return
		}
		ni.SinkManager, err = sinks.New(wg, config.Keys.SinkConfigFile)
		if err != nil {
			log.Error(err.Error())
			return
		}

		ni.SinkChannel = make(chan lp.CCMessage, 200)
		ni.SinkManager.AddInput(ni.SinkChannel)
		ni.SinkManager.Start()
	})
}

func Shutdown() {
	if ni.SinkManager != nil {
		log.Debug("Shutdown SinkManager...")
		ni.SinkManager.Close()
	}
}

func forwardJob() {
}
