// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"mime/multipart"
	"sync"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/domain/repository"
)

type FakeMinioRepo struct {
	DeleteFileStub        func(context.Context, string, string) error
	deleteFileMutex       sync.RWMutex
	deleteFileArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	deleteFileReturns struct {
		result1 error
	}
	deleteFileReturnsOnCall map[int]struct {
		result1 error
	}
	GenerateFileNameStub        func(*multipart.FileHeader, string, string) string
	generateFileNameMutex       sync.RWMutex
	generateFileNameArgsForCall []struct {
		arg1 *multipart.FileHeader
		arg2 string
		arg3 string
	}
	generateFileNameReturns struct {
		result1 string
	}
	generateFileNameReturnsOnCall map[int]struct {
		result1 string
	}
	UploadFileStub        func(context.Context, *multipart.FileHeader, string, string) error
	uploadFileMutex       sync.RWMutex
	uploadFileArgsForCall []struct {
		arg1 context.Context
		arg2 *multipart.FileHeader
		arg3 string
		arg4 string
	}
	uploadFileReturns struct {
		result1 error
	}
	uploadFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMinioRepo) DeleteFile(arg1 context.Context, arg2 string, arg3 string) error {
	fake.deleteFileMutex.Lock()
	ret, specificReturn := fake.deleteFileReturnsOnCall[len(fake.deleteFileArgsForCall)]
	fake.deleteFileArgsForCall = append(fake.deleteFileArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeleteFileStub
	fakeReturns := fake.deleteFileReturns
	fake.recordInvocation("DeleteFile", []interface{}{arg1, arg2, arg3})
	fake.deleteFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMinioRepo) DeleteFileCallCount() int {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	return len(fake.deleteFileArgsForCall)
}

func (fake *FakeMinioRepo) DeleteFileCalls(stub func(context.Context, string, string) error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = stub
}

func (fake *FakeMinioRepo) DeleteFileArgsForCall(i int) (context.Context, string, string) {
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	argsForCall := fake.deleteFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMinioRepo) DeleteFileReturns(result1 error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = nil
	fake.deleteFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMinioRepo) DeleteFileReturnsOnCall(i int, result1 error) {
	fake.deleteFileMutex.Lock()
	defer fake.deleteFileMutex.Unlock()
	fake.DeleteFileStub = nil
	if fake.deleteFileReturnsOnCall == nil {
		fake.deleteFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMinioRepo) GenerateFileName(arg1 *multipart.FileHeader, arg2 string, arg3 string) string {
	fake.generateFileNameMutex.Lock()
	ret, specificReturn := fake.generateFileNameReturnsOnCall[len(fake.generateFileNameArgsForCall)]
	fake.generateFileNameArgsForCall = append(fake.generateFileNameArgsForCall, struct {
		arg1 *multipart.FileHeader
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GenerateFileNameStub
	fakeReturns := fake.generateFileNameReturns
	fake.recordInvocation("GenerateFileName", []interface{}{arg1, arg2, arg3})
	fake.generateFileNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMinioRepo) GenerateFileNameCallCount() int {
	fake.generateFileNameMutex.RLock()
	defer fake.generateFileNameMutex.RUnlock()
	return len(fake.generateFileNameArgsForCall)
}

func (fake *FakeMinioRepo) GenerateFileNameCalls(stub func(*multipart.FileHeader, string, string) string) {
	fake.generateFileNameMutex.Lock()
	defer fake.generateFileNameMutex.Unlock()
	fake.GenerateFileNameStub = stub
}

func (fake *FakeMinioRepo) GenerateFileNameArgsForCall(i int) (*multipart.FileHeader, string, string) {
	fake.generateFileNameMutex.RLock()
	defer fake.generateFileNameMutex.RUnlock()
	argsForCall := fake.generateFileNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMinioRepo) GenerateFileNameReturns(result1 string) {
	fake.generateFileNameMutex.Lock()
	defer fake.generateFileNameMutex.Unlock()
	fake.GenerateFileNameStub = nil
	fake.generateFileNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMinioRepo) GenerateFileNameReturnsOnCall(i int, result1 string) {
	fake.generateFileNameMutex.Lock()
	defer fake.generateFileNameMutex.Unlock()
	fake.GenerateFileNameStub = nil
	if fake.generateFileNameReturnsOnCall == nil {
		fake.generateFileNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.generateFileNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeMinioRepo) UploadFile(arg1 context.Context, arg2 *multipart.FileHeader, arg3 string, arg4 string) error {
	fake.uploadFileMutex.Lock()
	ret, specificReturn := fake.uploadFileReturnsOnCall[len(fake.uploadFileArgsForCall)]
	fake.uploadFileArgsForCall = append(fake.uploadFileArgsForCall, struct {
		arg1 context.Context
		arg2 *multipart.FileHeader
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.UploadFileStub
	fakeReturns := fake.uploadFileReturns
	fake.recordInvocation("UploadFile", []interface{}{arg1, arg2, arg3, arg4})
	fake.uploadFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMinioRepo) UploadFileCallCount() int {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	return len(fake.uploadFileArgsForCall)
}

func (fake *FakeMinioRepo) UploadFileCalls(stub func(context.Context, *multipart.FileHeader, string, string) error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = stub
}

func (fake *FakeMinioRepo) UploadFileArgsForCall(i int) (context.Context, *multipart.FileHeader, string, string) {
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	argsForCall := fake.uploadFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeMinioRepo) UploadFileReturns(result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	fake.uploadFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMinioRepo) UploadFileReturnsOnCall(i int, result1 error) {
	fake.uploadFileMutex.Lock()
	defer fake.uploadFileMutex.Unlock()
	fake.UploadFileStub = nil
	if fake.uploadFileReturnsOnCall == nil {
		fake.uploadFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMinioRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteFileMutex.RLock()
	defer fake.deleteFileMutex.RUnlock()
	fake.generateFileNameMutex.RLock()
	defer fake.generateFileNameMutex.RUnlock()
	fake.uploadFileMutex.RLock()
	defer fake.uploadFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMinioRepo) recordInvocation(key string, args []interface{}) {
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

var _ repository.MinioRepo = new(FakeMinioRepo)
