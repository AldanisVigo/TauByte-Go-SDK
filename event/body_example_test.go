package event_test

import (
	"fmt"
	"io/ioutil"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/event"
	"bitbucket.org/taubyte/go-sdk/common"
	"bitbucket.org/taubyte/go-sdk/event"
)

func ExampleHttpEvent_Body() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		EventType: common.EventTypeHttp,
		Body:      []byte("Hello, world!"),
	}.Mock()

	var e event.Event
	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	body := httpEvent.Body()
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	err = body.Close()
	if err != nil {
		return
	}

	fmt.Println(string(data))
	// Output: Hello, world!
}
