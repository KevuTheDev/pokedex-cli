package main

import (
	"fmt"
	"os"
)

func callbackExit(cfg *config, args ...string) error {
	fmt.Println("Goodbye")
	fmt.Println()
	os.Exit(0)

	return nil
}
