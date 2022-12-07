package self

import (
	"testing"

	selfSym "bitbucket.org/taubyte/go-sdk-symbols/self"
	"bitbucket.org/taubyte/go-sdk/errno"
)

func TestCommitError(t *testing.T) {
	selfSym.CommitSize = func(sizePtr *uint32) errno.Error {
		return 1
	}

	_, err := Commit()
	if err == nil {
		t.Error("Expected error")
		return
	}

	selfSym.CommitSize = func(sizePtr *uint32) errno.Error {
		*sizePtr = 5
		return 0
	}

	selfSym.Commit = func(ptr *byte) errno.Error {
		return 1
	}

	_, err = Commit()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
