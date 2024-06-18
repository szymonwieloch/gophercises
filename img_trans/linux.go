//go:build linux

package main

import (
	"io/fs"
	"syscall"
	"time"
)

func accessTime(info fs.FileInfo) time.Time {
	var result time.Time
	ltime, ok := info.Sys().(*syscall.Stat_t)
	if ok {
		result = time.Unix(ltime.Atim.Sec, ltime.Atim.Nsec)
	}
	return result
}
