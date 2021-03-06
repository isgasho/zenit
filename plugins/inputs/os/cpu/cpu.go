package cpu

import (
	"fmt"

	"github.com/swapbyt3s/zenit/common/log"
	"github.com/swapbyt3s/zenit/config"
	"github.com/swapbyt3s/zenit/plugins/inputs"
	"github.com/swapbyt3s/zenit/plugins/lists/metrics"

	"github.com/shirou/gopsutil/cpu"
)

type InputOSCPU struct{}

func (l *InputOSCPU) Collect() {
	defer func() {
		if err := recover(); err != nil {
			log.Debug(fmt.Sprintf("Plugin - InputOSCPU - Panic (code %d) has been recover from somewhere.\n", err))
		}
	}()

	if !config.File.Inputs.OS.CPU {
		return
	}

	log.Info("Plugin - InputOSCPU")

	var a = metrics.Load()

	percentage, err := cpu.Percent(0, false)
	if err == nil {
		a.Add(metrics.Metric{
			Key: "os_cpu",
			Tags: []metrics.Tag{
				{"hostname", config.File.General.Hostname},
			},
			Values: []metrics.Value{
				{ "percentage", percentage[0] },
			},
		})
	}

	log.Debug(fmt.Sprintf("Plugin - InputOSCPU - CPU=%.2f", percentage))
}

func init() {
	inputs.Add("InputOSCPU", func() inputs.Input { return &InputOSCPU{} })
}
