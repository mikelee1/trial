package main_test

import (
	docker2 "github.com/fsouza/go-dockerclient"
	logger2 "myproj/try/common/logger"
	"testing"
	"fmt"
	dc "github.com/fsouza/go-dockerclient"
)

var (
	err    error
	logger = logger2.GetLogger()
	host   = "http://192.168.9.83:2375"
	//fsouzaCli, err = docker2.NewTLSClient(host,"./testdockerclient/client/cert.pem","./testdockerclient/client/key.pem","./testdockerclient/ca.pem")
	fsouzaCli, _ = docker2.NewClient(host)
)

func RemoveContainer(containerName string)  {
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

func Test_updateconfig(t *testing.T) {
	imgName := "192.168.9.87:5000/busybox"
	containerName := "test"

	RemoveContainer(containerName)

	_, err = fsouzaCli.InspectImage(imgName+":latest")
	if err != nil {
		logger.Info("pull new")
		err = fsouzaCli.PullImage(docker2.PullImageOptions{
			Repository: "192.168.9.87:5000/busybox",
			Tag:        "latest",
			Registry:   "192.168.9.87",
		}, docker2.AuthConfiguration{})
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info("good")
	}

	//创建容器
	resp, err := fsouzaCli.CreateContainer(dc.CreateContainerOptions{
		Name: containerName,
		Config: &dc.Config{
			Image: imgName,
			Tty:   true,
			Cmd:   []string{"sh", "-c", " sh"},
		},
		HostConfig: &dc.HostConfig{
			AutoRemove: true,
			//UsernsMode: "host",
			//ExtraHosts: []string{
			//	"test11:192.168.9.83",
			//},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}

	if err = fsouzaCli.StartContainer(resp.ID, nil); err != nil {
		panic(err)
	}

	container, err := fsouzaCli.InspectContainer(resp.ID)
	logger.Info("container.HostConfig.ExtraHosts: ",container.HostConfig.ExtraHosts)

	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("container: ",container)

	fsouzaCli.StopContainer(resp.ID,1)

	//创建容器
	resp, err = fsouzaCli.CreateContainer(dc.CreateContainerOptions{
		Name: containerName,
		Config: &dc.Config{
			Image: imgName,
			Tty:   true,
			Cmd:   []string{"sh", "-c", " sh"},
		},
		HostConfig: &dc.HostConfig{
			AutoRemove: true,
			//UsernsMode: "host",
			ExtraHosts: []string{
				"test11:192.168.9.83",
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}


	err = fsouzaCli.RestartContainer(resp.ID,1)
	if err != nil {
		logger.Error(err)
		return
	}

	container, err = fsouzaCli.InspectContainer(resp.ID)
	logger.Info("container.HostConfig.ExtraHosts: ",container.HostConfig.ExtraHosts)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("container: ",container)
	//
}

func Test_main(t *testing.T) {
	_, err = fsouzaCli.InspectImage("192.168.9.87:5000/busybox:latest")
	if err != nil {
		logger.Info("pull new")
		err = fsouzaCli.PullImage(docker2.PullImageOptions{
			Repository: "192.168.9.87:5000/busybox",
			Tag:        "latest",
			Registry:   "192.168.9.87:5000",
		}, docker2.AuthConfiguration{})
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info("good")
	} else {
		logger.Info("already exist")
	}
}

func Test_main2(t *testing.T) {
	_, err = fsouzaCli.InspectImage("192.168.9.87/busybox:latest")
	if err != nil {
		logger.Info("pull new")
		err = fsouzaCli.PullImage(docker2.PullImageOptions{
			Repository: "192.168.9.87:5000/busybox",
			Tag:        "latest",
			Registry:   "192.168.9.87",
		}, docker2.AuthConfiguration{})
		if err != nil {
			logger.Error(err)
			return
		}
		logger.Info("good")
	} else {
		logger.Info("already exist")
	}
}

//挂载方式1：创建mount的时候就把volume创建好
func Test_volume1(t *testing.T) {
	host = "http://192.168.9.83:2375"
	nfshost := "192.168.9.82"
	nfsdir := "/home/centos/testnfs/test"
	fsouzaCli, _ = docker2.NewClient(host)
	cname := "test1"
	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
		All: true,
	})
	for _, v := range cs {
		if v.Names[0] == "/"+cname {
			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
				ID:            v.ID,
				RemoveVolumes: true,
				Force:         true,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}
	//删除volume
	vs, err := fsouzaCli.ListVolumes(docker2.ListVolumesOptions{})
	for _, v := range vs {
		err = fsouzaCli.RemoveVolumeWithOptions(docker2.RemoveVolumeOptions{
			Name: v.Name,
		})
		if err != nil {
			continue
		}
		logger.Info("removed ", v.Name)
	}

	//验证挂载
	c, err := fsouzaCli.CreateContainer(docker2.CreateContainerOptions{
		Name: cname,
		Config: &docker2.Config{
			Tty:   true,
			Cmd:   []string{"sh", "-c", " sh"},
			Image: "alpine",
			Volumes: map[string]struct{}{
				"/nfs": struct{}{},
			},
		},
		HostConfig: &docker2.HostConfig{

			Mounts: []docker2.HostMount{
				//mount1
				docker2.HostMount{
					Type:   "volume",
					Target: "/var/hyperledger",
					VolumeOptions: &docker2.VolumeOptions{
						DriverConfig: docker2.VolumeDriverConfig{
							Name: "local",
							Options: map[string]string{
								"type":   "nfs",
								"device": fmt.Sprintf("%s:%s", nfshost, nfsdir),
								"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
							},
						},
					},
				},

				docker2.HostMount{
					Type:   "volume",
					Target: "/nfs",
					VolumeOptions: &docker2.VolumeOptions{
						DriverConfig: docker2.VolumeDriverConfig{
							Name: "local",
							Options: map[string]string{
								"type":   "nfs",
								"device": fmt.Sprintf("%s:%s", nfshost, nfsdir),
								"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
							},
						},
					},
				},

				////mount3,和mount1的挂载device一样
				//docker2.HostMount{
				//	Type:   "volume",
				//	Target: "/var/hyperledger/orderer2",
				//	VolumeOptions: &docker2.VolumeOptions{
				//		DriverConfig: docker2.VolumeDriverConfig{
				//			Name: "local",
				//			Options: map[string]string{
				//				"type":   "nfs",
				//				"device": fmt.Sprintf("%s:%s/orderer", nfshost, nfsdir),
				//				"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
				//			},
				//		},
				//	},
				//},
			},
			Binds: []string{
				//mike 宿主机上的地址进行映射
				//"/nfs:/var/hyperledger/orderer:rw",
				"/nfs/orderer:/var/hyperledger/orderer:rw",
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("good")
}

//挂载方式2：先创建volume，再mount到容器内
func Test_volume2(t *testing.T) {
	host = "http://192.168.9.83:2375"
	nfshost := "192.168.9.82"
	fsouzaCli, _ = docker2.NewClient(host)
	cname := "test1"
	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
		All: true,
	})
	for _, v := range cs {
		if v.Names[0] == "/"+cname {
			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
				ID:            v.ID,
				RemoveVolumes: true,
				Force:         true,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}
	//删除volume
	vs, err := fsouzaCli.ListVolumes(docker2.ListVolumesOptions{})
	for _, v := range vs {
		err = fsouzaCli.RemoveVolumeWithOptions(docker2.RemoveVolumeOptions{
			Name: v.Name,
		})
		if err != nil {
			continue
		}
		logger.Info("removed ", v.Name)
	}

	//创建volume
	vs1, err := fsouzaCli.CreateVolume(docker2.CreateVolumeOptions{
		Name:   "docker_nfs",
		Driver: "local",
		DriverOpts: map[string]string{
			"type":   "nfs",
			"device": nfshost + ":/home/centos/testnfs/test",
			"o":      "addr=" + nfshost + ",nolock,rw,tcp",
		},
	})

	vs2, err := fsouzaCli.CreateVolume(docker2.CreateVolumeOptions{
		Name:   "docker_nfs_common",
		Driver: "local",
		DriverOpts: map[string]string{
			"type":   "nfs",
			"device": nfshost + ":/home/centos/testnfs/test/common",
			"o":      "addr=" + nfshost + ",nolock,rw,tcp",
		},
	})

	//验证挂载
	c, err := fsouzaCli.CreateContainer(docker2.CreateContainerOptions{
		Name: cname,
		Config: &docker2.Config{
			Tty:   true,
			Cmd:   []string{"sh", "-c", " sh"},
			Image: "alpine",
		},
		HostConfig: &docker2.HostConfig{
			Mounts: []docker2.HostMount{
				docker2.HostMount{
					Type:   "volume",
					Source: vs1.Name,
					Target: "/var/hyperledger/orderer",
				},

				docker2.HostMount{
					Type:   "volume",
					Source: vs2.Name,
					Target: "/var/hyperledger/orderer/common",
				},

				docker2.HostMount{
					Type:   "volume",
					Source: vs1.Name,
					Target: "/var/hyperledger/orderer1",
				},
			},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("good")
}

//docker run -d -p 8081:8081 --restart=always -v docker_nfs:/nfs --name=test alpine

func Test_main1(t *testing.T) {
	images, err := fsouzaCli.ListImages(docker2.ListImagesOptions{
		All: true,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	for _, v := range images {
		logger.Info(v.ID)
	}
	logger.Info("good")
}
