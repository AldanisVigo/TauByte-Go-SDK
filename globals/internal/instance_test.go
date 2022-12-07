package internal_test

import (
	"testing"

	globalSym "bitbucket.org/taubyte/go-sdk-symbols/globals"
	"bitbucket.org/taubyte/go-sdk/errno"
	"bitbucket.org/taubyte/go-sdk/globals/str"
)

func TestInstanceError(t *testing.T) {
	globalSym.MockData{}.Mock()

	_, err := str.Get("name")
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	globalSym.GetOrCreateGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		return 1
	}

	_, err = str.GetOrCreate("name")
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	globalSym.GetOrCreateGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		return 0
	}

	_, err = str.GetOrCreate("name")
	if err != nil {
		t.Error("Expected error, got nil")
		return
	}

	globalSym.GetOrCreateGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		*valueSizePtr = uint32(6)
		return 0
	}

	_, err = str.GetOrCreate("name")
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}

	globalSym.GetGlobalValueSize = func(name string, application, function uint32, valueSizePtr *uint32) errno.Error {
		*valueSizePtr = uint32(6)
		return 0
	}

	_, err = str.Get("name")
	if err == nil {
		t.Error("Expected error, got nil")
		return
	}
}
