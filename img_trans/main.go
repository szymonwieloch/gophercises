package main

import (
	"log"
	"os"
)

func main() {
	args := parseArgs()

	createUploadDir(args.TmpDir)
	go backgroundClenup(args.TmpDir, args.CleanupCheck, args.Expire)
	runServer(args)
}

func createUploadDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal("Error when creating upload directory:", err)
	}
}
