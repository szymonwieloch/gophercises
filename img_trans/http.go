package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/szymonwieloch/gophercises/img_trans/prm"
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

		checksum, err := createOriginalFile(file, header.Filename, args.TmpDir)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error creating uploaded file: ", err)
			return
		}
		url := fmt.Sprintf("/choice/%s/%s", checksum, header.Filename)
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func choiceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	imgPath := strings.TrimPrefix(r.URL.Path, "/choice/")
	options := []string{"mode=triangle", "mode=rectangle"}
	vd := viewData{
		ImagePath: imgPath,
		Options:   options,
	}
	err := viewTemplate.Execute(w, &vd)
	if err != nil {
		log.Println("Error executing view template: ", err)
	}

}

func createImageHandler(args args) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		checksum, fileName := parseImagePath(r.URL.Path)
		if checksum == "" || fileName == "" {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error parsing image path: ", r.URL.Path)
			return
		}
		opts, err := parseImageOptions(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("Error parsing image options: ", err)
			return
		}
		file, err := openOrCreateFile(args.TmpDir, checksum, opts, path.Ext(fileName))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error opening or creating local file: ", err)
			return
		}
		defer file.Close()
		// TODO set content-type
		_, err = io.Copy(w, file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println("Error uploading local file: ", err)
			return
		}
	}
}

func parseImageOptions(r *http.Request) (imageOptions, error) {
	var result imageOptions
	modeStr := r.URL.Query().Get("mode")
	if modeStr != "" {
		mode, err := prm.ParseModeString(modeStr)
		if err != nil {
			return imageOptions{}, err
		}
		result.mode = &mode
	}
	nStr := r.URL.Query().Get("n")
	if nStr != "" {
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return imageOptions{}, err
		}
		if n < 0 {
			return imageOptions{}, errors.New("n needs to be greater than 0")
		}
		result.n = uint(n)
	}

	if result.n == 0 && result.mode != nil || result.mode == nil && result.n != 0 {
		return imageOptions{}, errors.New("inavalid combination of image parameters")
	}
	return result, nil
}

func parseImagePath(path string) (checksum string, fileName string) {
	idx := strings.LastIndex(path, "/")
	if idx == -1 {
		return "", ""
	}
	fileName = path[idx+1:]
	path = path[:idx]
	idx = strings.LastIndex(path, "/")
	if idx == -1 {
		return "", ""
	}
	checksum = path[idx+1:]
	return
}
