package client

import (
	"errors"
	"fmt"
	"io"

	ipfsClientSym "bitbucket.org/taubyte/go-sdk-symbols/ipfs/client"
	"bitbucket.org/taubyte/go-sdk/errno"
	"github.com/ipfs/go-cid"
)

// Write writes the passed in data into the file.
// Returns how much was written and an error.
func (c *Content) Write(p []byte) (int, error) {
	var n uint32
	err := ipfsClientSym.IpfsWriteFile(uint32(c.client), c.id, &p[0], uint32(len(p)), &n)
	if err != 0 {
		return 0, fmt.Errorf("Failed content write with %v", err)
	}

	return int(n), nil
}

// Close closes the file associated with the content.
// Returns an error.
func (c *Content) Close() error {
	err := ipfsClientSym.IpfsCloseFile(uint32(c.client), c.id)
	if err != 0 {
		return fmt.Errorf("Failed closing content with %v", err)
	}

	return nil
}

// Cid returns the cid of the file and an error.
func (c *Content) Cid() (cid.Cid, error) {
	_cid := make([]byte, CidBufferSize)

	err0 := ipfsClientSym.IpfsFileCid(uint32(c.client), c.id, &_cid[0])
	if err0 != 0 {
		return cid.Cid{}, fmt.Errorf("Failed getting cid with %v", err0)
	}

	_, cidFromBytes, err := cid.CidFromBytes(_cid)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("Failed decoding cid with: %v", err)
	}

	return cidFromBytes, err
}

func (c *Content) Read(p []byte) (int, error) {
	if p == nil || len(p) == 0 {
		return 0, errors.New("Invalid buffer to read into")
	}

	var counter uint32
	err := ipfsClientSym.IpfsReadFile(uint32(c.client), c.id, &p[0], uint32(len(p)), &counter)
	if err != 0 {
		if err == errno.ErrorEOF {
			return int(counter), io.EOF
		} else {
			return 0, fmt.Errorf("Failed reading content with %v", err)
		}
	}

	return int(counter), nil
}

// Seek moves to a position inside the file.
// Offset is how much to move the current position
// Whence has three options: 0 = SeekStart, 1 = SeekCurrent, or 2 = SeekEnd
// Combines both offset and whence to find a new offset inside the file
// Returns the new offset and an error
func (c *Content) Seek(offset int64, whence int) (int64, error) {
	var n int
	err := ipfsClientSym.IpfsSeekFile(uint32(c.client), c.id, offset, whence, &n)
	if err == errno.ErrorInvalidWhence {
		return 0, errors.New("Invalid Valid. Valid whence are 0 = SeekStart, 1 = SeekCurrent, or 2 = SeekEnd")
	}
	if err != 0 {
		return 0, fmt.Errorf("Failed seek with %v", err)
	}

	return int64(n), nil
}

// Push adds the file into the network then closes the file
// Returns cid and an error
func (c *Content) Push() (cid.Cid, error) {
	_, err := c.Seek(0, 0)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("Failed seeking 0, 0 of content with: %v", err)
	}

	_cid := make([]byte, CidBufferSize)
	err0 := ipfsClientSym.IpfsPushFile(uint32(c.client), c.id, &_cid[0])
	if err0 != 0 {
		return cid.Cid{}, fmt.Errorf("Failed reading content with: %v", err0)
	}

	_, cidFromBytes, err := cid.CidFromBytes(_cid)
	if err != nil {
		return cid.Cid{}, fmt.Errorf("Failed decoding cid with: %v", err)
	}

	// Closing content so it cannot be modified
	c.Close()

	return cidFromBytes, nil
}
