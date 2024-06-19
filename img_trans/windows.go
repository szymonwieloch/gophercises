//go:build windows

package main

import (
	"io/fs"
	"log"
	"syscall"
	"time"
)

func accessTime(info fs.FileInfo) time.Time {
	var result time.Time
	wtime, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		result = time.Unix(0, wtime.LastAccessTime.Nanoseconds())
	} else {
		log.Println("Could not obtain last access time for ", info.Name())
	}
	return result
}
