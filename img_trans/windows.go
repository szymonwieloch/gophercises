//go:build windows

package main

import (
	"io/fs"
	"syscall"
	"time"
)

func accessTime(info fs.FileInfo) time.Time {
	var result time.Time
	wtime, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if ok {
		result = time.Unix(0, wtime.LastAccessTime.Nanoseconds())
	}
	return result
}
