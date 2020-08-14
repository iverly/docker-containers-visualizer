package main

import (
	"docker-containers-visualizer/cli"
	"docker-containers-visualizer/cmd/dcv"
	"fmt"
	"os"
)

func main() {
	dcvCli := cli.NewDcvCli()

	if err := dcv.RunDcv(dcvCli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
