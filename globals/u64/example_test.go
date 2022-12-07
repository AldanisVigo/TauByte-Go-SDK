package u64_test

import (
	"fmt"

	globalSym "bitbucket.org/taubyte/go-sdk-symbols/globals"
	"bitbucket.org/taubyte/go-sdk/globals/scope"
	"bitbucket.org/taubyte/go-sdk/globals/u64"
)

func ExampleUint64() {
	// Mocking the calls to the vm for usage in tests and playground
	globalSym.MockData{
		Data: map[string][]uint8{
			"/uint64/count":             {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x10},
			"/application/uint64/count": {0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x15},
		},
	}.Mock()

	{
		count, err := u64.GetOrCreate("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope:", count.Value())

		err = count.Set(23)
		if err != nil {
			return
		}

		count, err = u64.Get("count")
		if err != nil {
			return
		}

		fmt.Println("Global scope set:", count.Value())
	}

	{
		count, err := u64.Get("count", scope.Application)
		if err != nil {
			return
		}

		fmt.Println("Application scope:", count.Value())
	}

	// Output: Global scope: 16
	// Global scope set: 23
	// Application scope: 533
}
