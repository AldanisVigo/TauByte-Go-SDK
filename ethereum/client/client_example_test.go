package ethereum_test

import (
	"fmt"
	"math/big"

	ethereumSym "bitbucket.org/taubyte/go-sdk-symbols/ethereum/client"
	ethereum "bitbucket.org/taubyte/go-sdk/ethereum/client"
)

func ExampleNew() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client: 5,
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	fmt.Printf("%d\n", client)
	// Output: 5
}

func ExampleClient_CurrentBlockNumber() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:             4,
		CurrentBlockNumber: 5,
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	blockNumber, err := client.CurrentBlockNumber()
	if err != nil {
		return
	}

	fmt.Printf("%d\n", blockNumber)
	// Output: 5
}

var block *ethereum.Block

func ExampleClient_BlockByNumber() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:        4,
		BlockByNumber: 5,
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	block, err = client.BlockByNumber(big.NewInt(20))
	if err != nil {
		return
	}

	fmt.Println("success")
	// Output: success
}

func ExampleClient_CurrentChainId() {
	// Mocking the calls to the vm for usage in tests and playground
	mockData := ethereumSym.MockData{
		Client:         4,
		CurrentChainId: big.NewInt(5),
	}
	mockData.Mock()

	client, err := ethereum.New("https://testRPC.url")
	if err != nil {
		return
	}

	chainId, err := client.CurrentChainId()
	if err != nil {
		return
	}

	fmt.Printf("%d\n", chainId)
	// Output: 5
}
