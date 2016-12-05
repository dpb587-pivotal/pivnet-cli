// This file was generated by counterfeiter
package logoutfakes

import (
	"sync"

	"github.com/pivotal-cf/pivnet-cli/commands/logout"
)

type FakeRCHandler struct {
	RemoveProfileWithNameStub        func(profileName string) error
	removeProfileWithNameMutex       sync.RWMutex
	removeProfileWithNameArgsForCall []struct {
		profileName string
	}
	removeProfileWithNameReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRCHandler) RemoveProfileWithName(profileName string) error {
	fake.removeProfileWithNameMutex.Lock()
	fake.removeProfileWithNameArgsForCall = append(fake.removeProfileWithNameArgsForCall, struct {
		profileName string
	}{profileName})
	fake.recordInvocation("RemoveProfileWithName", []interface{}{profileName})
	fake.removeProfileWithNameMutex.Unlock()
	if fake.RemoveProfileWithNameStub != nil {
		return fake.RemoveProfileWithNameStub(profileName)
	} else {
		return fake.removeProfileWithNameReturns.result1
	}
}

func (fake *FakeRCHandler) RemoveProfileWithNameCallCount() int {
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return len(fake.removeProfileWithNameArgsForCall)
}

func (fake *FakeRCHandler) RemoveProfileWithNameArgsForCall(i int) string {
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return fake.removeProfileWithNameArgsForCall[i].profileName
}

func (fake *FakeRCHandler) RemoveProfileWithNameReturns(result1 error) {
	fake.RemoveProfileWithNameStub = nil
	fake.removeProfileWithNameReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRCHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.removeProfileWithNameMutex.RLock()
	defer fake.removeProfileWithNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRCHandler) recordInvocation(key string, args []interface{}) {
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

var _ logout.RCHandler = new(FakeRCHandler)
