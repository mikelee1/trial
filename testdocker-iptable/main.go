package main

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
)

var host = "http://192.168.9.67:2375"
var logger *logging.Logger
var dockerClient *docker.Client

func init() {
	logger = logger2.GetLogger()
	dockerClient, _ = docker.NewClient(host)
}

func main() {
	containerName := "test"
	RemoveContainer(containerName)
	container, err := dockerClient.CreateContainer(docker.CreateContainerOptions{
		Name: containerName,
		Config: &docker.Config{
			Image: "alpine-ipts:latest",
			//./start.sh --rm -ti --cap-add=NET_ADMIN
			//Cmd: []string{"sh", "-c", "sh"},
			Cmd:        []string{"sh", "-c", "./start.sh && sh"},
			WorkingDir: "/home",
			Tty:        true,
			ExposedPorts: map[docker.Port]struct{}{
				"8882": struct{}{},
			},
		},
		HostConfig: &docker.HostConfig{
			Binds: []string{
				//mike 宿主机上的地址进行映射
				"/home/ubuntu/start.sh:/home/start.sh:rw",
			},
			Privileged: true,
			UsernsMode: "host",
			PortBindings: map[docker.Port][]docker.PortBinding{
				"8882": []docker.PortBinding{
					docker.PortBinding{
						HostIP:   "0.0.0.0",
						HostPort: "18882",
					},
				},
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
		logger.Info(value.Names[0])
		//删除test容器
		if value.Names[0] == "/"+containerName {
			dockerClient.InspectContainer(value.ID)
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
