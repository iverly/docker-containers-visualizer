package main

import (
	"docker-containers-visualizer/services/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.ExecuteRoot()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}
