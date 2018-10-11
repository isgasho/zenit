package cpu

import (
	"fmt"

	"github.com/swapbyt3s/zenit/common"
	"github.com/swapbyt3s/zenit/common/log"
	"github.com/swapbyt3s/zenit/config"
	"github.com/swapbyt3s/zenit/plugins/lists/accumulator"
	"github.com/swapbyt3s/zenit/plugins/lists/alerts"
)

func Check() {
	if ! config.File.OS.Alerts.CPU.Enable {
		log.Info("Require to enable OS CPU in config file.")
		return
	}

	var metrics = accumulator.Load()
	var message string = ""
	var value = metrics.FetchOne("os", "name", "cpu")
	var percentage = common.InterfaceToInt(value)

	message += fmt.Sprintf("*CPU:* %d\n", percentage)

	alerts.Load().Register(
		"cpu",
		"CPU",
		config.File.OS.Alerts.CPU.Duration,
		config.File.OS.Alerts.CPU.Warning,
		config.File.OS.Alerts.CPU.Critical,
		percentage,
		message,
	)
}