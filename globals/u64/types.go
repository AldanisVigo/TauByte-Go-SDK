package u64

import "bitbucket.org/taubyte/go-sdk/globals/internal"

type Uint64 interface {
	internal.BaseInterface[uint64]
	internal.NumberInterface[uint64]
}
