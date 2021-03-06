// This file was generated by counterfeiter
package fakes

import "sync"

type Bar struct {
	SetTotalStub        func(contentLength int64)
	setTotalMutex       sync.RWMutex
	setTotalArgsForCall []struct {
		contentLength int64
	}
	AddStub        func(totalWritten int) int
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		totalWritten int
	}
	addReturns struct {
		result1 int
	}
	KickoffStub        func()
	kickoffMutex       sync.RWMutex
	kickoffArgsForCall []struct{}
	FinishStub         func()
	finishMutex        sync.RWMutex
	finishArgsForCall  []struct{}
}

func (fake *Bar) SetTotal(contentLength int64) {
	fake.setTotalMutex.Lock()
	fake.setTotalArgsForCall = append(fake.setTotalArgsForCall, struct {
		contentLength int64
	}{contentLength})
	fake.setTotalMutex.Unlock()
	if fake.SetTotalStub != nil {
		fake.SetTotalStub(contentLength)
	}
}

func (fake *Bar) SetTotalCallCount() int {
	fake.setTotalMutex.RLock()
	defer fake.setTotalMutex.RUnlock()
	return len(fake.setTotalArgsForCall)
}

func (fake *Bar) SetTotalArgsForCall(i int) int64 {
	fake.setTotalMutex.RLock()
	defer fake.setTotalMutex.RUnlock()
	return fake.setTotalArgsForCall[i].contentLength
}

func (fake *Bar) Add(totalWritten int) int {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		totalWritten int
	}{totalWritten})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(totalWritten)
	} else {
		return fake.addReturns.result1
	}
}

func (fake *Bar) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *Bar) AddArgsForCall(i int) int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].totalWritten
}

func (fake *Bar) AddReturns(result1 int) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 int
	}{result1}
}

func (fake *Bar) Kickoff() {
	fake.kickoffMutex.Lock()
	fake.kickoffArgsForCall = append(fake.kickoffArgsForCall, struct{}{})
	fake.kickoffMutex.Unlock()
	if fake.KickoffStub != nil {
		fake.KickoffStub()
	}
}

func (fake *Bar) KickoffCallCount() int {
	fake.kickoffMutex.RLock()
	defer fake.kickoffMutex.RUnlock()
	return len(fake.kickoffArgsForCall)
}

func (fake *Bar) Finish() {
	fake.finishMutex.Lock()
	fake.finishArgsForCall = append(fake.finishArgsForCall, struct{}{})
	fake.finishMutex.Unlock()
	if fake.FinishStub != nil {
		fake.FinishStub()
	}
}

func (fake *Bar) FinishCallCount() int {
	fake.finishMutex.RLock()
	defer fake.finishMutex.RUnlock()
	return len(fake.finishArgsForCall)
}
