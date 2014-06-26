package osmain

import (
	"github.com/measure/os/cpustat"
	"github.com/measure/os/memstat"
	"github.com/measure/os/pidstat"
)

// these are implemented by all supported platforms
type OsIndependentStats struct {
	Cstat *cpustat.CPUStat
	Mstat *memstat.MemStat
	Procs *pidstat.ProcessStat
}
