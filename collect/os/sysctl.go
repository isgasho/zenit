package os

import (
  "gitlab.com/swapbyt3s/zenit/common"
  "gitlab.com/swapbyt3s/zenit/accumulator"
)

const NR_OPEN string  = "/proc/sys/fs/nr_open"
const FILE_MAX string = "/proc/sys/fs/file-max"

func GatherSysLimits(){
  accumulator.Load().AddItem(accumulator.Metric{
    Key: "os",
    Tags: []accumulator.Tag{accumulator.Tag{"system", "linux"},
                            accumulator.Tag{"setting", "sysctl"}},
    Values: []accumulator.Value{accumulator.Value{"nr_open", common.ValueFromFile(NR_OPEN)},
                                accumulator.Value{"file_max", common.ValueFromFile(FILE_MAX)}},
  })
}
