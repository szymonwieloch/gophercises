package main

import (
	"encoding/json"
	"os"
)

type keyCfg struct {
	ConsumerKey    string
	ConsumerSecret string
}

func readKeys(path string) keyCfg {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	var result keyCfg
	err = dec.Decode(&result)
	if err != nil {
		panic(err)
	}
	return result
}
