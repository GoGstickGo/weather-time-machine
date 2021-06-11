package main

import (
	"fmt"
	"os"
	"weather-api/utils/cmd"
)

func main() {
	if err := cmd.NewDefaultWTSCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
