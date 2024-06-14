package main

import "github.com/alexflint/go-arg"

type setCmd struct {
	Key    string `arg:"positional,required"`
	Secret string `arg:"positional"`
}

type getCmd struct {
	Key string `arg:"positional,required"`
}

type args struct {
	Password string  `arg:"-p,--password,env,required" help:"Password required to encrypt or decrypt the file"`
	File     string  `arg:"-f,--file" help:"Path to the secret file where the data is stored, default: '~/.secrets'"`
	Set      *setCmd `arg:"subcommand:set" help:"Sets a secret in the given file"`
	Get      *getCmd `arg:"subcommand:get" help:"Gets a secret from a secret file and prints it"`
}

func parseArgs() args {
	var result args
	p := arg.MustParse(&result)
	if result.Set == nil && result.Get == nil {
		p.Fail("Please choose subcommand")
	}
	return result
}
