package pkg

import (
	pisces "github.com/cryptnode-software/pisces/lib"
	"github.com/cryptnode-software/pisces/lib/auth"
)

var (
	service pisces.AuthService = nil
)

func NewAuthService() (pisces.AuthService, error) {
	var err error
	if service != nil {
		return service, nil
	}

	env, err := NewEnv()
	if err != nil {
		return nil, err
	}

	if service, err = auth.NewService(env); err != nil {
		return nil, err
	}

	return service, nil
}
