package main

import (
	"docker-containers-visualizer/services/cmd"
)

func main() {
	err := cmd.ExecuteRoot()
	if err != nil {
		panic(err)
	}
}
