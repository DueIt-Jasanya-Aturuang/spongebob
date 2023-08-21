// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type FakeProfileUsecase struct {
	GetProfileByIdStub        func(context.Context, string) (*dto.ProfileResp, error)
	getProfileByIdMutex       sync.RWMutex
	getProfileByIdArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getProfileByIdReturns struct {
		result1 *dto.ProfileResp
		result2 error
	}
	getProfileByIdReturnsOnCall map[int]struct {
		result1 *dto.ProfileResp
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProfileUsecase) GetProfileById(arg1 context.Context, arg2 string) (*dto.ProfileResp, error) {
	fake.getProfileByIdMutex.Lock()
	ret, specificReturn := fake.getProfileByIdReturnsOnCall[len(fake.getProfileByIdArgsForCall)]
	fake.getProfileByIdArgsForCall = append(fake.getProfileByIdArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetProfileByIdStub
	fakeReturns := fake.getProfileByIdReturns
	fake.recordInvocation("GetProfileById", []interface{}{arg1, arg2})
	fake.getProfileByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileUsecase) GetProfileByIdCallCount() int {
	fake.getProfileByIdMutex.RLock()
	defer fake.getProfileByIdMutex.RUnlock()
	return len(fake.getProfileByIdArgsForCall)
}

func (fake *FakeProfileUsecase) GetProfileByIdCalls(stub func(context.Context, string) (*dto.ProfileResp, error)) {
	fake.getProfileByIdMutex.Lock()
	defer fake.getProfileByIdMutex.Unlock()
	fake.GetProfileByIdStub = stub
}

func (fake *FakeProfileUsecase) GetProfileByIdArgsForCall(i int) (context.Context, string) {
	fake.getProfileByIdMutex.RLock()
	defer fake.getProfileByIdMutex.RUnlock()
	argsForCall := fake.getProfileByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileUsecase) GetProfileByIdReturns(result1 *dto.ProfileResp, result2 error) {
	fake.getProfileByIdMutex.Lock()
	defer fake.getProfileByIdMutex.Unlock()
	fake.GetProfileByIdStub = nil
	fake.getProfileByIdReturns = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) GetProfileByIdReturnsOnCall(i int, result1 *dto.ProfileResp, result2 error) {
	fake.getProfileByIdMutex.Lock()
	defer fake.getProfileByIdMutex.Unlock()
	fake.GetProfileByIdStub = nil
	if fake.getProfileByIdReturnsOnCall == nil {
		fake.getProfileByIdReturnsOnCall = make(map[int]struct {
			result1 *dto.ProfileResp
			result2 error
		})
	}
	fake.getProfileByIdReturnsOnCall[i] = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProfileByIdMutex.RLock()
	defer fake.getProfileByIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProfileUsecase) recordInvocation(key string, args []interface{}) {
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

var _ usecase.ProfileUsecase = new(FakeProfileUsecase)
