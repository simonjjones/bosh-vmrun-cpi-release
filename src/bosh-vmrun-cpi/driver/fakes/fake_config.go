// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"bosh-vmrun-cpi/driver"
	"sync"
	"time"
)

type FakeConfig struct {
	OvftoolPathStub        func() string
	ovftoolPathMutex       sync.RWMutex
	ovftoolPathArgsForCall []struct{}
	ovftoolPathReturns     struct {
		result1 string
	}
	ovftoolPathReturnsOnCall map[int]struct {
		result1 string
	}
	VmrunPathStub        func() string
	vmrunPathMutex       sync.RWMutex
	vmrunPathArgsForCall []struct{}
	vmrunPathReturns     struct {
		result1 string
	}
	vmrunPathReturnsOnCall map[int]struct {
		result1 string
	}
	VdiskmanagerPathStub        func() string
	vdiskmanagerPathMutex       sync.RWMutex
	vdiskmanagerPathArgsForCall []struct{}
	vdiskmanagerPathReturns     struct {
		result1 string
	}
	vdiskmanagerPathReturnsOnCall map[int]struct {
		result1 string
	}
	VmPathStub        func() string
	vmPathMutex       sync.RWMutex
	vmPathArgsForCall []struct{}
	vmPathReturns     struct {
		result1 string
	}
	vmPathReturnsOnCall map[int]struct {
		result1 string
	}
	VmStartMaxWaitStub        func() time.Duration
	vmStartMaxWaitMutex       sync.RWMutex
	vmStartMaxWaitArgsForCall []struct{}
	vmStartMaxWaitReturns     struct {
		result1 time.Duration
	}
	vmStartMaxWaitReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	VmSoftShutdownMaxWaitStub        func() time.Duration
	vmSoftShutdownMaxWaitMutex       sync.RWMutex
	vmSoftShutdownMaxWaitArgsForCall []struct{}
	vmSoftShutdownMaxWaitReturns     struct {
		result1 time.Duration
	}
	vmSoftShutdownMaxWaitReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConfig) OvftoolPath() string {
	fake.ovftoolPathMutex.Lock()
	ret, specificReturn := fake.ovftoolPathReturnsOnCall[len(fake.ovftoolPathArgsForCall)]
	fake.ovftoolPathArgsForCall = append(fake.ovftoolPathArgsForCall, struct{}{})
	fake.recordInvocation("OvftoolPath", []interface{}{})
	fake.ovftoolPathMutex.Unlock()
	if fake.OvftoolPathStub != nil {
		return fake.OvftoolPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.ovftoolPathReturns.result1
}

func (fake *FakeConfig) OvftoolPathCallCount() int {
	fake.ovftoolPathMutex.RLock()
	defer fake.ovftoolPathMutex.RUnlock()
	return len(fake.ovftoolPathArgsForCall)
}

func (fake *FakeConfig) OvftoolPathReturns(result1 string) {
	fake.OvftoolPathStub = nil
	fake.ovftoolPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) OvftoolPathReturnsOnCall(i int, result1 string) {
	fake.OvftoolPathStub = nil
	if fake.ovftoolPathReturnsOnCall == nil {
		fake.ovftoolPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.ovftoolPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VmrunPath() string {
	fake.vmrunPathMutex.Lock()
	ret, specificReturn := fake.vmrunPathReturnsOnCall[len(fake.vmrunPathArgsForCall)]
	fake.vmrunPathArgsForCall = append(fake.vmrunPathArgsForCall, struct{}{})
	fake.recordInvocation("VmrunPath", []interface{}{})
	fake.vmrunPathMutex.Unlock()
	if fake.VmrunPathStub != nil {
		return fake.VmrunPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vmrunPathReturns.result1
}

func (fake *FakeConfig) VmrunPathCallCount() int {
	fake.vmrunPathMutex.RLock()
	defer fake.vmrunPathMutex.RUnlock()
	return len(fake.vmrunPathArgsForCall)
}

func (fake *FakeConfig) VmrunPathReturns(result1 string) {
	fake.VmrunPathStub = nil
	fake.vmrunPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VmrunPathReturnsOnCall(i int, result1 string) {
	fake.VmrunPathStub = nil
	if fake.vmrunPathReturnsOnCall == nil {
		fake.vmrunPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.vmrunPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VdiskmanagerPath() string {
	fake.vdiskmanagerPathMutex.Lock()
	ret, specificReturn := fake.vdiskmanagerPathReturnsOnCall[len(fake.vdiskmanagerPathArgsForCall)]
	fake.vdiskmanagerPathArgsForCall = append(fake.vdiskmanagerPathArgsForCall, struct{}{})
	fake.recordInvocation("VdiskmanagerPath", []interface{}{})
	fake.vdiskmanagerPathMutex.Unlock()
	if fake.VdiskmanagerPathStub != nil {
		return fake.VdiskmanagerPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vdiskmanagerPathReturns.result1
}

func (fake *FakeConfig) VdiskmanagerPathCallCount() int {
	fake.vdiskmanagerPathMutex.RLock()
	defer fake.vdiskmanagerPathMutex.RUnlock()
	return len(fake.vdiskmanagerPathArgsForCall)
}

func (fake *FakeConfig) VdiskmanagerPathReturns(result1 string) {
	fake.VdiskmanagerPathStub = nil
	fake.vdiskmanagerPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VdiskmanagerPathReturnsOnCall(i int, result1 string) {
	fake.VdiskmanagerPathStub = nil
	if fake.vdiskmanagerPathReturnsOnCall == nil {
		fake.vdiskmanagerPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.vdiskmanagerPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VmPath() string {
	fake.vmPathMutex.Lock()
	ret, specificReturn := fake.vmPathReturnsOnCall[len(fake.vmPathArgsForCall)]
	fake.vmPathArgsForCall = append(fake.vmPathArgsForCall, struct{}{})
	fake.recordInvocation("VmPath", []interface{}{})
	fake.vmPathMutex.Unlock()
	if fake.VmPathStub != nil {
		return fake.VmPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vmPathReturns.result1
}

func (fake *FakeConfig) VmPathCallCount() int {
	fake.vmPathMutex.RLock()
	defer fake.vmPathMutex.RUnlock()
	return len(fake.vmPathArgsForCall)
}

func (fake *FakeConfig) VmPathReturns(result1 string) {
	fake.VmPathStub = nil
	fake.vmPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VmPathReturnsOnCall(i int, result1 string) {
	fake.VmPathStub = nil
	if fake.vmPathReturnsOnCall == nil {
		fake.vmPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.vmPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeConfig) VmStartMaxWait() time.Duration {
	fake.vmStartMaxWaitMutex.Lock()
	ret, specificReturn := fake.vmStartMaxWaitReturnsOnCall[len(fake.vmStartMaxWaitArgsForCall)]
	fake.vmStartMaxWaitArgsForCall = append(fake.vmStartMaxWaitArgsForCall, struct{}{})
	fake.recordInvocation("VmStartMaxWait", []interface{}{})
	fake.vmStartMaxWaitMutex.Unlock()
	if fake.VmStartMaxWaitStub != nil {
		return fake.VmStartMaxWaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vmStartMaxWaitReturns.result1
}

func (fake *FakeConfig) VmStartMaxWaitCallCount() int {
	fake.vmStartMaxWaitMutex.RLock()
	defer fake.vmStartMaxWaitMutex.RUnlock()
	return len(fake.vmStartMaxWaitArgsForCall)
}

func (fake *FakeConfig) VmStartMaxWaitReturns(result1 time.Duration) {
	fake.VmStartMaxWaitStub = nil
	fake.vmStartMaxWaitReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) VmStartMaxWaitReturnsOnCall(i int, result1 time.Duration) {
	fake.VmStartMaxWaitStub = nil
	if fake.vmStartMaxWaitReturnsOnCall == nil {
		fake.vmStartMaxWaitReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.vmStartMaxWaitReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) VmSoftShutdownMaxWait() time.Duration {
	fake.vmSoftShutdownMaxWaitMutex.Lock()
	ret, specificReturn := fake.vmSoftShutdownMaxWaitReturnsOnCall[len(fake.vmSoftShutdownMaxWaitArgsForCall)]
	fake.vmSoftShutdownMaxWaitArgsForCall = append(fake.vmSoftShutdownMaxWaitArgsForCall, struct{}{})
	fake.recordInvocation("VmSoftShutdownMaxWait", []interface{}{})
	fake.vmSoftShutdownMaxWaitMutex.Unlock()
	if fake.VmSoftShutdownMaxWaitStub != nil {
		return fake.VmSoftShutdownMaxWaitStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.vmSoftShutdownMaxWaitReturns.result1
}

func (fake *FakeConfig) VmSoftShutdownMaxWaitCallCount() int {
	fake.vmSoftShutdownMaxWaitMutex.RLock()
	defer fake.vmSoftShutdownMaxWaitMutex.RUnlock()
	return len(fake.vmSoftShutdownMaxWaitArgsForCall)
}

func (fake *FakeConfig) VmSoftShutdownMaxWaitReturns(result1 time.Duration) {
	fake.VmSoftShutdownMaxWaitStub = nil
	fake.vmSoftShutdownMaxWaitReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) VmSoftShutdownMaxWaitReturnsOnCall(i int, result1 time.Duration) {
	fake.VmSoftShutdownMaxWaitStub = nil
	if fake.vmSoftShutdownMaxWaitReturnsOnCall == nil {
		fake.vmSoftShutdownMaxWaitReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.vmSoftShutdownMaxWaitReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ovftoolPathMutex.RLock()
	defer fake.ovftoolPathMutex.RUnlock()
	fake.vmrunPathMutex.RLock()
	defer fake.vmrunPathMutex.RUnlock()
	fake.vdiskmanagerPathMutex.RLock()
	defer fake.vdiskmanagerPathMutex.RUnlock()
	fake.vmPathMutex.RLock()
	defer fake.vmPathMutex.RUnlock()
	fake.vmStartMaxWaitMutex.RLock()
	defer fake.vmStartMaxWaitMutex.RUnlock()
	fake.vmSoftShutdownMaxWaitMutex.RLock()
	defer fake.vmSoftShutdownMaxWaitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConfig) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ driver.Config = new(FakeConfig)
