package cmd

import (
	"docker-containers-visualizer/services/dapi"
	"fmt"
	tm "github.com/buger/goterm"
	"github.com/spf13/cobra"
	"os"
	"sort"
	"strings"
	"time"
)

var (
	rootCmd = &cobra.Command{
		Use:   "docker-containers-visualizer",
		Short: "Visualize all container running on the host with networks, stats and others information!",
		Long:  "\nVisualize all container running on the host with networks, stats and others information!\nMade by iverly with love in Go",
		Run:   runRoot,
	}
)

func ExecuteRoot() error {
	return rootCmd.Execute()
}

func runRoot(cmd *cobra.Command, args []string) {
	docker := dapi.DockerAPI{}
	err := docker.Init()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	for {
		tm.Clear()
		tm.MoveCursor(0, 0)

		table := tm.NewTable(0, 10, 5, ' ', 0)
		fmt.Fprintf(table, "ID\tName\tNetwork\tAddress\tCPU\tRAM\n")

		keys := make([]string, 0, len(docker.Containers))
		for k := range docker.Containers {
			keys = append(keys, docker.Containers[k].Name + ";;" + k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			c := docker.Containers[strings.Split(k, ";;")[1]]
			fmt.Fprintf(table, "%s\t%s\t%s\t%s\t%.2f%%\t%.2f%%\n", c.ID, c.Name, c.Networks[0].Name, c.Networks[0].Address, c.Stats.CPUPercentage, c.Stats.MemoryPercentage)
		}

		tm.Println(table)
		tm.Flush()
		time.Sleep(time.Second)
	}

}
