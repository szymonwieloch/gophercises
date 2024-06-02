package main

import (
	"fmt"
	"os"
	"path"
	"sort"
)

func main() {
	args := parseCmd()
	//fmt.Println(args)
	if len(args.Dirs) == 0 {
		processDir(".", args)
	} else {
		for _, dir := range args.Dirs {
			processDir(dir, args)
		}
	}
}

func matches(name, filter string) bool {
	if len(filter) == 0 {
		return true
	}
	isMatched, err := path.Match(filter, name)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error when applying filter:", err.Error())
		return false
	}
	return isMatched
}

func newName(oldName string, cfgName string, idx int, fileCnt int) string {
	return fmt.Sprintf("%s (%d of %d)%s", cfgName, (idx + 1), fileCnt, path.Ext(oldName))
}

func processDir(dir string, args args) {
	//fmt.Println("Processing directory: ", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed to process directory", dir, ":", err)
		return
	}
	var matched []string
	for _, file := range files {
		// fmt.Println("Checking file:", path.Join(dir, file.Name()))
		if file.IsDir() {
			if args.Recursive {
				processDir(path.Join(dir, file.Name()), args)
			}
		} else if matches(file.Name(), args.Filter) {
			matched = append(matched, file.Name())
		}
	}
	sort.Strings(matched)
	for idx, fileName := range matched {
		filePath := path.Join(dir, fileName)
		renameInto := newName(fileName, args.Name, idx, len(matched))
		renameInto = path.Join(dir, renameInto)
		fmt.Printf("Renaming '%s' into '%s'\n", filePath, renameInto)
		if !args.DryRun {
			err = os.Rename(filePath, renameInto)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Failed to rename file: ", err)
			}
		}
	}
}
