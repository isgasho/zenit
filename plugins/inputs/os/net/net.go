package net

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/swapbyt3s/zenit/common"
	"github.com/swapbyt3s/zenit/common/file"
	"github.com/swapbyt3s/zenit/common/log"
	"github.com/swapbyt3s/zenit/config"
	"github.com/swapbyt3s/zenit/plugins/inputs"
	"github.com/swapbyt3s/zenit/plugins/lists/metrics"
)

type InputOSNet struct{}

func (l *InputOSNet) Collect() {
	defer func() {
		if err := recover(); err != nil {
			log.Debug(fmt.Sprintf("Plugin - InputOSNet - Panic (code %d) has been recover from somewhere.\n", err))
		}
	}()

	if !config.File.Inputs.OS.Net {
		return
	}

	log.Info("Plugin - InputOSNet")

	var a = metrics.Load()

	reGroups := regexp.MustCompile(`(\d+)`)
	net := file.Read("/proc/net/dev")
	lines := strings.Split(net, "\n")

	for index, line := range lines {
		if index > 1 && len(line) > 0 {
			data := strings.Split(line, ":")
			dev := strings.TrimSpace(data[0])
			match := reGroups.FindAllString(data[1], -1)
			receive_bytes := common.StringToInt64(match[0])
			transmit_bytes := common.StringToInt64(match[8])

			a.Add(metrics.Metric{
				Key: "os_net",
				Tags: []metrics.Tag{
					{"hostname", config.File.General.Hostname},
					{"device", dev},
				},
				Values: []metrics.Value{
					{"receive", receive_bytes},
					{"transmit", transmit_bytes},
				},
			})
		}
	}
}

func init() {
	inputs.Add("InputOSNet", func() inputs.Input { return &InputOSNet{} })
}
