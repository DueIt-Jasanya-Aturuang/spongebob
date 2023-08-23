// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"database/sql"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/model"
	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
)

type FakeUserRepo struct {
	BeginTxStub        func(context.Context, *sql.TxOptions) error
	beginTxMutex       sync.RWMutex
	beginTxArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.TxOptions
	}
	beginTxReturns struct {
		result1 error
	}
	beginTxReturnsOnCall map[int]struct {
		result1 error
	}
	CallTxStub        func(*sql.Tx) error
	callTxMutex       sync.RWMutex
	callTxArgsForCall []struct {
		arg1 *sql.Tx
	}
	callTxReturns struct {
		result1 error
	}
	callTxReturnsOnCall map[int]struct {
		result1 error
	}
	CommitStub        func() error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
	}
	commitReturns struct {
		result1 error
	}
	commitReturnsOnCall map[int]struct {
		result1 error
	}
	GetTxStub        func() *sql.Tx
	getTxMutex       sync.RWMutex
	getTxArgsForCall []struct {
	}
	getTxReturns struct {
		result1 *sql.Tx
	}
	getTxReturnsOnCall map[int]struct {
		result1 *sql.Tx
	}
	GetUserByIDStub        func(context.Context, string) (*model.User, error)
	getUserByIDMutex       sync.RWMutex
	getUserByIDArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getUserByIDReturns struct {
		result1 *model.User
		result2 error
	}
	getUserByIDReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	RollbackStub        func() error
	rollbackMutex       sync.RWMutex
	rollbackArgsForCall []struct {
	}
	rollbackReturns struct {
		result1 error
	}
	rollbackReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateUserStub        func(context.Context, model.User) (*model.User, error)
	updateUserMutex       sync.RWMutex
	updateUserArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
	}
	updateUserReturns struct {
		result1 *model.User
		result2 error
	}
	updateUserReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	UpdateUsernameStub        func(context.Context, model.User) (*model.User, error)
	updateUsernameMutex       sync.RWMutex
	updateUsernameArgsForCall []struct {
		arg1 context.Context
		arg2 model.User
	}
	updateUsernameReturns struct {
		result1 *model.User
		result2 error
	}
	updateUsernameReturnsOnCall map[int]struct {
		result1 *model.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserRepo) BeginTx(arg1 context.Context, arg2 *sql.TxOptions) error {
	fake.beginTxMutex.Lock()
	ret, specificReturn := fake.beginTxReturnsOnCall[len(fake.beginTxArgsForCall)]
	fake.beginTxArgsForCall = append(fake.beginTxArgsForCall, struct {
		arg1 context.Context
		arg2 *sql.TxOptions
	}{arg1, arg2})
	stub := fake.BeginTxStub
	fakeReturns := fake.beginTxReturns
	fake.recordInvocation("BeginTx", []interface{}{arg1, arg2})
	fake.beginTxMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) BeginTxCallCount() int {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	return len(fake.beginTxArgsForCall)
}

func (fake *FakeUserRepo) BeginTxCalls(stub func(context.Context, *sql.TxOptions) error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = stub
}

func (fake *FakeUserRepo) BeginTxArgsForCall(i int) (context.Context, *sql.TxOptions) {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	argsForCall := fake.beginTxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) BeginTxReturns(result1 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	fake.beginTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) BeginTxReturnsOnCall(i int, result1 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	if fake.beginTxReturnsOnCall == nil {
		fake.beginTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.beginTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) CallTx(arg1 *sql.Tx) error {
	fake.callTxMutex.Lock()
	ret, specificReturn := fake.callTxReturnsOnCall[len(fake.callTxArgsForCall)]
	fake.callTxArgsForCall = append(fake.callTxArgsForCall, struct {
		arg1 *sql.Tx
	}{arg1})
	stub := fake.CallTxStub
	fakeReturns := fake.callTxReturns
	fake.recordInvocation("CallTx", []interface{}{arg1})
	fake.callTxMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) CallTxCallCount() int {
	fake.callTxMutex.RLock()
	defer fake.callTxMutex.RUnlock()
	return len(fake.callTxArgsForCall)
}

func (fake *FakeUserRepo) CallTxCalls(stub func(*sql.Tx) error) {
	fake.callTxMutex.Lock()
	defer fake.callTxMutex.Unlock()
	fake.CallTxStub = stub
}

func (fake *FakeUserRepo) CallTxArgsForCall(i int) *sql.Tx {
	fake.callTxMutex.RLock()
	defer fake.callTxMutex.RUnlock()
	argsForCall := fake.callTxArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUserRepo) CallTxReturns(result1 error) {
	fake.callTxMutex.Lock()
	defer fake.callTxMutex.Unlock()
	fake.CallTxStub = nil
	fake.callTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) CallTxReturnsOnCall(i int, result1 error) {
	fake.callTxMutex.Lock()
	defer fake.callTxMutex.Unlock()
	fake.CallTxStub = nil
	if fake.callTxReturnsOnCall == nil {
		fake.callTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) Commit() error {
	fake.commitMutex.Lock()
	ret, specificReturn := fake.commitReturnsOnCall[len(fake.commitArgsForCall)]
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
	}{})
	stub := fake.CommitStub
	fakeReturns := fake.commitReturns
	fake.recordInvocation("Commit", []interface{}{})
	fake.commitMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeUserRepo) CommitCalls(stub func() error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeUserRepo) CommitReturns(result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) CommitReturnsOnCall(i int, result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	if fake.commitReturnsOnCall == nil {
		fake.commitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.commitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) GetTx() *sql.Tx {
	fake.getTxMutex.Lock()
	ret, specificReturn := fake.getTxReturnsOnCall[len(fake.getTxArgsForCall)]
	fake.getTxArgsForCall = append(fake.getTxArgsForCall, struct {
	}{})
	stub := fake.GetTxStub
	fakeReturns := fake.getTxReturns
	fake.recordInvocation("GetTx", []interface{}{})
	fake.getTxMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) GetTxCallCount() int {
	fake.getTxMutex.RLock()
	defer fake.getTxMutex.RUnlock()
	return len(fake.getTxArgsForCall)
}

func (fake *FakeUserRepo) GetTxCalls(stub func() *sql.Tx) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = stub
}

func (fake *FakeUserRepo) GetTxReturns(result1 *sql.Tx) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = nil
	fake.getTxReturns = struct {
		result1 *sql.Tx
	}{result1}
}

func (fake *FakeUserRepo) GetTxReturnsOnCall(i int, result1 *sql.Tx) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = nil
	if fake.getTxReturnsOnCall == nil {
		fake.getTxReturnsOnCall = make(map[int]struct {
			result1 *sql.Tx
		})
	}
	fake.getTxReturnsOnCall[i] = struct {
		result1 *sql.Tx
	}{result1}
}

func (fake *FakeUserRepo) GetUserByID(arg1 context.Context, arg2 string) (*model.User, error) {
	fake.getUserByIDMutex.Lock()
	ret, specificReturn := fake.getUserByIDReturnsOnCall[len(fake.getUserByIDArgsForCall)]
	fake.getUserByIDArgsForCall = append(fake.getUserByIDArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetUserByIDStub
	fakeReturns := fake.getUserByIDReturns
	fake.recordInvocation("GetUserByID", []interface{}{arg1, arg2})
	fake.getUserByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUserRepo) GetUserByIDCallCount() int {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	return len(fake.getUserByIDArgsForCall)
}

func (fake *FakeUserRepo) GetUserByIDCalls(stub func(context.Context, string) (*model.User, error)) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = stub
}

func (fake *FakeUserRepo) GetUserByIDArgsForCall(i int) (context.Context, string) {
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	argsForCall := fake.getUserByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) GetUserByIDReturns(result1 *model.User, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	fake.getUserByIDReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) GetUserByIDReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.getUserByIDMutex.Lock()
	defer fake.getUserByIDMutex.Unlock()
	fake.GetUserByIDStub = nil
	if fake.getUserByIDReturnsOnCall == nil {
		fake.getUserByIDReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.getUserByIDReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) Rollback() error {
	fake.rollbackMutex.Lock()
	ret, specificReturn := fake.rollbackReturnsOnCall[len(fake.rollbackArgsForCall)]
	fake.rollbackArgsForCall = append(fake.rollbackArgsForCall, struct {
	}{})
	stub := fake.RollbackStub
	fakeReturns := fake.rollbackReturns
	fake.recordInvocation("Rollback", []interface{}{})
	fake.rollbackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeUserRepo) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeUserRepo) RollbackCalls(stub func() error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *FakeUserRepo) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) RollbackReturnsOnCall(i int, result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	if fake.rollbackReturnsOnCall == nil {
		fake.rollbackReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.rollbackReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeUserRepo) UpdateUser(arg1 context.Context, arg2 model.User) (*model.User, error) {
	fake.updateUserMutex.Lock()
	ret, specificReturn := fake.updateUserReturnsOnCall[len(fake.updateUserArgsForCall)]
	fake.updateUserArgsForCall = append(fake.updateUserArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
	}{arg1, arg2})
	stub := fake.UpdateUserStub
	fakeReturns := fake.updateUserReturns
	fake.recordInvocation("UpdateUser", []interface{}{arg1, arg2})
	fake.updateUserMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
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

func (fake *FakeUserRepo) UpdateUserCalls(stub func(context.Context, model.User) (*model.User, error)) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = stub
}

func (fake *FakeUserRepo) UpdateUserArgsForCall(i int) (context.Context, model.User) {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	argsForCall := fake.updateUserArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) UpdateUserReturns(result1 *model.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	fake.updateUserReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUserReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.updateUserMutex.Lock()
	defer fake.updateUserMutex.Unlock()
	fake.UpdateUserStub = nil
	if fake.updateUserReturnsOnCall == nil {
		fake.updateUserReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.updateUserReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUsername(arg1 context.Context, arg2 model.User) (*model.User, error) {
	fake.updateUsernameMutex.Lock()
	ret, specificReturn := fake.updateUsernameReturnsOnCall[len(fake.updateUsernameArgsForCall)]
	fake.updateUsernameArgsForCall = append(fake.updateUsernameArgsForCall, struct {
		arg1 context.Context
		arg2 model.User
	}{arg1, arg2})
	stub := fake.UpdateUsernameStub
	fakeReturns := fake.updateUsernameReturns
	fake.recordInvocation("UpdateUsername", []interface{}{arg1, arg2})
	fake.updateUsernameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
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

func (fake *FakeUserRepo) UpdateUsernameCalls(stub func(context.Context, model.User) (*model.User, error)) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = stub
}

func (fake *FakeUserRepo) UpdateUsernameArgsForCall(i int) (context.Context, model.User) {
	fake.updateUsernameMutex.RLock()
	defer fake.updateUsernameMutex.RUnlock()
	argsForCall := fake.updateUsernameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUserRepo) UpdateUsernameReturns(result1 *model.User, result2 error) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = nil
	fake.updateUsernameReturns = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) UpdateUsernameReturnsOnCall(i int, result1 *model.User, result2 error) {
	fake.updateUsernameMutex.Lock()
	defer fake.updateUsernameMutex.Unlock()
	fake.UpdateUsernameStub = nil
	if fake.updateUsernameReturnsOnCall == nil {
		fake.updateUsernameReturnsOnCall = make(map[int]struct {
			result1 *model.User
			result2 error
		})
	}
	fake.updateUsernameReturnsOnCall[i] = struct {
		result1 *model.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUserRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	fake.callTxMutex.RLock()
	defer fake.callTxMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.getTxMutex.RLock()
	defer fake.getTxMutex.RUnlock()
	fake.getUserByIDMutex.RLock()
	defer fake.getUserByIDMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
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

var _ repository.UserRepo = new(FakeUserRepo)
