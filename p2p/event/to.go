package event

import (
	"fmt"

	p2pEventSym "bitbucket.org/taubyte/go-sdk-symbols/p2p/event"
	"github.com/ipfs/go-cid"
)

// To returns a cid.Cid of the receiving node and an error.
func (e Event) To() (cid.Cid, error) {
	var size uint32
	err := p2pEventSym.GetP2PEventToSize(uint32(e), &size)
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Getting to address data size failed with: %s", err)
	}
	if size == 0 {
		return cid.Cid{}, nil
	}

	cidBytes := make([]byte, size)
	err = p2pEventSym.GetP2PEventTo(uint32(e), &cidBytes[0])
	if err != 0 {
		return cid.Cid{}, fmt.Errorf("Getting to address data failed with: %s", err)
	}

	_, cidToBytes, err0 := cid.CidFromBytes(cidBytes)
	return cidToBytes, err0
}
