package main

import (
	"fmt"

	"github.com/szymonwieloch/gophercises/tasks/cmd"
)

func main() {
	fmt.Println("CLI task manager")
	fmt.Println()
	cmd.Execute()
}
