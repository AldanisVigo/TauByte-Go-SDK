package event_test

import (
	"fmt"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/event"
	"bitbucket.org/taubyte/go-sdk/common"
	"bitbucket.org/taubyte/go-sdk/event"
)

func ExampleHttpEvent_Headers() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockData{
		EventType: common.EventTypeHttp,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}.Mock()

	var e event.Event
	httpEvent, err := e.HTTP()
	if err != nil {
		return
	}

	headers := httpEvent.Headers()

	contentType, err := headers.Get("Content-Type")
	if err != nil {
		return
	}
	fmt.Println("Received content type:", contentType)

	headersList, err := headers.List()
	if err != nil {
		return
	}
	fmt.Println("Headers:", headersList)

	err = headers.Set("Content-Type", "text/csv")
	if err != nil {
		return
	}

	contentType, err = headers.Get("Content-Type")
	if err != nil {
		return
	}
	fmt.Println("New content type:", contentType)

	// Output: Received content type: application/json
	// Headers: [Content-Type]
	// New content type: text/csv
}
