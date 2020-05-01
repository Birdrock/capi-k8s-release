// Code generated by counterfeiter. DO NOT EDIT.
package image_registryfakes

import (
	"capi_kpack_watcher/image_registry"
	"sync"
)

type FakeManifestFetcher struct {
	FetchRawManifestFromImageReferenceStub        func(string) ([]byte, error)
	fetchRawManifestFromImageReferenceMutex       sync.RWMutex
	fetchRawManifestFromImageReferenceArgsForCall []struct {
		arg1 string
	}
	fetchRawManifestFromImageReferenceReturns struct {
		result1 []byte
		result2 error
	}
	fetchRawManifestFromImageReferenceReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReference(arg1 string) ([]byte, error) {
	fake.fetchRawManifestFromImageReferenceMutex.Lock()
	ret, specificReturn := fake.fetchRawManifestFromImageReferenceReturnsOnCall[len(fake.fetchRawManifestFromImageReferenceArgsForCall)]
	fake.fetchRawManifestFromImageReferenceArgsForCall = append(fake.fetchRawManifestFromImageReferenceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("FetchRawManifestFromImageReference", []interface{}{arg1})
	fake.fetchRawManifestFromImageReferenceMutex.Unlock()
	if fake.FetchRawManifestFromImageReferenceStub != nil {
		return fake.FetchRawManifestFromImageReferenceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchRawManifestFromImageReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReferenceCallCount() int {
	fake.fetchRawManifestFromImageReferenceMutex.RLock()
	defer fake.fetchRawManifestFromImageReferenceMutex.RUnlock()
	return len(fake.fetchRawManifestFromImageReferenceArgsForCall)
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReferenceCalls(stub func(string) ([]byte, error)) {
	fake.fetchRawManifestFromImageReferenceMutex.Lock()
	defer fake.fetchRawManifestFromImageReferenceMutex.Unlock()
	fake.FetchRawManifestFromImageReferenceStub = stub
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReferenceArgsForCall(i int) string {
	fake.fetchRawManifestFromImageReferenceMutex.RLock()
	defer fake.fetchRawManifestFromImageReferenceMutex.RUnlock()
	argsForCall := fake.fetchRawManifestFromImageReferenceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReferenceReturns(result1 []byte, result2 error) {
	fake.fetchRawManifestFromImageReferenceMutex.Lock()
	defer fake.fetchRawManifestFromImageReferenceMutex.Unlock()
	fake.FetchRawManifestFromImageReferenceStub = nil
	fake.fetchRawManifestFromImageReferenceReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestFetcher) FetchRawManifestFromImageReferenceReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.fetchRawManifestFromImageReferenceMutex.Lock()
	defer fake.fetchRawManifestFromImageReferenceMutex.Unlock()
	fake.FetchRawManifestFromImageReferenceStub = nil
	if fake.fetchRawManifestFromImageReferenceReturnsOnCall == nil {
		fake.fetchRawManifestFromImageReferenceReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.fetchRawManifestFromImageReferenceReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeManifestFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchRawManifestFromImageReferenceMutex.RLock()
	defer fake.fetchRawManifestFromImageReferenceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManifestFetcher) recordInvocation(key string, args []interface{}) {
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

var _ image_registry.ManifestFetcher = new(FakeManifestFetcher)