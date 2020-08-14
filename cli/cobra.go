package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

var HelpCommand = &cobra.Command{
	Use:               "help [command]",
	Short:             "Help about the command",
	PersistentPreRun:  func(cmd *cobra.Command, args []string) {},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {},
	RunE: func(c *cobra.Command, args []string) error {
		cmd, args, e := c.Root().Find(args)
		if cmd == nil || e != nil || len(args) > 0 {
			return errors.Errorf("unknown help topic: %v", strings.Join(args, " "))
		}

		helpFunc := cmd.HelpFunc()
		helpFunc(cmd, args)
		return nil
	},
}

func AddCommands(cmd *cobra.Command, dcvCli Cli) {
	cmd.AddCommand()
}

type TopLevelCommand struct {
	cmd    *cobra.Command
	dcvCli *DcvCli
	flags  *pflag.FlagSet
	args   []string
}

func NewTopLevelCommand(cmd *cobra.Command, dcvCli *DcvCli, flags *pflag.FlagSet) *TopLevelCommand {
	return &TopLevelCommand{cmd: cmd, dcvCli: dcvCli, flags: flags, args: os.Args[1:]}
}

func (tcmd *TopLevelCommand) HandleGlobalFlags() (*cobra.Command, []string, error) {
	cmd := tcmd.cmd

	flags := pflag.NewFlagSet(cmd.Name(), pflag.ContinueOnError)

	flags.SetInterspersed(false)

	flags.AddFlagSet(cmd.Flags())
	flags.AddFlagSet(cmd.PersistentFlags())

	if err := flags.Parse(tcmd.args); err != nil {
		if err := tcmd.Initialize(); err != nil {
			return nil, nil, err
		}
		return nil, nil, cmd.FlagErrorFunc()(cmd, err)
	}

	return cmd, flags.Args(), nil
}

func (tcmd *TopLevelCommand) Initialize() error {
	return tcmd.dcvCli.Initialize()
}
