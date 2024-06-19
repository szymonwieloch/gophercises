//go:build linux

package main

import (
	"io/fs"
	"log"
	"syscall"
	"time"
)

func accessTime(info fs.FileInfo) time.Time {
	var result time.Time
	ltime, ok := info.Sys().(*syscall.Stat_t)
	if ok {

		result = time.Unix(ltime.Atim.Sec, ltime.Atim.Nsec)
	} else {
		log.Println("Could not obtain last access time for ", info.Name())
	}
	return result
}
