package cmd

import (
	"docker-containers-visualizer/services/dapi"
	"github.com/spf13/cobra"
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
		panic(err)
	}
}
