// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"database/sql"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
)

type FakeUnitOfWork struct {
	BeginTxStub        func(context.Context, *sql.TxOptions) (*sql.Tx, error)
	beginTxMutex       sync.RWMutex
	beginTxArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.TxOptions
	}
	beginTxReturns struct {
		result1 *sql.Tx
		result2 error
	}
	beginTxReturnsOnCall map[int]struct {
		result1 *sql.Tx
		result2 error
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
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnitOfWork) BeginTx(arg1 context.Context, arg2 *sql.TxOptions) (*sql.Tx, error) {
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
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUnitOfWork) BeginTxCallCount() int {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	return len(fake.beginTxArgsForCall)
}

func (fake *FakeUnitOfWork) BeginTxCalls(stub func(context.Context, *sql.TxOptions) (*sql.Tx, error)) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = stub
}

func (fake *FakeUnitOfWork) BeginTxArgsForCall(i int) (context.Context, *sql.TxOptions) {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	argsForCall := fake.beginTxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnitOfWork) BeginTxReturns(result1 *sql.Tx, result2 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	fake.beginTxReturns = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeUnitOfWork) BeginTxReturnsOnCall(i int, result1 *sql.Tx, result2 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	if fake.beginTxReturnsOnCall == nil {
		fake.beginTxReturnsOnCall = make(map[int]struct {
			result1 *sql.Tx
			result2 error
		})
	}
	fake.beginTxReturnsOnCall[i] = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeUnitOfWork) Commit() error {
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

func (fake *FakeUnitOfWork) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeUnitOfWork) CommitCalls(stub func() error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = stub
}

func (fake *FakeUnitOfWork) CommitReturns(result1 error) {
	fake.commitMutex.Lock()
	defer fake.commitMutex.Unlock()
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnitOfWork) CommitReturnsOnCall(i int, result1 error) {
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

func (fake *FakeUnitOfWork) Rollback() error {
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

func (fake *FakeUnitOfWork) RollbackCallCount() int {
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	return len(fake.rollbackArgsForCall)
}

func (fake *FakeUnitOfWork) RollbackCalls(stub func() error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = stub
}

func (fake *FakeUnitOfWork) RollbackReturns(result1 error) {
	fake.rollbackMutex.Lock()
	defer fake.rollbackMutex.Unlock()
	fake.RollbackStub = nil
	fake.rollbackReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeUnitOfWork) RollbackReturnsOnCall(i int, result1 error) {
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

func (fake *FakeUnitOfWork) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	fake.rollbackMutex.RLock()
	defer fake.rollbackMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnitOfWork) recordInvocation(key string, args []interface{}) {
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

var _ repository.UnitOfWork = new(FakeUnitOfWork)
