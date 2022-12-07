package pubsub

import (
	"testing"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/pubsub"
)

func TestChannel(t *testing.T) {
	symbols.MockData{Channel: "someChannel"}.Mock()

	_, err := Channel("")
	if err == nil {
		t.Error("Expected error")
		return
	}
}
