package main

import (
	"time"

	"github.com/alexflint/go-arg"
)

type args struct {
	Port         uint16        `arg:"-p,--port" help:"Port to open the HTTP server on" default:"3000"`
	TmpDir       string        `arg:"-d,--dir" help:"Directory to store the temporal images" default:"./imgs"`
	Expire       time.Duration `arg:"-e,--expire" help:"Remove temporal files older than the provided expiration period", default:"1h"`
	CleanupCheck time.Duration `arg:"-c,--cleanup" help:"Expiration check running period", default:"1m"`
}

func parseArgs() args {
	var result args
	arg.MustParse(&result)
	return result
}
