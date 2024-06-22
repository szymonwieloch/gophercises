package main

import "github.com/alexflint/go-arg"

type args struct {
	Input  string   `arg:"-i,--input,required" help:"Path to a CSV file containing data"`
	X      string   `arg:"-x,required" help:"Name of the CSV column containing X axis data"`
	Y      []string `arg:"-y,required" help:"Name of the CSV column containing Y axis data"`
	Output string   `arg:"-o,--output" help:"Path to the output CSV file" default:"out.svg"`
}

func parseArgs() args {
	var args args
	p := arg.MustParse(&args)
	if len(args.Y) == 0 {
		p.Fail("At least one CSV column needs to be specified with the '-y' paramter")
	}
	return args
}
