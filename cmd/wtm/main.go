package main

import (
	"fmt"
	"os"
	"weather-api/utils/cmd"
)

func main() {
	if err := cmd.NewDefaultWTMCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
