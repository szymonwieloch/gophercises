package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func main() {
	args := parseArgs()
	switch {
	case args.Set != nil:
		checkErr(handleSet(args))
	case args.Get != nil:
		checkErr(handleGet(args))
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func handleGet(args args) error {
	secrets, err := readSecrets(args.File)
	if err != nil {
		return err
	}
	secret, ok := secrets[args.Get.Key]
	if !ok {
		return fmt.Errorf("secret '%s' not found", args.Get.Key)
	}
	fmt.Println(secret)
	return nil
}

func handleSet(args args) error {
	secret := args.Set.Secret
	if secret == "" {
		var err error
		secret, err = readCLISecret()
		if err != nil {
			return err
		}
	}
	secrets, err := readSecrets(args.File)
	if err != nil {
		return err
	}
	secrets[args.Set.Key] = secret
	path, err := filePath(args.File)
	if err != nil {
		return err
	}
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	err = enc.Encode(secrets)
	if err != nil {
		return err
	}
	return nil
}

func filePath(cfgPath string) (string, error) {
	if cfgPath != "" {
		return cfgPath, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return path.Join(home, ".secrets"), nil
}

func readSecrets(cfgPath string) (map[string]string, error) {
	path, err := filePath(cfgPath)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file does not exist: '%s'", path)
		} else {
			return nil, fmt.Errorf("error opening file '%s': %w", path, err)
		}
	}
	defer file.Close()
	var result map[string]string
	dec := json.NewDecoder(file)
	err = dec.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func readCLISecret() (string, error) {
	fmt.Println("Provide the secret:")
	byteSecret, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}

	secret := strings.TrimSpace(string(byteSecret))
	if len(secret) == 0 {
		return "", errors.New("the provided password is empty")
	}
	return secret, nil
}
