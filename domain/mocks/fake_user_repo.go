// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"database/sql"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain"
)

type FakeUserRepo struct {
	GetUserByIdStub        func(context.Context, *sql.DB, string) (*domain.User, error)
	getUserByIdMutex       sync.RWMutex
	getUserByIdArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.DB
		arg3 string
	}
	getUserByIdReturns struct {
		result1 *domain.User
		result2 error
	}
	getUserByIdReturnsOnCall map[int]struct {
		result1 *domain.User
		result2 error
	}
	UpdateUserStub        func(context.Context, *sql.Tx, domain.User) (*domain.User, error)
	updateUserMutex       sync.RWMutex
	updateUserArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.Tx
		arg3 domain.User
	}
	updateUserReturns struct {
		result1 *domain.User
		result2 error
	}
	updateUserReturnsOnCall map[int]struct {
		result1 *domain.User
		result2 error
	}
	UpdateUsernameStub        func(context.Context, *sql.Tx, domain.User) (*domain.User, error)
	updateUsernameMutex       sync.RWMutex
	updateUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.Tx
		arg3 domain.User
	}
	updateUsernameReturns struct {
		result1 *domain.User
		result2 error
	}
	updateUsernameReturnsOnCall map[int]struct {
		result1 *domain.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserRepo) GetUserById(arg1 context.Context, arg2 *sql.DB, arg3 string) (*domain.User, error) {
	fake.getUserByIdMutex.Lock()
	ret, specificReturn := fake.getUserByIdReturnsOnCall[len(fake.getUserByIdArgsForCall)]
	fake.getUserByIdArgsForCall = append(fake.getUserByIdArgsForCall, struct {
		arg1 context.Context
		arg2 *sql.DB
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetUserByIdStub
	fakeReturns := fake.getUserByIdReturns
	fake.recordInvocation("GetUserById", []interface{}{arg1, arg2, arg3})
	fake.getUserByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) GetUserByIdCallCount() int {
	fake.getUserByIdMutex.RLock()
	defer fake.getUserByIdMutex.RUnlock()
	return len(fake.getUserByIdArgsForCall)
}

func (fake *FakeUserRepo) GetUserByIdCalls(stub func(context.Context, *sql.DB, string) (*domain.User, error)) {
	fake.getUserByIdMutex.Lock()
	defer fake.getUserByIdMutex.Unlock()
	fake.GetUserByIdStub = stub
}

func (fake *FakeUserRepo) GetUserByIdArgsForCall(i int) (context.Context, *sql.DB, string) {
	fake.getUserByIdMutex.RLock()
	defer fake.getUserByIdMutex.RUnlock()
	argsForCall := fake.getUserByIdArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserRepo) GetUserByIdReturns(result1 *domain.User, result2 error) {
	fake.getUserByIdMutex.Lock()
	defer fake.getUserByIdMutex.Unlock()
	fake.GetUserByIdStub = nil
	fake.getUserByIdReturns = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) GetUserByIdReturnsOnCall(i int, result1 *domain.User, result2 error) {
	fake.getUserByIdMutex.Lock()
	defer fake.getUserByIdMutex.Unlock()
	fake.GetUserByIdStub = nil
	if fake.getUserByIdReturnsOnCall == nil {
		fake.getUserByIdReturnsOnCall = make(map[int]struct {
			result1 *domain.User
			result2 error
		})
	}
	fake.getUserByIdReturnsOnCall[i] = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUser(arg1 context.Context, arg2 *sql.Tx, arg3 domain.User) (*domain.User, error) {
	fake.updateUserMutex.Lock()
	ret, specificReturn := fake.updateUserReturnsOnCall[len(fake.updateUserArgsForCall)]
	fake.updateUserArgsForCall = append(fake.updateUserArgsForCall, struct {
		arg1 context.Context
		arg2 *sql.Tx
		arg3 domain.User
	}{arg1, arg2, arg3})
	stub := fake.UpdateUserStub
	fakeReturns := fake.updateUserReturns
	fake.recordInvocation("UpdateUser", []interface{}{arg1, arg2, arg3})
	fake.updateUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) UpdateUserCallCount() int {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	return len(fake.updateUserArgsForCall)
}

func (fake *FakeUserRepo) UpdateUserCalls(stub func(context.Context, *sql.Tx, domain.User) (*domain.User, error)) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = stub
}

func (fake *FakeUserRepo) UpdateUserArgsForCall(i int) (context.Context, *sql.Tx, domain.User) {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	argsForCall := fake.updateUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserRepo) UpdateUserReturns(result1 *domain.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	fake.updateUserReturns = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUserReturnsOnCall(i int, result1 *domain.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	if fake.updateUserReturnsOnCall == nil {
		fake.updateUserReturnsOnCall = make(map[int]struct {
			result1 *domain.User
			result2 error
		})
	}
	fake.updateUserReturnsOnCall[i] = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUsername(arg1 context.Context, arg2 *sql.Tx, arg3 domain.User) (*domain.User, error) {
	fake.updateUsernameMutex.Lock()
	ret, specificReturn := fake.updateUsernameReturnsOnCall[len(fake.updateUsernameArgsForCall)]
	fake.updateUsernameArgsForCall = append(fake.updateUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 *sql.Tx
		arg3 domain.User
	}{arg1, arg2, arg3})
	stub := fake.UpdateUsernameStub
	fakeReturns := fake.updateUsernameReturns
	fake.recordInvocation("UpdateUsername", []interface{}{arg1, arg2, arg3})
	fake.updateUsernameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) UpdateUsernameCallCount() int {
	fake.updateUsernameMutex.RLock()
	defer fake.updateUsernameMutex.RUnlock()
	return len(fake.updateUsernameArgsForCall)
}

func (fake *FakeUserRepo) UpdateUsernameCalls(stub func(context.Context, *sql.Tx, domain.User) (*domain.User, error)) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = stub
}

func (fake *FakeUserRepo) UpdateUsernameArgsForCall(i int) (context.Context, *sql.Tx, domain.User) {
	fake.updateUsernameMutex.RLock()
	defer fake.updateUsernameMutex.RUnlock()
	argsForCall := fake.updateUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeUserRepo) UpdateUsernameReturns(result1 *domain.User, result2 error) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = nil
	fake.updateUsernameReturns = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUsernameReturnsOnCall(i int, result1 *domain.User, result2 error) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = nil
	if fake.updateUsernameReturnsOnCall == nil {
		fake.updateUsernameReturnsOnCall = make(map[int]struct {
			result1 *domain.User
			result2 error
		})
	}
	fake.updateUsernameReturnsOnCall[i] = struct {
		result1 *domain.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getUserByIdMutex.RLock()
	defer fake.getUserByIdMutex.RUnlock()
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	fake.updateUsernameMutex.RLock()
	defer fake.updateUsernameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUserRepo) recordInvocation(key string, args []interface{}) {
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

var _ domain.UserRepo = new(FakeUserRepo)
