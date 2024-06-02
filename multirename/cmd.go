package main

import "github.com/alexflint/go-arg"

type args struct {
	Name      string   `arg:"-n,--name,required" help:"New nae for the files"`
	Recursive bool     `arg:"-r,--recursive" help:"Recursive lookup."`
	Filter    string   `arg:"-f,--filter" help:"Filters files using the file name. Uses common bash syntax."`
	DryRun    bool     `arg:"-d,--dry" help:"Dry run."`
	Dirs      []string `arg:"positional" help:"List of directories to process. By default uses the current directory."`
}

func parseCmd() args {
	var result args
	arg.MustParse(&result)
	return result
}
