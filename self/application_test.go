package self

import (
	"testing"

	selfSym "bitbucket.org/taubyte/go-sdk-symbols/self"
	"bitbucket.org/taubyte/go-sdk/errno"
)

func TestApplicationError(t *testing.T) {
	selfSym.ApplicationSize = func(SizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Application()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.ApplicationSize = func(SizePtr *uint32) errno.Error {
		*SizePtr = 5
		return 0
	}

	selfSym.Application = func(Ptr *byte) errno.Error {
		return 1
	}

	_, err = Application()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
