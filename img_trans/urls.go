package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/szymonwieloch/gophercises/img_trans/prm"
)

const (
	modeKey      = "mode"
	shapesKey    = "shapes"
	viewPrefix   = "/view/"
	choicePrefix = "/choice/"
	imagePrefix  = "/image/"
)

func parseImageOptions(r *http.Request) (imageOptions, error) {
	var result imageOptions
	modeStr := r.URL.Query().Get(modeKey)
	if modeStr != "" {
		mode, err := prm.ParseModeString(modeStr)
		if err != nil {
			return imageOptions{}, err
		}
		result.mode = &mode
	}
	nStr := r.URL.Query().Get(shapesKey)
	if nStr != "" {
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return imageOptions{}, err
		}
		if n < 0 {
			return imageOptions{}, errors.New("number of shapes needs to be greater than 0")
		}
		result.n = uint(n)
	}

	// if result.n == 0 && result.mode != nil || result.mode == nil && result.n != 0 {
	// 	return imageOptions{}, errors.New("inavalid combination of image parameters")
	// }
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

func imgQueryString(opts imageOptions) string {
	params := url.Values{}
	if opts.mode != nil {
		params.Add(modeKey, opts.mode.String())
	}
	if opts.n != 0 {
		params.Add(shapesKey, fmt.Sprint(opts.n))
	}
	return params.Encode()
}

func viewUrl(path string, opts imageOptions) string {

	return fmt.Sprintf("%s%s?%s", viewPrefix, path, imgQueryString(opts))
}

func imageUrl(path string, opts imageOptions) string {
	return fmt.Sprintf("%s%s?%s", imagePrefix, path, imgQueryString(opts))
}

func choiceUrl(path string, opts imageOptions) string {
	return fmt.Sprintf("%s%s?%s", choicePrefix, path, imgQueryString(opts))
}
