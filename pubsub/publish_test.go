package pubsub

import (
	"testing"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/pubsub"
)

func TestPublish(t *testing.T) {
	symbols.MockData{}.Mock()

	channel := &ChannelObject{name: "someChannel"}

	err := channel.Publish([]byte("Hello, world!"))
	if err == nil {
		t.Error("Expected error")
		return
	}
}
