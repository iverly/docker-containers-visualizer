package main

import (
	"docker-containers-visualizer/cli"
	"docker-containers-visualizer/cli/command/root"
	"fmt"
	"os"
)

func newDcvCommand(dcvCli *cli.DcvCli) *cli.TopLevelCommand {
	cmd := root.NewRootCommand(dcvCli)
	return cli.NewTopLevelCommand(cmd, dcvCli, cmd.Flags())
}

func runDcv(dcvCli *cli.DcvCli) error {
	tcmd := newDcvCommand(dcvCli)
	cmd, args, err := tcmd.HandleGlobalFlags()
	if err != nil {
		return err
	}

	if err := tcmd.Initialize(); err != nil {
		return err
	}

	cmd.SetArgs(args)
	return cmd.Execute()
}

func main() {
	dcvCli := cli.NewDcvCli()

	if err := runDcv(dcvCli); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
