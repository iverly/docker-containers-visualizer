package cli

import (
	"github.com/docker/docker/client"
	"github.com/spf13/cobra"
)

type Cli interface {
	Client() *client.Client
}

type DcvCli struct {
	client *client.Client
}

func NewDcvCli() *DcvCli {
	return &DcvCli{}
}

func (d *DcvCli) Client() *client.Client {
	return d.client
}

func (d *DcvCli) Initialize() error {
	var err error
	d.client, err = client.NewEnvClient()
	return err
}

func ShowHelp() func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}
