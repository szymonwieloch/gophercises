package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func runServer(args args) {
	port := fmt.Sprintf(":%d", args.Port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/upload", createUploadHandler(args))
	mux.HandleFunc("/choice/", choiceHandler)
	mux.HandleFunc("/image/", createImageHandler(args))
	log.Fatal(http.ListenAndServe(port, mux))
}

//go:embed index.htm
var indexPage []byte

func rootHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_, err := w.Write(indexPage)
	if err != nil {
		log.Println("Error when writing the index page: ", err)
	}
}

func createUploadHandler(args args) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		file, header, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error receiving file: ", err)
			return
		}
		defer file.Close()

		path, err := createFile(file, header.Filename, args.TmpDir)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error crating uploaded file: ", err)
			return
		}
		url := fmt.Sprintf("/choice/%s", path)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func createFile(r io.Reader, fileName string, tmpDir string) (string, error) {
	hash := md5.New()
	tee := io.TeeReader(r, hash)
	tmpFile, err := os.CreateTemp("", fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		tmpFile.Close()
		if err != nil {
			os.Remove(tmpFile.Name())
		}
	}()
	_, err = io.Copy(tmpFile, tee)
	if err != nil {
		return "", err
	}
	tmpFile.Close()
	checksum := hex.EncodeToString(hash.Sum(nil))
	uniqueDir := path.Join(tmpDir, checksum)
	err = os.Mkdir(uniqueDir, os.ModePerm)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	finalPath := path.Join(uniqueDir, fileName)
	// during reupload the file might exist
	_, err = os.Stat(finalPath)
	if err != nil && !os.IsExist(err) {
		return "", err
	}
	if err == nil {
		// file alrady exists, remove
		err = os.Remove(finalPath)
		if err != nil {
			return "", err
		}
	}
	err = os.Rename(tmpFile.Name(), finalPath)
	if err != nil {
		return "", err
	}
	return path.Join(checksum, fileName), nil

}

func choiceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	_ = strings.TrimPrefix(r.URL.Path, "/choice/")

}

func createImageHandler(args args) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		p := strings.TrimPrefix(r.URL.Path, "/image/")
		completePath := path.Join(args.TmpDir, p)
		file, err := os.Open(completePath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error opening local file: ", err)
			return
		}
		defer file.Close()
		// TODO set content-type
		_, err = io.Copy(w, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error opening local file: ", err)
			return
		}
	}
}
