// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/dto"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/usecase"
)

type FakeProfileUsecase struct {
	GetProfileByIDStub        func(context.Context, *dto.GetProfileReq) (*dto.ProfileResp, error)
	getProfileByIDMutex       sync.RWMutex
	getProfileByIDArgsForCall []struct {
		arg1 context.Context
		arg2 *dto.GetProfileReq
	}
	getProfileByIDReturns struct {
		result1 *dto.ProfileResp
		result2 error
	}
	getProfileByIDReturnsOnCall map[int]struct {
		result1 *dto.ProfileResp
		result2 error
	}
	StoreProfileStub        func(context.Context, *dto.StoreProfileReq) (*dto.ProfileResp, error)
	storeProfileMutex       sync.RWMutex
	storeProfileArgsForCall []struct {
		arg1 context.Context
		arg2 *dto.StoreProfileReq
	}
	storeProfileReturns struct {
		result1 *dto.ProfileResp
		result2 error
	}
	storeProfileReturnsOnCall map[int]struct {
		result1 *dto.ProfileResp
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProfileUsecase) GetProfileByID(arg1 context.Context, arg2 *dto.GetProfileReq) (*dto.ProfileResp, error) {
	fake.getProfileByIDMutex.Lock()
	ret, specificReturn := fake.getProfileByIDReturnsOnCall[len(fake.getProfileByIDArgsForCall)]
	fake.getProfileByIDArgsForCall = append(fake.getProfileByIDArgsForCall, struct {
		arg1 context.Context
		arg2 *dto.GetProfileReq
	}{arg1, arg2})
	stub := fake.GetProfileByIDStub
	fakeReturns := fake.getProfileByIDReturns
	fake.recordInvocation("GetProfileByID", []interface{}{arg1, arg2})
	fake.getProfileByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileUsecase) GetProfileByIDCallCount() int {
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	return len(fake.getProfileByIDArgsForCall)
}

func (fake *FakeProfileUsecase) GetProfileByIDCalls(stub func(context.Context, *dto.GetProfileReq) (*dto.ProfileResp, error)) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = stub
}

func (fake *FakeProfileUsecase) GetProfileByIDArgsForCall(i int) (context.Context, *dto.GetProfileReq) {
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	argsForCall := fake.getProfileByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileUsecase) GetProfileByIDReturns(result1 *dto.ProfileResp, result2 error) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = nil
	fake.getProfileByIDReturns = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) GetProfileByIDReturnsOnCall(i int, result1 *dto.ProfileResp, result2 error) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = nil
	if fake.getProfileByIDReturnsOnCall == nil {
		fake.getProfileByIDReturnsOnCall = make(map[int]struct {
			result1 *dto.ProfileResp
			result2 error
		})
	}
	fake.getProfileByIDReturnsOnCall[i] = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) StoreProfile(arg1 context.Context, arg2 *dto.StoreProfileReq) (*dto.ProfileResp, error) {
	fake.storeProfileMutex.Lock()
	ret, specificReturn := fake.storeProfileReturnsOnCall[len(fake.storeProfileArgsForCall)]
	fake.storeProfileArgsForCall = append(fake.storeProfileArgsForCall, struct {
		arg1 context.Context
		arg2 *dto.StoreProfileReq
	}{arg1, arg2})
	stub := fake.StoreProfileStub
	fakeReturns := fake.storeProfileReturns
	fake.recordInvocation("StoreProfile", []interface{}{arg1, arg2})
	fake.storeProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileUsecase) StoreProfileCallCount() int {
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
	return len(fake.storeProfileArgsForCall)
}

func (fake *FakeProfileUsecase) StoreProfileCalls(stub func(context.Context, *dto.StoreProfileReq) (*dto.ProfileResp, error)) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = stub
}

func (fake *FakeProfileUsecase) StoreProfileArgsForCall(i int) (context.Context, *dto.StoreProfileReq) {
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
	argsForCall := fake.storeProfileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileUsecase) StoreProfileReturns(result1 *dto.ProfileResp, result2 error) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = nil
	fake.storeProfileReturns = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) StoreProfileReturnsOnCall(i int, result1 *dto.ProfileResp, result2 error) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = nil
	if fake.storeProfileReturnsOnCall == nil {
		fake.storeProfileReturnsOnCall = make(map[int]struct {
			result1 *dto.ProfileResp
			result2 error
		})
	}
	fake.storeProfileReturnsOnCall[i] = struct {
		result1 *dto.ProfileResp
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileUsecase) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
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
