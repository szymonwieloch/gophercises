package main

import (
	"io/fs"
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

func backgroundClenup(tmpDir string, checkPeriod, expiration time.Duration) {
	for {
		runCleanup(tmpDir, expiration)
		time.Sleep(checkPeriod)
	}
}

func runCleanup(tmpDir string, expiration time.Duration) {
	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		log.Println("Error listing directory while doing cleanup: ", err)
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		cleanupDir(path.Join(tmpDir, entry.Name()), expiration)
	}
}

func cleanupDir(dir string, expiration time.Duration) {
	lat, err := lastAccessTime(dir)
	var emptyTime time.Time
	if err != nil || lat == emptyTime {
		log.Println("Could not get last access time of dir ", dir, " : ", err)
		return
	}
	if time.Now().After(lat.Add(expiration)) {
		err = os.RemoveAll(dir)
		if err != nil {
			log.Println("Could not remove directory ", dir, " : ", err)
		}
	}
}

func lastAccessTime(dir string) (time.Time, error) {
	var result time.Time
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println("Error walking dir ", dir, " : ", err)
			return nil
		}
		at := getAccessTime(info)
		if at.After(result) {
			result = at
		}
		return nil
	})
	return result, err
}

func getAccessTime(info fs.FileInfo) time.Time {
	result := info.ModTime()
	at := accessTime(info)
	if at.After(result) {
		result = at
	}
	return result
}
