package main

import (
	"time"

	"github.com/alexflint/go-arg"
)

type args struct {
	Keys    string        `arg:"-k,--keys" default:"keys.json" help:"JSON file with account key and secreat key"`
	Users   string        `arg:"-u,--users" default:"users.csv" help:"CSV file where the retweeting user names are going to be stored"`
	Tweet   string        `arg:"-t,--tweet,required" help:"Tweet ID"`
	Period  time.Duration `arg:"-p,--period" help:"Period for checking for new retweets" default:"20s"`
	Winners uint          `arg:"-w,--winners" help:"Number of winneres to pick" default:"0"`
}

func parseArgs() args {
	var a args
	arg.MustParse(&a)
	return a
}
