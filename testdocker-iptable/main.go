package main

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/op/go-logging"
	logger2 "myproj.lee/try/common/logger"
)

var host = "http://192.168.9.82:2375"
var logger *logging.Logger
var dockerClient *docker.Client

func init() {
	logger = logger2.GetLogger()
	dockerClient, _ = docker.NewClient(host)
}

func main() {
	containerName := "test32"
	RemoveContainer(containerName)

	//networks, _ := dockerClient.ListNetworks()
	//for _, network := range networks {
	//	if network.Name == "myoverlay" {
	//		fmt.Println(network.Name, network.Driver, network.ID)
	//	}
	//}
	//return

	container, err := dockerClient.CreateContainer(docker.CreateContainerOptions{
		Name: containerName,
		Config: &docker.Config{
			Image:      "192.168.9.87:5000/nginx:latest",
			Cmd:        []string{"sh","-c","nginx -c /etc/nginx/conf/nginx.conf; sh"},
			//WorkingDir: "/home",
			Tty:        true,
			//ExposedPorts: map[docker.Port]struct{}{
			//	"8883": struct{}{},
			//},
		},
		HostConfig: &docker.HostConfig{
			Privileged: true,
			UsernsMode: "host",
			Binds: []string{
				"/home/centos/go/src/wasabi/backEnd/msp/orderer-2-baas1-nginx.conf:/etc/nginx/conf/nginx.conf",
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	err = dockerClient.StartContainer(container.ID, &docker.HostConfig{})
	if err != nil {
		logger.Error(err)
		return
	}

}

func RemoveContainer(containerName string) {
	//获取容器列表
	conlist, err := dockerClient.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {
		logger.Error(err)
		return
	}

	//删除容器
	for _, value := range conlist {
		//logger.Info(value.Names[0])
		//删除test容器
		if value.Names[0] == "/"+containerName {
			//dockerClient.InspectContainer(value.ID)
			err = dockerClient.RemoveContainer(docker.RemoveContainerOptions{
				Force: true,
				ID:    value.ID,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}
}
