package dapi

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"sort"
	"sync"
)

type DockerAPI struct {
	Client     *client.Client
	Containers map[string]*Container
	Networks   []string
	sync.Mutex
}

type Container struct {
	ID       string
	Name     string
	Image    string
	Running  bool
	Stats    *ContainerStats
	Networks []ContainerNetwork
}

type ContainerNetwork struct {
	Name    string
	Address string
}

type Network struct {
	Name string
}

func (d *DockerAPI) Init() error {
	var err error

	d.Containers = make(map[string]*Container)
	d.Client, err = client.NewEnvClient()
	if err != nil {
		return err
	}

	containers, err := d.GetContainers()
	if err != nil {
		return err
	}
	for _, c := range containers {
		d.Containers[c.ID] = c
	}

	networks, err := d.GetNetworks()
	if err != nil {
		return err
	}
	d.Networks = networks

	err = d.StartMonitoringRunningContainers()
	if err != nil {
		return err
	}

	return nil
}

func (d *DockerAPI) GetContainer(id string) (*Container, error) {
	container, found := d.Containers[id]
	if !found {
		tc, err := d.Client.ContainerInspect(context.Background(), id)
		if err != nil {
			return nil, err
		}

		stats := &ContainerStats{
			doneChan:         make(chan bool),
			previousCPUStats: &CPUStats{},
			StatsEntry: StatsEntry{
				CPUPercentage:    0,
				Memory:           0,
				MemoryLimit:      0,
				MemoryPercentage: 0,
			},
		}

		container = &Container{
			ID:       tc.ID[:12],
			Name:     tc.Name[1:],
			Image:    tc.Image,
			Running:  tc.State.Running,
			Stats:    stats,
			Networks: make([]ContainerNetwork, len(tc.NetworkSettings.Networks)),
		}

		i := 0
		for name, setting := range tc.NetworkSettings.Networks {
			container.Networks[i].Name = name
			container.Networks[i].Address = setting.IPAddress
			i++
		}
	}

	return container, nil
}

func (d *DockerAPI) GetNetworks() ([]string, error) {
	networks, err := d.Client.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return nil, err
	}

	ns := make([]string, len(networks))
	i := 0
	for _, network := range networks {
		ns[i] = network.Name
		i++
	}

	sort.Strings(ns)
	return ns, nil
}

func (d *DockerAPI) GetContainers() ([]*Container, error) {
	containers, err := d.Client.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	cs := make([]*Container, len(containers))
	i := 0
	for _, container := range containers {
		c, err := d.GetContainer(container.ID)
		if err != nil {
			return nil, err
		}
		cs[i] = c
		i++
	}

	return cs, nil
}

func (d *DockerAPI) GetContainersOnNetwork(network string) []*Container  {
	var cs []*Container
	for _, c := range d.Containers {
		for _, nn := range c.Networks {
			if nn.Name == network {
				cs = append(cs, c)
			}
		}
	}
	return cs
}
