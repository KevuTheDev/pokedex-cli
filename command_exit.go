package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("Goodbye")
	fmt.Println()
	os.Exit(0)

	return nil
}
