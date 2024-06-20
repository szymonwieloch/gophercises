package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strings"

	"github.com/szymonwieloch/gophercises/img_trans/prm"
)

func runServer(args args) {
	port := fmt.Sprintf(":%d", args.Port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/upload", createUploadHandler(args))
	mux.HandleFunc(choicePrefix, choiceHandler)
	mux.HandleFunc(imagePrefix, createImageHandler(args))
	mux.HandleFunc(viewPrefix, viewHandler)
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
		//url := fmt.Sprintf("%s%s/%s", choicePrefix, checksum, header.Filename)
		path := fmt.Sprintf("%s/%s", checksum, header.Filename)
		mode := prm.Circle
		url := viewUrl(path, imageOptions{
			mode: &mode,
			n:    20,
		})
		http.Redirect(w, r, url, http.StatusFound)
	}
}

func choiceHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	imgPath := strings.TrimPrefix(r.URL.Path, viewPrefix)
	imgOpts, err := parseImageOptions(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad image options: ", r.URL)
		return
	}

	cd := choiceData{}
	if imgOpts.mode == nil {
		cd.Choice = "Mode"
		cd.Options = []choiceOptionData{}
		for _, mode := range prm.ModeValues() {
			newOpts := imgOpts
			newOpts.mode = &mode
			cd.Options = append(cd.Options, choiceOptionData{
				ImgLink:  imageUrl(imgPath, newOpts),
				ViewLink: viewUrl(imgPath, newOpts),
			})
		}
	} else if imgOpts.n == 0 {
		cd.Choice = "Number Of Elements"
		cd.Options = []choiceOptionData{}
		nOptions := []uint{5, 10, 20, 40, 80}
		for _, n := range nOptions {
			newOpts := imgOpts
			newOpts.n = n
			cd.Options = append(cd.Options, choiceOptionData{
				ImgLink:  imageUrl(imgPath, newOpts),
				ViewLink: viewUrl(imgPath, newOpts),
			})
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad choice link: ", r.URL)
		return
	}
	err = choiceTemplate.Execute(w, &cd)
	if err != nil {
		log.Println("Error executing choice template: ", err)
	}

}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	imgPath := strings.TrimPrefix(r.URL.Path, viewPrefix)
	imgOpts, err := parseImageOptions(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Bad image options: ", r.URL)
		return
	}
	shapeOpts := imgOpts
	shapeOpts.mode = nil
	nOpts := imgOpts
	nOpts.n = 0
	choices := []viewChoiceData{
		{Name: "Shape", ChoiceLink: choiceUrl(imgPath, shapeOpts)},
		{Name: "Number Of Shapes", ChoiceLink: choiceUrl(imgPath, nOpts)},
	}
	vd := viewData{
		ImgLink: imageUrl(imgPath, imgOpts),
		Choices: choices,
	}

	err = viewTemplate.Execute(w, &vd)
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
