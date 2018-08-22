package vmx

import (
	govmx "github.com/hooklift/govmx"
)

type VM struct {
	// Disable swap: http://artykul8.com/2012/06/vmware-performance-enhancing/
	MinVmMemPct                 int  `vmx:"prefvmx.minVmMemPct"`
	MemTrimRate                 int  `vmx:"MemTrimRate"`
	UseNamedFile                bool `vmx:"mainMem.useNamedFile"`
	Pshare                      bool `vmx:"sched.mem.pshare.enable"`
	UseRecommendedLockedMemSize bool `vmx:"prefvmx.useRecommendedLockedMemSize"`

	govmx.VirtualMachine
}
