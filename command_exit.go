package main

import (
	"os"
)

func callbackExit() error {
	os.Exit(0)
	return nil
}
