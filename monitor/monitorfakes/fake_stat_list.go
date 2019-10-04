// Code generated by counterfeiter. DO NOT EDIT.
package monitorfakes

import (
	"sync"

	"github.com/athornton2012/http_monitor/monitor"
)

type FakeStatList struct {
	FlushAllStub        func() string
	flushAllMutex       sync.RWMutex
	flushAllArgsForCall []struct {
	}
	flushAllReturns struct {
		result1 string
	}
	flushAllReturnsOnCall map[int]struct {
		result1 string
	}
	UpdateStatListStub        func() string
	updateStatListMutex       sync.RWMutex
	updateStatListArgsForCall []struct {
	}
	updateStatListReturns struct {
		result1 string
	}
	updateStatListReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStatList) FlushAll() string {
	fake.flushAllMutex.Lock()
	ret, specificReturn := fake.flushAllReturnsOnCall[len(fake.flushAllArgsForCall)]
	fake.flushAllArgsForCall = append(fake.flushAllArgsForCall, struct {
	}{})
	fake.recordInvocation("FlushAll", []interface{}{})
	fake.flushAllMutex.Unlock()
	if fake.FlushAllStub != nil {
		return fake.FlushAllStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.flushAllReturns
	return fakeReturns.result1
}

func (fake *FakeStatList) FlushAllCallCount() int {
	fake.flushAllMutex.RLock()
	defer fake.flushAllMutex.RUnlock()
	return len(fake.flushAllArgsForCall)
}

func (fake *FakeStatList) FlushAllCalls(stub func() string) {
	fake.flushAllMutex.Lock()
	defer fake.flushAllMutex.Unlock()
	fake.FlushAllStub = stub
}

func (fake *FakeStatList) FlushAllReturns(result1 string) {
	fake.flushAllMutex.Lock()
	defer fake.flushAllMutex.Unlock()
	fake.FlushAllStub = nil
	fake.flushAllReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStatList) FlushAllReturnsOnCall(i int, result1 string) {
	fake.flushAllMutex.Lock()
	defer fake.flushAllMutex.Unlock()
	fake.FlushAllStub = nil
	if fake.flushAllReturnsOnCall == nil {
		fake.flushAllReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.flushAllReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStatList) UpdateStatList() string {
	fake.updateStatListMutex.Lock()
	ret, specificReturn := fake.updateStatListReturnsOnCall[len(fake.updateStatListArgsForCall)]
	fake.updateStatListArgsForCall = append(fake.updateStatListArgsForCall, struct {
	}{})
	fake.recordInvocation("UpdateStatList", []interface{}{})
	fake.updateStatListMutex.Unlock()
	if fake.UpdateStatListStub != nil {
		return fake.UpdateStatListStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.updateStatListReturns
	return fakeReturns.result1
}

func (fake *FakeStatList) UpdateStatListCallCount() int {
	fake.updateStatListMutex.RLock()
	defer fake.updateStatListMutex.RUnlock()
	return len(fake.updateStatListArgsForCall)
}

func (fake *FakeStatList) UpdateStatListCalls(stub func() string) {
	fake.updateStatListMutex.Lock()
	defer fake.updateStatListMutex.Unlock()
	fake.UpdateStatListStub = stub
}

func (fake *FakeStatList) UpdateStatListReturns(result1 string) {
	fake.updateStatListMutex.Lock()
	defer fake.updateStatListMutex.Unlock()
	fake.UpdateStatListStub = nil
	fake.updateStatListReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStatList) UpdateStatListReturnsOnCall(i int, result1 string) {
	fake.updateStatListMutex.Lock()
	defer fake.updateStatListMutex.Unlock()
	fake.UpdateStatListStub = nil
	if fake.updateStatListReturnsOnCall == nil {
		fake.updateStatListReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.updateStatListReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStatList) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.flushAllMutex.RLock()
	defer fake.flushAllMutex.RUnlock()
	fake.updateStatListMutex.RLock()
	defer fake.updateStatListMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStatList) recordInvocation(key string, args []interface{}) {
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

var _ monitor.StatList = new(FakeStatList)