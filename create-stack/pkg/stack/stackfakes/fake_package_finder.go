// Code generated by counterfeiter. DO NOT EDIT.
package stackfakes

import (
	"sync"

	"github.com/paketo-buildpacks/stacks/create-stack/pkg/stack"
)

type FakePackageFinder struct {
	GetPackageMetadataStub        func(string) (string, error)
	getPackageMetadataMutex       sync.RWMutex
	getPackageMetadataArgsForCall []struct {
		arg1 string
	}
	getPackageMetadataReturns struct {
		result1 string
		result2 error
	}
	getPackageMetadataReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetPackagesListStub        func(string) ([]string, error)
	getPackagesListMutex       sync.RWMutex
	getPackagesListArgsForCall []struct {
		arg1 string
	}
	getPackagesListReturns struct {
		result1 []string
		result2 error
	}
	getPackagesListReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePackageFinder) GetPackageMetadata(arg1 string) (string, error) {
	fake.getPackageMetadataMutex.Lock()
	ret, specificReturn := fake.getPackageMetadataReturnsOnCall[len(fake.getPackageMetadataArgsForCall)]
	fake.getPackageMetadataArgsForCall = append(fake.getPackageMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPackageMetadata", []interface{}{arg1})
	fake.getPackageMetadataMutex.Unlock()
	if fake.GetPackageMetadataStub != nil {
		return fake.GetPackageMetadataStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPackageMetadataReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePackageFinder) GetPackageMetadataCallCount() int {
	fake.getPackageMetadataMutex.RLock()
	defer fake.getPackageMetadataMutex.RUnlock()
	return len(fake.getPackageMetadataArgsForCall)
}

func (fake *FakePackageFinder) GetPackageMetadataCalls(stub func(string) (string, error)) {
	fake.getPackageMetadataMutex.Lock()
	defer fake.getPackageMetadataMutex.Unlock()
	fake.GetPackageMetadataStub = stub
}

func (fake *FakePackageFinder) GetPackageMetadataArgsForCall(i int) string {
	fake.getPackageMetadataMutex.RLock()
	defer fake.getPackageMetadataMutex.RUnlock()
	argsForCall := fake.getPackageMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePackageFinder) GetPackageMetadataReturns(result1 string, result2 error) {
	fake.getPackageMetadataMutex.Lock()
	defer fake.getPackageMetadataMutex.Unlock()
	fake.GetPackageMetadataStub = nil
	fake.getPackageMetadataReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePackageFinder) GetPackageMetadataReturnsOnCall(i int, result1 string, result2 error) {
	fake.getPackageMetadataMutex.Lock()
	defer fake.getPackageMetadataMutex.Unlock()
	fake.GetPackageMetadataStub = nil
	if fake.getPackageMetadataReturnsOnCall == nil {
		fake.getPackageMetadataReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getPackageMetadataReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakePackageFinder) GetPackagesList(arg1 string) ([]string, error) {
	fake.getPackagesListMutex.Lock()
	ret, specificReturn := fake.getPackagesListReturnsOnCall[len(fake.getPackagesListArgsForCall)]
	fake.getPackagesListArgsForCall = append(fake.getPackagesListArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetPackagesList", []interface{}{arg1})
	fake.getPackagesListMutex.Unlock()
	if fake.GetPackagesListStub != nil {
		return fake.GetPackagesListStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getPackagesListReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePackageFinder) GetPackagesListCallCount() int {
	fake.getPackagesListMutex.RLock()
	defer fake.getPackagesListMutex.RUnlock()
	return len(fake.getPackagesListArgsForCall)
}

func (fake *FakePackageFinder) GetPackagesListCalls(stub func(string) ([]string, error)) {
	fake.getPackagesListMutex.Lock()
	defer fake.getPackagesListMutex.Unlock()
	fake.GetPackagesListStub = stub
}

func (fake *FakePackageFinder) GetPackagesListArgsForCall(i int) string {
	fake.getPackagesListMutex.RLock()
	defer fake.getPackagesListMutex.RUnlock()
	argsForCall := fake.getPackagesListArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePackageFinder) GetPackagesListReturns(result1 []string, result2 error) {
	fake.getPackagesListMutex.Lock()
	defer fake.getPackagesListMutex.Unlock()
	fake.GetPackagesListStub = nil
	fake.getPackagesListReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePackageFinder) GetPackagesListReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getPackagesListMutex.Lock()
	defer fake.getPackagesListMutex.Unlock()
	fake.GetPackagesListStub = nil
	if fake.getPackagesListReturnsOnCall == nil {
		fake.getPackagesListReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getPackagesListReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakePackageFinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getPackageMetadataMutex.RLock()
	defer fake.getPackageMetadataMutex.RUnlock()
	fake.getPackagesListMutex.RLock()
	defer fake.getPackagesListMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePackageFinder) recordInvocation(key string, args []interface{}) {
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

var _ stack.PackageFinder = new(FakePackageFinder)
