// This file was generated by counterfeiter
package releasedependencyfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/go-pivnet/cmd/pivnet/commands/releasedependency"
)

type FakePivnetClient struct {
	ReleaseForVersionStub        func(productSlug string, releaseVersion string) (go_pivnet.Release, error)
	releaseForVersionMutex       sync.RWMutex
	releaseForVersionArgsForCall []struct {
		productSlug    string
		releaseVersion string
	}
	releaseForVersionReturns struct {
		result1 go_pivnet.Release
		result2 error
	}
	ReleaseDependenciesStub        func(productSlug string, releaseID int) ([]go_pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseDependenciesReturns struct {
		result1 []go_pivnet.ReleaseDependency
		result2 error
	}
	AddReleaseDependencyStub        func(productSlug string, releaseID int, dependentReleaseID int) error
	addReleaseDependencyMutex       sync.RWMutex
	addReleaseDependencyArgsForCall []struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}
	addReleaseDependencyReturns struct {
		result1 error
	}
	RemoveReleaseDependencyStub        func(productSlug string, releaseID int, dependentReleaseID int) error
	removeReleaseDependencyMutex       sync.RWMutex
	removeReleaseDependencyArgsForCall []struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}
	removeReleaseDependencyReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseForVersion(productSlug string, releaseVersion string) (go_pivnet.Release, error) {
	fake.releaseForVersionMutex.Lock()
	fake.releaseForVersionArgsForCall = append(fake.releaseForVersionArgsForCall, struct {
		productSlug    string
		releaseVersion string
	}{productSlug, releaseVersion})
	fake.recordInvocation("ReleaseForVersion", []interface{}{productSlug, releaseVersion})
	fake.releaseForVersionMutex.Unlock()
	if fake.ReleaseForVersionStub != nil {
		return fake.ReleaseForVersionStub(productSlug, releaseVersion)
	} else {
		return fake.releaseForVersionReturns.result1, fake.releaseForVersionReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseForVersionCallCount() int {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return len(fake.releaseForVersionArgsForCall)
}

func (fake *FakePivnetClient) ReleaseForVersionArgsForCall(i int) (string, string) {
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	return fake.releaseForVersionArgsForCall[i].productSlug, fake.releaseForVersionArgsForCall[i].releaseVersion
}

func (fake *FakePivnetClient) ReleaseForVersionReturns(result1 go_pivnet.Release, result2 error) {
	fake.ReleaseForVersionStub = nil
	fake.releaseForVersionReturns = struct {
		result1 go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependencies(productSlug string, releaseID int) ([]go_pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseDependencies", []interface{}{productSlug, releaseID})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(productSlug, releaseID)
	} else {
		return fake.releaseDependenciesReturns.result1, fake.releaseDependenciesReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return fake.releaseDependenciesArgsForCall[i].productSlug, fake.releaseDependenciesArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseDependenciesReturns(result1 []go_pivnet.ReleaseDependency, result2 error) {
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []go_pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AddReleaseDependency(productSlug string, releaseID int, dependentReleaseID int) error {
	fake.addReleaseDependencyMutex.Lock()
	fake.addReleaseDependencyArgsForCall = append(fake.addReleaseDependencyArgsForCall, struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}{productSlug, releaseID, dependentReleaseID})
	fake.recordInvocation("AddReleaseDependency", []interface{}{productSlug, releaseID, dependentReleaseID})
	fake.addReleaseDependencyMutex.Unlock()
	if fake.AddReleaseDependencyStub != nil {
		return fake.AddReleaseDependencyStub(productSlug, releaseID, dependentReleaseID)
	} else {
		return fake.addReleaseDependencyReturns.result1
	}
}

func (fake *FakePivnetClient) AddReleaseDependencyCallCount() int {
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	return len(fake.addReleaseDependencyArgsForCall)
}

func (fake *FakePivnetClient) AddReleaseDependencyArgsForCall(i int) (string, int, int) {
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	return fake.addReleaseDependencyArgsForCall[i].productSlug, fake.addReleaseDependencyArgsForCall[i].releaseID, fake.addReleaseDependencyArgsForCall[i].dependentReleaseID
}

func (fake *FakePivnetClient) AddReleaseDependencyReturns(result1 error) {
	fake.AddReleaseDependencyStub = nil
	fake.addReleaseDependencyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) RemoveReleaseDependency(productSlug string, releaseID int, dependentReleaseID int) error {
	fake.removeReleaseDependencyMutex.Lock()
	fake.removeReleaseDependencyArgsForCall = append(fake.removeReleaseDependencyArgsForCall, struct {
		productSlug        string
		releaseID          int
		dependentReleaseID int
	}{productSlug, releaseID, dependentReleaseID})
	fake.recordInvocation("RemoveReleaseDependency", []interface{}{productSlug, releaseID, dependentReleaseID})
	fake.removeReleaseDependencyMutex.Unlock()
	if fake.RemoveReleaseDependencyStub != nil {
		return fake.RemoveReleaseDependencyStub(productSlug, releaseID, dependentReleaseID)
	} else {
		return fake.removeReleaseDependencyReturns.result1
	}
}

func (fake *FakePivnetClient) RemoveReleaseDependencyCallCount() int {
	fake.removeReleaseDependencyMutex.RLock()
	defer fake.removeReleaseDependencyMutex.RUnlock()
	return len(fake.removeReleaseDependencyArgsForCall)
}

func (fake *FakePivnetClient) RemoveReleaseDependencyArgsForCall(i int) (string, int, int) {
	fake.removeReleaseDependencyMutex.RLock()
	defer fake.removeReleaseDependencyMutex.RUnlock()
	return fake.removeReleaseDependencyArgsForCall[i].productSlug, fake.removeReleaseDependencyArgsForCall[i].releaseID, fake.removeReleaseDependencyArgsForCall[i].dependentReleaseID
}

func (fake *FakePivnetClient) RemoveReleaseDependencyReturns(result1 error) {
	fake.RemoveReleaseDependencyStub = nil
	fake.removeReleaseDependencyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseForVersionMutex.RLock()
	defer fake.releaseForVersionMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	fake.addReleaseDependencyMutex.RLock()
	defer fake.addReleaseDependencyMutex.RUnlock()
	fake.removeReleaseDependencyMutex.RLock()
	defer fake.removeReleaseDependencyMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
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

var _ releasedependency.PivnetClient = new(FakePivnetClient)