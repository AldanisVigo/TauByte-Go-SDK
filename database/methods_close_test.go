package database

import (
	"testing"

	databaseSym "bitbucket.org/taubyte/go-sdk-symbols/database"
)

func TestCloseError(t *testing.T) {
	databaseSym.MockClose(5)

	err := Database(7).Close()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
