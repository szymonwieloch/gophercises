package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/szymonwieloch/gophercises/img_trans/prm"
)

type imageOptions struct {
	mode *prm.Mode
	n    uint
}

func imageFileName(opts imageOptions, ext string) string {
	mode := "none"
	if opts.mode != nil {
		mode = strings.ToLower(opts.mode.String())
	}

	return fmt.Sprintf("%s_%d%s", mode, opts.n, ext)
}

func createOriginalFile(r io.Reader, fileName string, tmpDir string) (string, error) {
	hash := md5.New()
	tee := io.TeeReader(r, hash)
	tmpFile, err := os.CreateTemp("", "")
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

	finalPath := path.Join(uniqueDir, imageFileName(imageOptions{}, path.Ext(fileName)))
	// during reupload the file might exist
	_, err = os.Stat(finalPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
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
	return checksum, nil

}

func openOrCreateFile(tmpDir string, checksum string, opts imageOptions, ext string) (*os.File, error) {
	completePath := path.Join(tmpDir, checksum, imageFileName(opts, ext))
	file, err := os.Open(completePath)
	if err == nil {
		return file, nil
	}
	if !os.IsNotExist(err) {
		return nil, err
	}
	if opts.mode == nil {
		return nil, errors.New("original file does not exist")
	}
	originalPath := path.Join(tmpDir, checksum, imageFileName(imageOptions{}, ext))
	err = prm.Transform(originalPath, completePath, opts.n, *opts.mode)
	if err != nil {
		return nil, err
	}
	return os.Open(completePath)
}
