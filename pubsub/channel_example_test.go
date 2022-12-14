package pubsub_test

import (
	"fmt"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/pubsub"
	"bitbucket.org/taubyte/go-sdk/pubsub"
)

func ExampleChannel() {
	// Mocking the calls to the vm for usage in tests and playground
	m := symbols.MockData{
		Channel:      "someChannel",
		WebSocketURL: "ws-QmQq71tkq1yKmYobFUhPWF2MejK5CrvpY4h7HQDDT8b8Zb/someChannel",
	}.Mock()

	channel, err := pubsub.Channel("someChannel")
	if err != nil {
		return
	}

	fmt.Println("Name:", channel.Name())

	err = channel.Subscribe()
	if err != nil {
		return
	}

	fmt.Println("Subscriptions:", m.Subscriptions)

	err = channel.Publish([]byte("Hello, world!"))
	if err != nil {
		return
	}

	fmt.Println("Published:", string(m.PublishedData))

	url, err := channel.WebSocket().Url()
	if err != nil {
		return
	}

	fmt.Println("Url:", url.Path)

	// Output: Name: someChannel
	// Subscriptions: [someChannel]
	// Published: Hello, world!
	// Url: ws-QmQq71tkq1yKmYobFUhPWF2MejK5CrvpY4h7HQDDT8b8Zb/someChannel
}
