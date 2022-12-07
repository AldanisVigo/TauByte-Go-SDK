package event

import (
	"testing"

	symbols "bitbucket.org/taubyte/go-sdk-symbols/p2p/event"
	"bitbucket.org/taubyte/go-sdk/errno"
	"github.com/ipfs/go-cid"
)

func TestFrom(t *testing.T) {
	var e Event

	testFrom, err := cid.Parse("QmZjBpQzRw8UovKaMoWJ3qvQrHZvErqQuXNyXNusm4XYK3")
	if err != nil {
		return
	}

	symbols.MockData{ClientId: 1, From: testFrom}.Mock()

	_, err = e.From()
	if err == nil {
		t.Error("Expected error")
		return
	}

	symbols.MockData{From: testFrom}.Mock()

	from, err := e.From()
	if err != nil {
		t.Error(err)
		return
	}
	if string(from.Bytes()) != string(testFrom.Bytes()) {
		t.Errorf("Expected %s, got %s", string(testFrom.Bytes()), string(from.Bytes()))
		return
	}

	// Size 0
	symbols.MockData{From: cid.Cid{}}.Mock()

	_, err = e.From()
	if err != nil {
		t.Error(err)
		return
	}

	// Basic get error
	symbols.MockData{From: testFrom}.Mock()

	symbols.GetP2PEventFrom = func(eventId uint32, bufPtr *byte) (error errno.Error) {
		return 1
	}

	_, err = e.From()
	if err == nil {
		t.Error("Expected error")
		return
	}
}
