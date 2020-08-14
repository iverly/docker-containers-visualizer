package root

import (
	"context"
	"docker-containers-visualizer/cli"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/docker/docker/api/types"
	"github.com/spf13/cobra"
	"strings"
)

func NewRootCommand(dcvCli *cli.DcvCli) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dcv COMMAND [ARG...]",
		Short: "Visualize all container running on the host with networks, stats and others information !",
		SilenceUsage: true,
		SilenceErrors: true,
		TraverseChildren: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return runRoot(dcvCli, args)
			}
			return fmt.Errorf("dcv: '%s' is not a dcv command.\nSee 'dcv --help'", args[0])
		},
		Version: "1.0",
		DisableFlagsInUseLine: true,
	}

	addCommands(cmd, dcvCli)
	cmd.SetHelpCommand(cli.HelpCommand)

	cmd.PersistentFlags().BoolP("help", "h", false, "Show this help")
	cmd.PersistentFlags().BoolP("version", "v", false, "Print version information and quit")

	cmd.SetVersionTemplate("Dcv version {{.Version}}\n")
	return cmd
}

func addCommands(cmd *cobra.Command, dcvCli *cli.DcvCli) {
	cmd.AddCommand()
}

func runRoot(dcvCli *cli.DcvCli, args []string) error {
	networkList, err := dcvCli.Client().NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return err
	}

	tm.Clear()
	tm.MoveCursor(0, 0)

	table := tm.NewTable(0, 10, 5, ' ', 0)
	_, _ = fmt.Fprintf(table, "Network\tID\tName\tAddress/Subnet\n")

	for _, network := range networkList {
		if len(network.Containers) == 0 {
			continue
		}

		_, _ = fmt.Fprintf(table, "%s\t-\t-\t-\t\n", network.Name)

		for id, container := range network.Containers {
			inspect, err := dcvCli.Client().ContainerInspect(context.Background(), id)
			if err != nil {
				return err
			}
			localIP := strings.Split(container.IPv4Address, "/")[0]
			_, _ = fmt.Fprintf(table, "-\t%s\t%s\t%s\n", id[:12], inspect.Name[1:], localIP)
		}
	}

	_, _ = tm.Println(table)
	tm.Flush()
	return nil
}
