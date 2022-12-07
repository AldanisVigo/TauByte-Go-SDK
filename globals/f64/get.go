package f64

import "bitbucket.org/taubyte/go-sdk/globals/internal"

func Get(name string, scope ...internal.Option) (Float64, error) {
	u := new(name, scope...)

	return u, u.Get()
}

func GetOrCreate(name string, scope ...internal.Option) (Float64, error) {
	u := new(name, scope...)

	return u, u.GetOrCreate()
}
