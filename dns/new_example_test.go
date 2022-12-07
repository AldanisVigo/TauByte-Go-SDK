package dns_test

import (
	"fmt"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/dns"
	"bitbucket.org/taubyte/go-sdk/dns"
)

func ExampleNewResolver() {
	// Mocking the calls to the vm for usage in tests and playground
	symbols.MockNew(5)

	resolver, err := dns.NewResolver()
	if err != nil {
		return
	}

	fmt.Printf("%d\n", resolver)
	// Output: 5
}
