package main_test

import (
	docker2 "github.com/fsouza/go-dockerclient"
	logger2 "myproj.lee/try/common/logger"
	"testing"
	dc "github.com/fsouza/go-dockerclient"
	"github.com/docker/docker/api/types/swarm"
	"fmt"
)

var (
	err    error
	logger = logger2.GetLogger()
	host   = "http://192.168.9.67:2375"
	//fsouzaCli, _ = docker2.NewTLSClient(host,"./testdockerclient/client/cert.pem","./testdockerclient/client/key.pem","./testdockerclient/ca.pem")
	fsouzaCli, _ = docker2.NewClient(host)
)

func RemoveContainer(containerName string) {
	//获取容器列表
	conlist, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
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
			err = fsouzaCli.RemoveContainer(dc.RemoveContainerOptions{
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

func Test_swarm(t *testing.T) {
	host = "http://192.168.9.82:2375"
	fsouzaCli, _ = docker2.NewClient(host)
	s, err := fsouzaCli.CreateService(dc.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			Mode: swarm.ServiceMode{},
			Annotations: swarm.Annotations{
				Name: "test3",
				Labels: map[string]string{
					"name": "test3",
				},
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image:   "alpine",
					TTY:     true,
					Command: []string{"sh", "-c", "sh"},
				},
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(s.ID)
	ss, err := fsouzaCli.ListServices(dc.ListServicesOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	for _, s := range ss {
		logger.Info(s.Spec.Name)
	}
}

func Test_swarmupdate(t *testing.T) {
	host = "http://192.168.9.82:2375"
	fsouzaCli, _ = docker2.NewClient(host)
	ss, err := fsouzaCli.ListServices(dc.ListServicesOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	for _, s := range ss {
		logger.Info(s.Spec.Name)
		if s.Spec.Name == "test3" {
			c, err := fsouzaCli.InspectService(s.ID)
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info(c.Spec.TaskTemplate.ContainerSpec.Hosts)

			c.Spec.TaskTemplate.ContainerSpec.Hosts = append(c.Spec.TaskTemplate.ContainerSpec.Hosts, "host2 127.0.0.1")
			err = fsouzaCli.UpdateService(s.ID, dc.UpdateServiceOptions{
				ServiceSpec: c.Spec,
				Version:     c.Version.Index,
			})

			if err != nil {
				logger.Error(err)
				return
			}
		}

	}
}

func Test_swarmupdate1(t *testing.T) {
	host = "http://192.168.9.82:2375"
	fsouzaCli, _ = docker2.NewClient(host)

	ss, err := fsouzaCli.ListServices(dc.ListServicesOptions{})
	if err != nil {
		logger.Error(err)
		return
	}

	for _, s := range ss {
		logger.Info(s.Spec.Name)
		if s.Spec.Name == "test3" {
			cc, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
				Filters: map[string][]string{
					"name": []string{"test3"},
				},
				All: true,
			})
			if err != nil {
				logger.Error(err)
				return
			}
			for _, oneContainer := range cc {
				logger.Info(oneContainer.Names, oneContainer.State)
				if oneContainer.State != "running" {
					err = fsouzaCli.RemoveContainer(dc.RemoveContainerOptions{
						ID: oneContainer.ID,
					})
					if err != nil {
						logger.Error(err)
						return
					}
				}
			}
			logger.Info(cc)
			c, err := fsouzaCli.InspectService(s.ID)
			if err != nil {
				logger.Error(err)
				return
			}
			c.Spec.TaskTemplate.ContainerSpec.Hosts = append(c.Spec.TaskTemplate.ContainerSpec.Hosts, "127.0.0.8 host4")
			err = fsouzaCli.UpdateService(s.ID, dc.UpdateServiceOptions{
				ServiceSpec: c.Spec,
				Version:     c.Version.Index,
			})

			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info(c.Spec.Name)

		}
	}
}

func Contains(s string, a []string) bool {
	for _, e := range a {
		if e == s {
			return true
		}
	}
	return false
}

var networkid string

func Test_SwarmOverlay(t *testing.T) {
	host = "http://192.168.9.82:2375"
	fsouzaCli, err = docker2.NewClient(host)
	if err != nil {
		logger.Infof("err: ", err)
		return
	}
	sss, err := fsouzaCli.ListServices(dc.ListServicesOptions{})
	if err != nil {
		logger.Error(err)
		return
	}
	for _, s := range sss {
		if s.Spec.Name == "test1"||s.Spec.Name == "test2"{
			fsouzaCli.RemoveService(dc.RemoveServiceOptions{
				ID:s.ID,
			})
		}

	}

	err = CreateOverlayNetwork(fsouzaCli, "test-overlay")
	if err != nil {
		logger.Infof("err: ", err)
		return
	}


	_, err = fsouzaCli.CreateService(dc.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			Mode: swarm.ServiceMode{},
			Annotations: swarm.Annotations{
				Name: "test1",
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image:   "alpine:latest",
					TTY:     true,
					Command: []string{"sh", "-c", "sh"},
				},
			},
			Networks: []swarm.NetworkAttachmentConfig{
				swarm.NetworkAttachmentConfig{Target: "test-overlay"},
			},
		},
	})
	if err != nil {
		logger.Infof("err: ", err)
		return
	}

	_, err = fsouzaCli.CreateService(dc.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			Mode: swarm.ServiceMode{},
			Annotations: swarm.Annotations{
				Name: "test2",
			},
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image:   "alpine:latest",
					TTY:     true,
					Command: []string{"sh", "-c", "sh"},
				},
			},
			Networks: []swarm.NetworkAttachmentConfig{
				swarm.NetworkAttachmentConfig{Target: "test-overlay"},
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}

	//sss, err = fsouzaCli.ListServices(dc.ListServicesOptions{})
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//for _, s := range sss {
	//	logger.Info(s.Spec.Name)
	//}
}

func CreateOverlayNetwork(client *dc.Client, overlayName string) error {
	logger.Info(client.Info())
	networks, err := client.ListNetworks()
	if err != nil {
		logger.Infof("err1: ", err)
		return err
	}
	logger.Infof("err: ",err)

	for _, network := range networks {
		if network.Name == overlayName {
			if network.Driver != "overlay" {
				logger.Warning("network %s has already exist, but not overlay type", overlayName)
				return fmt.Errorf("network %s has already exist, but not overlay type", overlayName)
			}
			logger.Warning("network %s has already exist, use old network", overlayName)
			return nil
		}
	}

	nopt := dc.CreateNetworkOptions{}
	nopt.Name = overlayName
	nopt.CheckDuplicate = true
	nopt.Driver = "overlay"
	nopt.Internal = false
	// 与fabric1.2版本一致的fsouza，不支持这两个参数
	// nopt.Attachable = true
	// nopt.Ingress = false
	var configs []dc.IPAMConfig

	subnet := "10.0.3.0/24"
	config := dc.IPAMConfig{Subnet: subnet}
	configs = append(configs, config)
	nopt.IPAM = &dc.IPAMOptions{
		Config: configs,
	}
	nopt.Attachable = true

	overlayNetwork, err := client.CreateNetwork(nopt)
	if err != nil {
		return err
	}
	networkid = overlayNetwork.ID
	logger.Info("successfully create overlay network, %+v", overlayNetwork)
	return nil
}
