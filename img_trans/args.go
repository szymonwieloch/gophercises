package main

import (
	"time"

	"github.com/alexflint/go-arg"
)

type args struct {
	Port    uint16        `arg:"-p,--port" help:"Port to open the HTTP server on" default:"3000"`
	TmpDir  string        `arg:"-d,--dir" help:"Directory to store the temporal images" default:"./imgs"`
	Cleanup time.Duration `arg:"-c,--cleanup" help:"Remove temporal files period", default:"1h"`
}

func parseArgs() args {
	var result args
	arg.MustParse(&result)
	return result
}
