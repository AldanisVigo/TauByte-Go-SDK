package f32

import "bitbucket.org/taubyte/go-sdk/globals/internal"

type Float32 interface {
	internal.BaseInterface[float32]
	internal.NumberInterface[float32]
}
