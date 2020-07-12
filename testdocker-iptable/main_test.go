package main

import (
	"github.com/docker/docker/api/types/swarm"
	"github.com/fsouza/go-dockerclient"
	"fmt"
	"testing"
	"github.com/docker/docker/api/types/mount"
)

var err error

func init() {
	Init()
}
func Init() {
	dockerClient, err = docker.NewClient(host)
	if err != nil {
		panic(err)
	}
}

func Test_createoverlay(t *testing.T) {
	n, err := dockerClient.CreateNetwork(docker.CreateNetworkOptions{
		Name:       "myoverlay",
		Driver:     "overlay",
		Attachable: true,
		//Ingress:    true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(n.ID)
}

func Test_createbridge(t *testing.T) {
	networks, err := dockerClient.ListNetworks()
	for _, network := range networks {
		if network.Name == "mybridge" {
			fmt.Println("id: ", network.ID, network.Name)
			err = dockerClient.RemoveNetwork(network.ID)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	n, err := dockerClient.CreateNetwork(docker.CreateNetworkOptions{
		Name:       "mybridge",
		Driver:     "bridge",
		Attachable: true,
		Scope:      "swarm",
		//Internal:   true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(n.ID)
}

var serviceName = "test11"

func Test_start(t *testing.T) {
	err := CreateIptableService(serviceName, 7052, "192.168.9.87", 30021)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateIptableService(containerName string, fromPort int32, toIp string, toPort int32) error {

	logger.Info("ipts暴露的内部端口: ", fromPort)

	services, err := dockerClient.ListServices(docker.ListServicesOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range services {
		//fmt.Println(v)
		if v.Spec.Name == serviceName {
			err = dockerClient.RemoveService(docker.RemoveServiceOptions{
				ID: v.ID,
			})
			if err != nil {
				return err
			}
		}
	}

	_, err = dockerClient.CreateService(docker.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image:   "nginx:latest",
					Command: []string{"sh", "-c", "sh"},
					TTY: true,
				},
			},
			Annotations: swarm.Annotations{
				Name: containerName,
			},
			Networks: []swarm.NetworkAttachmentConfig{
				//swarm.NetworkAttachmentConfig{Target: "myoverlay"},
				swarm.NetworkAttachmentConfig{Target: "mybridge"},
			},
			EndpointSpec: &swarm.EndpointSpec{Mode: swarm.ResolutionModeVIP, Ports: []swarm.PortConfig{
				swarm.PortConfig{
					Name:          "MainPort",
					TargetPort:    uint32(fromPort),
					PublishedPort: uint32(fromPort),
					PublishMode:   swarm.PortConfigPublishModeHost,
				},
			}},
		},
	})
	if err != nil {
		fmt.Errorf("err: %s", err.Error())
		return err
	}

	return nil
}

func Test_start1(t *testing.T) {
	serviceName := "nginxservice"
	err := CreateNginxService(serviceName)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateNginxService(containerName string) error {

	services, err := dockerClient.ListServices(docker.ListServicesOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, v := range services {
		fmt.Println(v.Spec.Name)
		if v.Spec.Name == serviceName {
			err = dockerClient.RemoveService(docker.RemoveServiceOptions{
				ID: v.ID,
			})
			if err != nil {
				return err
			}
		}
	}

	_, err = dockerClient.CreateService(docker.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: "nginx",
					TTY:   true,
					Mounts: []mount.Mount{
						mount.Mount{
							Type:   mount.TypeBind,
							Source: "/home/ubuntu/nginx.conf",
							Target: "/etc/nginx/nginx.conf",
						},
					},
				},
			},
			Annotations: swarm.Annotations{
				Name: containerName,
			},
			Networks: []swarm.NetworkAttachmentConfig{
				swarm.NetworkAttachmentConfig{Target: "myoverlay"},
			},
			EndpointSpec: &swarm.EndpointSpec{Mode: swarm.ResolutionModeVIP, Ports: []swarm.PortConfig{
				swarm.PortConfig{
					Name:       "MainPort",
					TargetPort: uint32(7777),
					//PublishedPort: uint32(fromPort),
					//PublishMode:   swarm.PortConfigPublishModeHost,
				},
			}},
		},
	})
	if err != nil {
		fmt.Errorf("err: %s", err.Error())
		return err
	}

	return nil
}
