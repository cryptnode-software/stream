package pkg

import (
	"os"

	pisces "github.com/cryptnode-software/pisces/lib"
)

var (
	env *pisces.Env = nil
)

func NewEnv() (*pisces.Env, error) {
	if env != nil {
		return env, nil
	}

	environ := pisces.Environment(
		os.Getenv("ENV"),
	)

	env = pisces.NewEnv(
		pisces.NewLogger(environ),
	)

	return env, nil
}
