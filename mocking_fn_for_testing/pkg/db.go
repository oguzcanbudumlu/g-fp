package pkg

import "os"

type authorizationFn func() bool

type Db struct {
	AuthorizationFn authorizationFn
}

func (d *Db) IsAuthorized() bool {
	return d.AuthorizationFn()
}

func NewDB() *Db {
	return &Db{
		AuthorizationFn: argsAuthorization,
	}
}

func argsAuthorization() bool {
	user := os.Args[1]

	if user == "admin" {
		return true
	}

	return false
}
