package f32

import "bitbucket.org/taubyte/go-sdk/globals/internal"

func Get(name string, scope ...internal.Option) (Float32, error) {
	u := new(name, scope...)

	return u, u.Get()
}

func GetOrCreate(name string, scope ...internal.Option) (Float32, error) {
	u := new(name, scope...)

	return u, u.GetOrCreate()
}
