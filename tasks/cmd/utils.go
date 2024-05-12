package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/szymonwieloch/gophercises/tasks/pkg"
)

func parseTaskID(id string) (pkg.TaskID, error) {
	number, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID: %w", err)
	}
	return pkg.TaskID(number), nil
}

func handleResult(err error) {
	if err == nil {
		return
	}
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
