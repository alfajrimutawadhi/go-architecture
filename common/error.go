package common

import (
	"strings"

	log "github.com/sirupsen/logrus"
)

type Error struct {
	Flag string
	Err  error
	Args string
}

func WrapError(err error, flag string, args ...string) {
	e := Error{
		Flag: strings.ToUpper(flag),
		Err:  err,
	}
	if args != nil {
		e.Args = args[0]
	}

	log.Errorf("[%s] - %s : %v", e.Flag, e.Args, e.Err)
}
