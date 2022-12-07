package pubsub_test

import (
	"fmt"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/pubsub"
	"bitbucket.org/taubyte/go-sdk/pubsub"
)

func ExampleChannelObject_Subscribe() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		Channel: "someChannel",
	}.Mock()

	channel, err := pubsub.Channel("someChannel")
	if err != nil {
		return
	}

	err = channel.Subscribe()
	if err != nil {
		return
	}

	fmt.Println("Subs:", m.Subscriptions)
	// Output: Subs: [someChannel]
}
