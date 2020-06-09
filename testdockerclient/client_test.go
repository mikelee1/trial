package main_test

import (
	"github.com/fsouza/go-dockerclient"
	"fmt"
	"testing"
)

var (
	host string
)

func init() {
	host = "http://192.168.9.67:2375"
}

//列出所有容器
func Test_containers(t *testing.T) {
	var err error
	fsouzaCli, err := docker.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	// 获取所有的容器
	allContainers, err := fsouzaCli.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {
		panic(err)
		return
	}
	for _, c := range allContainers {
		fmt.Println(c.Names)
	}
}

//列出所有网络
func Test_networks(t *testing.T) {
	var err error
	fsouzaCli, err := docker.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	// 获取所有的容器
	allNetworks, err := fsouzaCli.ListNetworks()
	if err != nil {
		panic(err)
		return
	}
	for _, c := range allNetworks {
		fmt.Println(c.Name)
	}
}
