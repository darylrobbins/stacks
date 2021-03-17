// Code generated by counterfeiter. DO NOT EDIT.
package stackfakes

import (
	"sync"

	"github.com/paketo-buildpacks/stacks/create-stack/pkg/stack"
)

type FakeStack struct {
	GetBaseBuildArgsStub        func() []string
	getBaseBuildArgsMutex       sync.RWMutex
	getBaseBuildArgsArgsForCall []struct {
	}
	getBaseBuildArgsReturns struct {
		result1 []string
	}
	getBaseBuildArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetBaseBuildDockerfilePathStub        func() string
	getBaseBuildDockerfilePathMutex       sync.RWMutex
	getBaseBuildDockerfilePathArgsForCall []struct {
	}
	getBaseBuildDockerfilePathReturns struct {
		result1 string
	}
	getBaseBuildDockerfilePathReturnsOnCall map[int]struct {
		result1 string
	}
	GetBaseRunArgsStub        func() []string
	getBaseRunArgsMutex       sync.RWMutex
	getBaseRunArgsArgsForCall []struct {
	}
	getBaseRunArgsReturns struct {
		result1 []string
	}
	getBaseRunArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetBaseRunDockerfilePathStub        func() string
	getBaseRunDockerfilePathMutex       sync.RWMutex
	getBaseRunDockerfilePathArgsForCall []struct {
	}
	getBaseRunDockerfilePathReturns struct {
		result1 string
	}
	getBaseRunDockerfilePathReturnsOnCall map[int]struct {
		result1 string
	}
	GetBuildDescriptionStub        func() string
	getBuildDescriptionMutex       sync.RWMutex
	getBuildDescriptionArgsForCall []struct {
	}
	getBuildDescriptionReturns struct {
		result1 string
	}
	getBuildDescriptionReturnsOnCall map[int]struct {
		result1 string
	}
	GetCNBBuildArgsStub        func() []string
	getCNBBuildArgsMutex       sync.RWMutex
	getCNBBuildArgsArgsForCall []struct {
	}
	getCNBBuildArgsReturns struct {
		result1 []string
	}
	getCNBBuildArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetCNBBuildDockerfilePathStub        func() string
	getCNBBuildDockerfilePathMutex       sync.RWMutex
	getCNBBuildDockerfilePathArgsForCall []struct {
	}
	getCNBBuildDockerfilePathReturns struct {
		result1 string
	}
	getCNBBuildDockerfilePathReturnsOnCall map[int]struct {
		result1 string
	}
	GetCNBRunArgsStub        func() []string
	getCNBRunArgsMutex       sync.RWMutex
	getCNBRunArgsArgsForCall []struct {
	}
	getCNBRunArgsReturns struct {
		result1 []string
	}
	getCNBRunArgsReturnsOnCall map[int]struct {
		result1 []string
	}
	GetCNBRunDockerfilePathStub        func() string
	getCNBRunDockerfilePathMutex       sync.RWMutex
	getCNBRunDockerfilePathArgsForCall []struct {
	}
	getCNBRunDockerfilePathReturns struct {
		result1 string
	}
	getCNBRunDockerfilePathReturnsOnCall map[int]struct {
		result1 string
	}
	GetNameStub        func() string
	getNameMutex       sync.RWMutex
	getNameArgsForCall []struct {
	}
	getNameReturns struct {
		result1 string
	}
	getNameReturnsOnCall map[int]struct {
		result1 string
	}
	GetRunDescriptionStub        func() string
	getRunDescriptionMutex       sync.RWMutex
	getRunDescriptionArgsForCall []struct {
	}
	getRunDescriptionReturns struct {
		result1 string
	}
	getRunDescriptionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStack) GetBaseBuildArgs() []string {
	fake.getBaseBuildArgsMutex.Lock()
	ret, specificReturn := fake.getBaseBuildArgsReturnsOnCall[len(fake.getBaseBuildArgsArgsForCall)]
	fake.getBaseBuildArgsArgsForCall = append(fake.getBaseBuildArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBaseBuildArgs", []interface{}{})
	fake.getBaseBuildArgsMutex.Unlock()
	if fake.GetBaseBuildArgsStub != nil {
		return fake.GetBaseBuildArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBaseBuildArgsReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetBaseBuildArgsCallCount() int {
	fake.getBaseBuildArgsMutex.RLock()
	defer fake.getBaseBuildArgsMutex.RUnlock()
	return len(fake.getBaseBuildArgsArgsForCall)
}

func (fake *FakeStack) GetBaseBuildArgsCalls(stub func() []string) {
	fake.getBaseBuildArgsMutex.Lock()
	defer fake.getBaseBuildArgsMutex.Unlock()
	fake.GetBaseBuildArgsStub = stub
}

func (fake *FakeStack) GetBaseBuildArgsReturns(result1 []string) {
	fake.getBaseBuildArgsMutex.Lock()
	defer fake.getBaseBuildArgsMutex.Unlock()
	fake.GetBaseBuildArgsStub = nil
	fake.getBaseBuildArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetBaseBuildArgsReturnsOnCall(i int, result1 []string) {
	fake.getBaseBuildArgsMutex.Lock()
	defer fake.getBaseBuildArgsMutex.Unlock()
	fake.GetBaseBuildArgsStub = nil
	if fake.getBaseBuildArgsReturnsOnCall == nil {
		fake.getBaseBuildArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getBaseBuildArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetBaseBuildDockerfilePath() string {
	fake.getBaseBuildDockerfilePathMutex.Lock()
	ret, specificReturn := fake.getBaseBuildDockerfilePathReturnsOnCall[len(fake.getBaseBuildDockerfilePathArgsForCall)]
	fake.getBaseBuildDockerfilePathArgsForCall = append(fake.getBaseBuildDockerfilePathArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBaseBuildDockerfilePath", []interface{}{})
	fake.getBaseBuildDockerfilePathMutex.Unlock()
	if fake.GetBaseBuildDockerfilePathStub != nil {
		return fake.GetBaseBuildDockerfilePathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBaseBuildDockerfilePathReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetBaseBuildDockerfilePathCallCount() int {
	fake.getBaseBuildDockerfilePathMutex.RLock()
	defer fake.getBaseBuildDockerfilePathMutex.RUnlock()
	return len(fake.getBaseBuildDockerfilePathArgsForCall)
}

func (fake *FakeStack) GetBaseBuildDockerfilePathCalls(stub func() string) {
	fake.getBaseBuildDockerfilePathMutex.Lock()
	defer fake.getBaseBuildDockerfilePathMutex.Unlock()
	fake.GetBaseBuildDockerfilePathStub = stub
}

func (fake *FakeStack) GetBaseBuildDockerfilePathReturns(result1 string) {
	fake.getBaseBuildDockerfilePathMutex.Lock()
	defer fake.getBaseBuildDockerfilePathMutex.Unlock()
	fake.GetBaseBuildDockerfilePathStub = nil
	fake.getBaseBuildDockerfilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetBaseBuildDockerfilePathReturnsOnCall(i int, result1 string) {
	fake.getBaseBuildDockerfilePathMutex.Lock()
	defer fake.getBaseBuildDockerfilePathMutex.Unlock()
	fake.GetBaseBuildDockerfilePathStub = nil
	if fake.getBaseBuildDockerfilePathReturnsOnCall == nil {
		fake.getBaseBuildDockerfilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getBaseBuildDockerfilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetBaseRunArgs() []string {
	fake.getBaseRunArgsMutex.Lock()
	ret, specificReturn := fake.getBaseRunArgsReturnsOnCall[len(fake.getBaseRunArgsArgsForCall)]
	fake.getBaseRunArgsArgsForCall = append(fake.getBaseRunArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBaseRunArgs", []interface{}{})
	fake.getBaseRunArgsMutex.Unlock()
	if fake.GetBaseRunArgsStub != nil {
		return fake.GetBaseRunArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBaseRunArgsReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetBaseRunArgsCallCount() int {
	fake.getBaseRunArgsMutex.RLock()
	defer fake.getBaseRunArgsMutex.RUnlock()
	return len(fake.getBaseRunArgsArgsForCall)
}

func (fake *FakeStack) GetBaseRunArgsCalls(stub func() []string) {
	fake.getBaseRunArgsMutex.Lock()
	defer fake.getBaseRunArgsMutex.Unlock()
	fake.GetBaseRunArgsStub = stub
}

func (fake *FakeStack) GetBaseRunArgsReturns(result1 []string) {
	fake.getBaseRunArgsMutex.Lock()
	defer fake.getBaseRunArgsMutex.Unlock()
	fake.GetBaseRunArgsStub = nil
	fake.getBaseRunArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetBaseRunArgsReturnsOnCall(i int, result1 []string) {
	fake.getBaseRunArgsMutex.Lock()
	defer fake.getBaseRunArgsMutex.Unlock()
	fake.GetBaseRunArgsStub = nil
	if fake.getBaseRunArgsReturnsOnCall == nil {
		fake.getBaseRunArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getBaseRunArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetBaseRunDockerfilePath() string {
	fake.getBaseRunDockerfilePathMutex.Lock()
	ret, specificReturn := fake.getBaseRunDockerfilePathReturnsOnCall[len(fake.getBaseRunDockerfilePathArgsForCall)]
	fake.getBaseRunDockerfilePathArgsForCall = append(fake.getBaseRunDockerfilePathArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBaseRunDockerfilePath", []interface{}{})
	fake.getBaseRunDockerfilePathMutex.Unlock()
	if fake.GetBaseRunDockerfilePathStub != nil {
		return fake.GetBaseRunDockerfilePathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBaseRunDockerfilePathReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetBaseRunDockerfilePathCallCount() int {
	fake.getBaseRunDockerfilePathMutex.RLock()
	defer fake.getBaseRunDockerfilePathMutex.RUnlock()
	return len(fake.getBaseRunDockerfilePathArgsForCall)
}

func (fake *FakeStack) GetBaseRunDockerfilePathCalls(stub func() string) {
	fake.getBaseRunDockerfilePathMutex.Lock()
	defer fake.getBaseRunDockerfilePathMutex.Unlock()
	fake.GetBaseRunDockerfilePathStub = stub
}

func (fake *FakeStack) GetBaseRunDockerfilePathReturns(result1 string) {
	fake.getBaseRunDockerfilePathMutex.Lock()
	defer fake.getBaseRunDockerfilePathMutex.Unlock()
	fake.GetBaseRunDockerfilePathStub = nil
	fake.getBaseRunDockerfilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetBaseRunDockerfilePathReturnsOnCall(i int, result1 string) {
	fake.getBaseRunDockerfilePathMutex.Lock()
	defer fake.getBaseRunDockerfilePathMutex.Unlock()
	fake.GetBaseRunDockerfilePathStub = nil
	if fake.getBaseRunDockerfilePathReturnsOnCall == nil {
		fake.getBaseRunDockerfilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getBaseRunDockerfilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetBuildDescription() string {
	fake.getBuildDescriptionMutex.Lock()
	ret, specificReturn := fake.getBuildDescriptionReturnsOnCall[len(fake.getBuildDescriptionArgsForCall)]
	fake.getBuildDescriptionArgsForCall = append(fake.getBuildDescriptionArgsForCall, struct {
	}{})
	fake.recordInvocation("GetBuildDescription", []interface{}{})
	fake.getBuildDescriptionMutex.Unlock()
	if fake.GetBuildDescriptionStub != nil {
		return fake.GetBuildDescriptionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getBuildDescriptionReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetBuildDescriptionCallCount() int {
	fake.getBuildDescriptionMutex.RLock()
	defer fake.getBuildDescriptionMutex.RUnlock()
	return len(fake.getBuildDescriptionArgsForCall)
}

func (fake *FakeStack) GetBuildDescriptionCalls(stub func() string) {
	fake.getBuildDescriptionMutex.Lock()
	defer fake.getBuildDescriptionMutex.Unlock()
	fake.GetBuildDescriptionStub = stub
}

func (fake *FakeStack) GetBuildDescriptionReturns(result1 string) {
	fake.getBuildDescriptionMutex.Lock()
	defer fake.getBuildDescriptionMutex.Unlock()
	fake.GetBuildDescriptionStub = nil
	fake.getBuildDescriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetBuildDescriptionReturnsOnCall(i int, result1 string) {
	fake.getBuildDescriptionMutex.Lock()
	defer fake.getBuildDescriptionMutex.Unlock()
	fake.GetBuildDescriptionStub = nil
	if fake.getBuildDescriptionReturnsOnCall == nil {
		fake.getBuildDescriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getBuildDescriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetCNBBuildArgs() []string {
	fake.getCNBBuildArgsMutex.Lock()
	ret, specificReturn := fake.getCNBBuildArgsReturnsOnCall[len(fake.getCNBBuildArgsArgsForCall)]
	fake.getCNBBuildArgsArgsForCall = append(fake.getCNBBuildArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCNBBuildArgs", []interface{}{})
	fake.getCNBBuildArgsMutex.Unlock()
	if fake.GetCNBBuildArgsStub != nil {
		return fake.GetCNBBuildArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCNBBuildArgsReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetCNBBuildArgsCallCount() int {
	fake.getCNBBuildArgsMutex.RLock()
	defer fake.getCNBBuildArgsMutex.RUnlock()
	return len(fake.getCNBBuildArgsArgsForCall)
}

func (fake *FakeStack) GetCNBBuildArgsCalls(stub func() []string) {
	fake.getCNBBuildArgsMutex.Lock()
	defer fake.getCNBBuildArgsMutex.Unlock()
	fake.GetCNBBuildArgsStub = stub
}

func (fake *FakeStack) GetCNBBuildArgsReturns(result1 []string) {
	fake.getCNBBuildArgsMutex.Lock()
	defer fake.getCNBBuildArgsMutex.Unlock()
	fake.GetCNBBuildArgsStub = nil
	fake.getCNBBuildArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetCNBBuildArgsReturnsOnCall(i int, result1 []string) {
	fake.getCNBBuildArgsMutex.Lock()
	defer fake.getCNBBuildArgsMutex.Unlock()
	fake.GetCNBBuildArgsStub = nil
	if fake.getCNBBuildArgsReturnsOnCall == nil {
		fake.getCNBBuildArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getCNBBuildArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetCNBBuildDockerfilePath() string {
	fake.getCNBBuildDockerfilePathMutex.Lock()
	ret, specificReturn := fake.getCNBBuildDockerfilePathReturnsOnCall[len(fake.getCNBBuildDockerfilePathArgsForCall)]
	fake.getCNBBuildDockerfilePathArgsForCall = append(fake.getCNBBuildDockerfilePathArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCNBBuildDockerfilePath", []interface{}{})
	fake.getCNBBuildDockerfilePathMutex.Unlock()
	if fake.GetCNBBuildDockerfilePathStub != nil {
		return fake.GetCNBBuildDockerfilePathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCNBBuildDockerfilePathReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetCNBBuildDockerfilePathCallCount() int {
	fake.getCNBBuildDockerfilePathMutex.RLock()
	defer fake.getCNBBuildDockerfilePathMutex.RUnlock()
	return len(fake.getCNBBuildDockerfilePathArgsForCall)
}

func (fake *FakeStack) GetCNBBuildDockerfilePathCalls(stub func() string) {
	fake.getCNBBuildDockerfilePathMutex.Lock()
	defer fake.getCNBBuildDockerfilePathMutex.Unlock()
	fake.GetCNBBuildDockerfilePathStub = stub
}

func (fake *FakeStack) GetCNBBuildDockerfilePathReturns(result1 string) {
	fake.getCNBBuildDockerfilePathMutex.Lock()
	defer fake.getCNBBuildDockerfilePathMutex.Unlock()
	fake.GetCNBBuildDockerfilePathStub = nil
	fake.getCNBBuildDockerfilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetCNBBuildDockerfilePathReturnsOnCall(i int, result1 string) {
	fake.getCNBBuildDockerfilePathMutex.Lock()
	defer fake.getCNBBuildDockerfilePathMutex.Unlock()
	fake.GetCNBBuildDockerfilePathStub = nil
	if fake.getCNBBuildDockerfilePathReturnsOnCall == nil {
		fake.getCNBBuildDockerfilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getCNBBuildDockerfilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetCNBRunArgs() []string {
	fake.getCNBRunArgsMutex.Lock()
	ret, specificReturn := fake.getCNBRunArgsReturnsOnCall[len(fake.getCNBRunArgsArgsForCall)]
	fake.getCNBRunArgsArgsForCall = append(fake.getCNBRunArgsArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCNBRunArgs", []interface{}{})
	fake.getCNBRunArgsMutex.Unlock()
	if fake.GetCNBRunArgsStub != nil {
		return fake.GetCNBRunArgsStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCNBRunArgsReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetCNBRunArgsCallCount() int {
	fake.getCNBRunArgsMutex.RLock()
	defer fake.getCNBRunArgsMutex.RUnlock()
	return len(fake.getCNBRunArgsArgsForCall)
}

func (fake *FakeStack) GetCNBRunArgsCalls(stub func() []string) {
	fake.getCNBRunArgsMutex.Lock()
	defer fake.getCNBRunArgsMutex.Unlock()
	fake.GetCNBRunArgsStub = stub
}

func (fake *FakeStack) GetCNBRunArgsReturns(result1 []string) {
	fake.getCNBRunArgsMutex.Lock()
	defer fake.getCNBRunArgsMutex.Unlock()
	fake.GetCNBRunArgsStub = nil
	fake.getCNBRunArgsReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetCNBRunArgsReturnsOnCall(i int, result1 []string) {
	fake.getCNBRunArgsMutex.Lock()
	defer fake.getCNBRunArgsMutex.Unlock()
	fake.GetCNBRunArgsStub = nil
	if fake.getCNBRunArgsReturnsOnCall == nil {
		fake.getCNBRunArgsReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getCNBRunArgsReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeStack) GetCNBRunDockerfilePath() string {
	fake.getCNBRunDockerfilePathMutex.Lock()
	ret, specificReturn := fake.getCNBRunDockerfilePathReturnsOnCall[len(fake.getCNBRunDockerfilePathArgsForCall)]
	fake.getCNBRunDockerfilePathArgsForCall = append(fake.getCNBRunDockerfilePathArgsForCall, struct {
	}{})
	fake.recordInvocation("GetCNBRunDockerfilePath", []interface{}{})
	fake.getCNBRunDockerfilePathMutex.Unlock()
	if fake.GetCNBRunDockerfilePathStub != nil {
		return fake.GetCNBRunDockerfilePathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getCNBRunDockerfilePathReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetCNBRunDockerfilePathCallCount() int {
	fake.getCNBRunDockerfilePathMutex.RLock()
	defer fake.getCNBRunDockerfilePathMutex.RUnlock()
	return len(fake.getCNBRunDockerfilePathArgsForCall)
}

func (fake *FakeStack) GetCNBRunDockerfilePathCalls(stub func() string) {
	fake.getCNBRunDockerfilePathMutex.Lock()
	defer fake.getCNBRunDockerfilePathMutex.Unlock()
	fake.GetCNBRunDockerfilePathStub = stub
}

func (fake *FakeStack) GetCNBRunDockerfilePathReturns(result1 string) {
	fake.getCNBRunDockerfilePathMutex.Lock()
	defer fake.getCNBRunDockerfilePathMutex.Unlock()
	fake.GetCNBRunDockerfilePathStub = nil
	fake.getCNBRunDockerfilePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetCNBRunDockerfilePathReturnsOnCall(i int, result1 string) {
	fake.getCNBRunDockerfilePathMutex.Lock()
	defer fake.getCNBRunDockerfilePathMutex.Unlock()
	fake.GetCNBRunDockerfilePathStub = nil
	if fake.getCNBRunDockerfilePathReturnsOnCall == nil {
		fake.getCNBRunDockerfilePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getCNBRunDockerfilePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetName() string {
	fake.getNameMutex.Lock()
	ret, specificReturn := fake.getNameReturnsOnCall[len(fake.getNameArgsForCall)]
	fake.getNameArgsForCall = append(fake.getNameArgsForCall, struct {
	}{})
	fake.recordInvocation("GetName", []interface{}{})
	fake.getNameMutex.Unlock()
	if fake.GetNameStub != nil {
		return fake.GetNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getNameReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetNameCallCount() int {
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	return len(fake.getNameArgsForCall)
}

func (fake *FakeStack) GetNameCalls(stub func() string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = stub
}

func (fake *FakeStack) GetNameReturns(result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	fake.getNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetNameReturnsOnCall(i int, result1 string) {
	fake.getNameMutex.Lock()
	defer fake.getNameMutex.Unlock()
	fake.GetNameStub = nil
	if fake.getNameReturnsOnCall == nil {
		fake.getNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetRunDescription() string {
	fake.getRunDescriptionMutex.Lock()
	ret, specificReturn := fake.getRunDescriptionReturnsOnCall[len(fake.getRunDescriptionArgsForCall)]
	fake.getRunDescriptionArgsForCall = append(fake.getRunDescriptionArgsForCall, struct {
	}{})
	fake.recordInvocation("GetRunDescription", []interface{}{})
	fake.getRunDescriptionMutex.Unlock()
	if fake.GetRunDescriptionStub != nil {
		return fake.GetRunDescriptionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getRunDescriptionReturns
	return fakeReturns.result1
}

func (fake *FakeStack) GetRunDescriptionCallCount() int {
	fake.getRunDescriptionMutex.RLock()
	defer fake.getRunDescriptionMutex.RUnlock()
	return len(fake.getRunDescriptionArgsForCall)
}

func (fake *FakeStack) GetRunDescriptionCalls(stub func() string) {
	fake.getRunDescriptionMutex.Lock()
	defer fake.getRunDescriptionMutex.Unlock()
	fake.GetRunDescriptionStub = stub
}

func (fake *FakeStack) GetRunDescriptionReturns(result1 string) {
	fake.getRunDescriptionMutex.Lock()
	defer fake.getRunDescriptionMutex.Unlock()
	fake.GetRunDescriptionStub = nil
	fake.getRunDescriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) GetRunDescriptionReturnsOnCall(i int, result1 string) {
	fake.getRunDescriptionMutex.Lock()
	defer fake.getRunDescriptionMutex.Unlock()
	fake.GetRunDescriptionStub = nil
	if fake.getRunDescriptionReturnsOnCall == nil {
		fake.getRunDescriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.getRunDescriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getBaseBuildArgsMutex.RLock()
	defer fake.getBaseBuildArgsMutex.RUnlock()
	fake.getBaseBuildDockerfilePathMutex.RLock()
	defer fake.getBaseBuildDockerfilePathMutex.RUnlock()
	fake.getBaseRunArgsMutex.RLock()
	defer fake.getBaseRunArgsMutex.RUnlock()
	fake.getBaseRunDockerfilePathMutex.RLock()
	defer fake.getBaseRunDockerfilePathMutex.RUnlock()
	fake.getBuildDescriptionMutex.RLock()
	defer fake.getBuildDescriptionMutex.RUnlock()
	fake.getCNBBuildArgsMutex.RLock()
	defer fake.getCNBBuildArgsMutex.RUnlock()
	fake.getCNBBuildDockerfilePathMutex.RLock()
	defer fake.getCNBBuildDockerfilePathMutex.RUnlock()
	fake.getCNBRunArgsMutex.RLock()
	defer fake.getCNBRunArgsMutex.RUnlock()
	fake.getCNBRunDockerfilePathMutex.RLock()
	defer fake.getCNBRunDockerfilePathMutex.RUnlock()
	fake.getNameMutex.RLock()
	defer fake.getNameMutex.RUnlock()
	fake.getRunDescriptionMutex.RLock()
	defer fake.getRunDescriptionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStack) recordInvocation(key string, args []interface{}) {
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

var _ stack.Stack = new(FakeStack)
