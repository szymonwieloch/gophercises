package main

import (
	"log"
	"os"
)

func main() {
	args := parseArgs()

	createUploadDir(args.TmpDir)

	// err := prm.Transform("example.jpg", "out.jpg", 100, prm.Triangle)
	// if err != nil {
	// 	panic(err)
	// }

	runServer(args)
}

func createUploadDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal("Error when creating upload directory:", err)
	}
}
