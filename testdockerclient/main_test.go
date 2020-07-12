//package main_test
//
//import (
//	docker2 "github.com/fsouza/go-dockerclient"
//	logger2 "myproj.lee/try/common/logger"
//	"testing"
//	"fmt"
//	dc "github.com/fsouza/go-dockerclient"
//	dt "github.com/docker/docker/client"
//	"github.com/docker/docker/api/types/filters"
//	"context"
//	"github.com/docker/docker/api/types"
//	"github.com/ghodss/yaml"
//	"os"
//	"io/ioutil"
//)
//
//var (
//	err    error
//	logger = logger2.GetLogger()
//	host   = "http://192.168.9.82:2375"
//	//fsouzaCli, err = docker2.NewTLSClient(host,"./testdockerclient/client/cert.pem","./testdockerclient/client/key.pem","./testdockerclient/ca.pem")
//	fsouzaCli, _  = docker2.NewVersionedClient(host, "1.37")
//	fsouzaCli1, _ = dt.NewClientWithOpts(dt.WithHost(host), dt.WithVersion("1.37"))
//)
//
//func Test_listcontainers(t *testing.T) {
//	fliter := make(map[string][]string)
//	fliter["ancestor"] = []string{"hyperledger"}
//	fliter["status"] = []string{"running", "paused", "exited"}
//	svcContainers, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
//		Filters: fliter,
//	})
//	if err != nil {
//		logger.Error("ListContainers err: ", err)
//		return
//	}
//
//	for _, svcContainers := range svcContainers {
//		logger.Info(svcContainers.Names)
//		//err = fsouzaCli.RemoveContainer(dc.RemoveContainerOptions{ID: svcContainers.ID})
//		//if err != nil {
//		//	logger.Error("Error RemoveServiceRouter ", err)
//		//	return
//		//}
//	}
//	return
//}
//
//func RemoveContainer(containerName string) {
//	//获取容器列表
//	conlist, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
//		All: true,
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//
//	//删除容器
//	for _, value := range conlist {
//		logger.Info(value.Names[0])
//		//删除test容器
//		if value.Names[0] == "/"+containerName {
//			err = fsouzaCli.RemoveContainer(dc.RemoveContainerOptions{
//				Force: true,
//				ID:    value.ID,
//			})
//			if err != nil {
//				logger.Error(err)
//				return
//			}
//		}
//	}
//}
//
//func Test_ContainerStatus(t *testing.T) {
//	host = "http://192.168.9.87:2375"
//	fsouzaCli, _ := dt.NewClientWithOpts(dt.WithHost(host))
//
//	fsArgs := filters.NewArgs()
//
//	fsArgs.Add("name", "peer-1-baas3")
//
//	//获取容器列表
//	conlist, err := fsouzaCli.ContainerList(context.TODO(), types.ContainerListOptions{
//		Filters: fsArgs,
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	//logger.Info(conlist)
//	for _, con := range conlist {
//		logger.Info("con: ", con.Names)
//	}
//	if conlist[0].Names[0] != "peer-1-baas3" {
//		logger.Info("down")
//	}
//	if len(conlist) == 0 || conlist[0].State != "running" {
//		//logger.Info(conlist[0].State)
//		logger.Info(conlist)
//		logger.Info("down")
//		return
//	}
//	logger.Info("good")
//
//}
//
//func Test_updateconfig(t *testing.T) {
//	imgName := "192.168.9.87:5000/busybox"
//	containerName := "test"
//
//	RemoveContainer(containerName)
//
//	_, err = fsouzaCli.InspectImage(imgName + ":latest")
//	if err != nil {
//		logger.Info("pull new")
//		err = fsouzaCli.PullImage(docker2.PullImageOptions{
//			Repository: "192.168.9.87:5000/busybox",
//			Tag:        "latest",
//			Registry:   "192.168.9.87",
//		}, docker2.AuthConfiguration{})
//		if err != nil {
//			logger.Error(err)
//			return
//		}
//		logger.Info("good")
//	}
//
//	//创建容器
//	resp, err := fsouzaCli.CreateContainer(dc.CreateContainerOptions{
//		Name: containerName,
//		Config: &dc.Config{
//			Image: imgName,
//			Tty:   true,
//			Cmd:   []string{"sh", "-c", " sh"},
//		},
//		HostConfig: &dc.HostConfig{
//		},
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//
//	if err = fsouzaCli.StartContainer(resp.ID, nil); err != nil {
//		logger.Error(err)
//		panic(err)
//	}
//
//	container, err := fsouzaCli.InspectContainer(resp.ID)
//	logger.Info("container.HostConfig.ExtraHosts: ", container.HostConfig.ExtraHosts)
//
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//
//	////创建容器
//	e, err := fsouzaCli.CreateExec(dc.CreateExecOptions{
//		Container:  resp.ID,
//		Cmd:        []string{"sh", "-c", "echo '127.0.0.2 aaa' >> /etc/hosts"},
//		Privileged: true,
//		Tty:        true,
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartExec(e.ID, dc.StartExecOptions{
//		Detach: true,
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	//
//
//	container, err = fsouzaCli.InspectContainer(resp.ID)
//	logger.Info("container.HostConfig.ExtraHosts: ", container.HostConfig.ExtraHosts)
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	logger.Info("container: ", container)
//	//
//}
//
//func Test_main(t *testing.T) {
//	_, err = fsouzaCli.InspectImage("192.168.9.87:5000/busybox:latest")
//	if err != nil {
//		logger.Info("pull new")
//		err = fsouzaCli.PullImage(docker2.PullImageOptions{
//			Repository: "192.168.9.87:5000/busybox",
//			Tag:        "latest",
//			Registry:   "192.168.9.87:5000",
//		}, docker2.AuthConfiguration{})
//		if err != nil {
//			logger.Error(err)
//			return
//		}
//		logger.Info("good")
//	} else {
//		logger.Info("already exist")
//	}
//}
//
//func Test_main2(t *testing.T) {
//	_, err = fsouzaCli.InspectImage("192.168.9.87/busybox:latest")
//	if err != nil {
//		logger.Info("pull new")
//		err = fsouzaCli.PullImage(docker2.PullImageOptions{
//			Repository: "192.168.9.87:5000/busybox",
//			Tag:        "latest",
//			Registry:   "192.168.9.87",
//		}, docker2.AuthConfiguration{})
//		if err != nil {
//			logger.Error(err)
//			return
//		}
//		logger.Info("good")
//	} else {
//		logger.Info("already exist")
//	}
//}
//
////挂载方式1：创建mount的时候就把volume创建好
//func Test_volume1(t *testing.T) {
//	host = "http://192.168.9.61:2375"
//	nfshost := "192.168.9.82"
//	nfsdir := "/home/centos/testnfs/test"
//	fsouzaCli, _ = docker2.NewClient(host)
//	cname := "test1"
//	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
//		All: true,
//	})
//	for _, v := range cs {
//		if v.Names[0] == "/"+cname {
//			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
//				ID:            v.ID,
//				RemoveVolumes: true,
//				Force:         true,
//			})
//			if err != nil {
//				logger.Error(err)
//				return
//			}
//		}
//	}
//	//删除volume
//	vs, err := fsouzaCli.ListVolumes(docker2.ListVolumesOptions{})
//	for _, v := range vs {
//		err = fsouzaCli.RemoveVolumeWithOptions(docker2.RemoveVolumeOptions{
//			Name: v.Name,
//		})
//		if err != nil {
//			continue
//		}
//		logger.Info("removed ", v.Name)
//	}
//
//	//验证挂载
//	c, err := fsouzaCli.CreateContainer(docker2.CreateContainerOptions{
//		Name: cname,
//		Config: &docker2.Config{
//			Tty:   true,
//			Cmd:   []string{"sh", "-c", " sh"},
//			Image: "alpine",
//			Volumes: map[string]struct{}{
//				"/nfs": struct{}{},
//			},
//		},
//		HostConfig: &docker2.HostConfig{
//
//			Mounts: []docker2.HostMount{
//				//mount1
//				docker2.HostMount{
//					Type:   "volume",
//					Target: "/var/hyperledger",
//					VolumeOptions: &docker2.VolumeOptions{
//						DriverConfig: docker2.VolumeDriverConfig{
//							Name: "local",
//							Options: map[string]string{
//								"type":   "nfs",
//								"device": fmt.Sprintf("%s:%s", nfshost, nfsdir),
//								"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
//							},
//						},
//					},
//				},
//
//				docker2.HostMount{
//					Type:   "volume",
//					Target: "/nfs",
//					VolumeOptions: &docker2.VolumeOptions{
//						DriverConfig: docker2.VolumeDriverConfig{
//							Name: "local",
//							Options: map[string]string{
//								"type":   "nfs",
//								"device": fmt.Sprintf("%s:%s", nfshost, nfsdir),
//								"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
//							},
//						},
//					},
//				},
//
//				////mount3,和mount1的挂载device一样
//				//docker2.HostMount{
//				//	Type:   "volume",
//				//	Target: "/var/hyperledger/orderer2",
//				//	VolumeOptions: &docker2.VolumeOptions{
//				//		DriverConfig: docker2.VolumeDriverConfig{
//				//			Name: "local",
//				//			Options: map[string]string{
//				//				"type":   "nfs",
//				//				"device": fmt.Sprintf("%s:%s/orderer", nfshost, nfsdir),
//				//				"o":      fmt.Sprintf("addr=%s,rw,tcp,nolock", nfshost),
//				//			},
//				//		},
//				//	},
//				//},
//			},
//			Binds: []string{
//				//mike 宿主机上的地址进行映射
//				//"/nfs:/var/hyperledger/orderer:rw",
//				"/nfs/orderer:/var/hyperledger/orderer:rw",
//			},
//		},
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	logger.Info("good")
//}
//
////挂载方式2：先创建volume，再mount到容器内
//func Test_volume2(t *testing.T) {
//	host = "http://192.168.9.87:2375"
//	nfshost := "192.168.9.82"
//	fsouzaCli, _ = docker2.NewClient(host)
//	cname := "test1"
//	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
//		All: true,
//	})
//	for _, v := range cs {
//		if v.Names[0] == "/"+cname {
//			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
//				ID:            v.ID,
//				RemoveVolumes: true,
//				Force:         true,
//			})
//			if err != nil {
//				logger.Error(err)
//				return
//			}
//		}
//	}
//	//删除volume
//	vs, err := fsouzaCli.ListVolumes(docker2.ListVolumesOptions{})
//	for _, v := range vs {
//		err = fsouzaCli.RemoveVolumeWithOptions(docker2.RemoveVolumeOptions{
//			Name: v.Name,
//		})
//		if err != nil {
//			continue
//		}
//		logger.Info("removed ", v.Name)
//	}
//
//	//创建volume
//	vs1, err := fsouzaCli.CreateVolume(docker2.CreateVolumeOptions{
//		Name:   "docker_nfs",
//		Driver: "local",
//		DriverOpts: map[string]string{
//			"type":   "nfs",
//			"device": nfshost + ":/home/centos/testnfs/test",
//			"o":      "addr=" + nfshost + ",nolock,rw,tcp",
//		},
//	})
//
//	vs2, err := fsouzaCli.CreateVolume(docker2.CreateVolumeOptions{
//		Name:   "docker_nfs_common",
//		Driver: "local",
//		DriverOpts: map[string]string{
//			"type":   "nfs",
//			"device": nfshost + ":/home/centos/testnfs/test/common",
//			"o":      "addr=" + nfshost + ",nolock,rw,tcp",
//		},
//	})
//
//	//验证挂载
//	c, err := fsouzaCli.CreateContainer(docker2.CreateContainerOptions{
//		Name: cname,
//		Config: &docker2.Config{
//			Tty:   true,
//			Cmd:   []string{"sh", "-c", " sh"},
//			Image: "alpine",
//		},
//		HostConfig: &docker2.HostConfig{
//			Mounts: []docker2.HostMount{
//				docker2.HostMount{
//					Type:   "volume",
//					Source: vs1.Name,
//					Target: "/var/hyperledger/orderer",
//				},
//
//				docker2.HostMount{
//					Type:   "volume",
//					Source: vs2.Name,
//					Target: "/var/hyperledger/orderer/common",
//				},
//
//				docker2.HostMount{
//					Type:   "volume",
//					Source: vs1.Name,
//					Target: "/var/hyperledger/orderer1",
//				},
//			},
//		},
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	logger.Info("good")
//}
//
////docker run -d -p 8081:8081 --restart=always -v docker_nfs:/nfs --name=test alpine
//
//func Test_main1(t *testing.T) {
//	images, err := fsouzaCli.ListNetworks()
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	for _, v := range images {
//		logger.Info(v.ID)
//	}
//
//}
//
////创建容器并把宿主机的hosts挂在进去
//func Test_HostsVolume(t *testing.T) {
//	host = "http://192.168.9.87:2375"
//	fsouzaCli, _ = docker2.NewClient(host)
//	cname := "test1"
//	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
//		All: true,
//	})
//	for _, v := range cs {
//		if v.Names[0] == "/"+cname {
//			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
//				ID:            v.ID,
//				RemoveVolumes: true,
//				Force:         true,
//			})
//			if err != nil {
//				logger.Error(err)
//				return
//			}
//		}
//	}
//
//	//验证挂载
//	c, err := fsouzaCli.CreateContainer(docker2.CreateContainerOptions{
//		Name: cname,
//		Config: &docker2.Config{
//			Tty:   true,
//			Cmd:   []string{"sh", "-c", " sh"},
//			Image: "alpine",
//		},
//		HostConfig: &docker2.HostConfig{
//			Binds: []string{
//				"/home/ubuntu/nfs/hosts1:/etc/hosts1:rw", //绑定文件
//			},
//			UsernsMode: "host",
//		},
//	})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	logger.Info("good")
//}
//
//func Test_SaveCreateContainerYaml(t *testing.T) {
//	host = "http://192.168.9.87:2375"
//	fsouzaCli, _ = docker2.NewClient(host)
//	cname := "test1"
//	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
//		All: true,
//	})
//	for _, v := range cs {
//		if v.Names[0] == "/"+cname || v.Names[0] == "/aaa" {
//			err = fsouzaCli.RemoveContainer(docker2.RemoveContainerOptions{
//				ID:            v.ID,
//				RemoveVolumes: true,
//				Force:         true,
//			})
//			if err != nil {
//				logger.Error(err)
//				return
//			}
//		}
//	}
//	opt := docker2.CreateContainerOptions{
//		Name: cname,
//		Config: &docker2.Config{
//			Tty:   true,
//			Cmd:   []string{"sh", "-c", " sh"},
//			Image: "alpine",
//			Volumes: map[string]struct{}{
//				"/nfs": struct{}{},
//			},
//		},
//		HostConfig: &docker2.HostConfig{},
//	}
//
//	//验证挂载
//	c, err := fsouzaCli.CreateContainer(opt)
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartContainer(c.ID, &docker2.HostConfig{})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//
//	err = saveDockerContainerYaml(opt)
//	if err != nil {
//		logger.Error(err)
//	}
//
//	logger.Info("good")
//	opt1 := readContainerFromYaml("conf/docker-yaml-file/a-container.yaml")
//	opt1.HostConfig.ExtraHosts = append(opt1.HostConfig.ExtraHosts, "host4:127.0.0.4")
//	c1, err := fsouzaCli.CreateContainer(*opt1)
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//	err = fsouzaCli.StartContainer(c1.ID, &docker2.HostConfig{})
//	if err != nil {
//		logger.Error(err)
//		return
//	}
//}
//

//
////list 容器
//func Test_listcontainers1(t *testing.T) {
//	host = "http://192.168.9.82:2375"
//	fsouzaCli, _ = docker2.NewClient(host)
//	cs, _ := fsouzaCli.ListContainers(docker2.ListContainersOptions{
//		All: true,
//	})
//	for _, v := range cs {
//		logger.Info("container: ", v.Names)
//
//	}
//}
//
////list swarm集群容器
//func Test_listnodes(t *testing.T) {
//	host = "http://127.0.0.1:2375"
//	fsouzaCli, err := docker2.NewClient(host)
//	if err != nil {
//		panic(err)
//	}
//	cs, err := fsouzaCli.ListNodes(docker2.ListNodesOptions{
//	})
//	if err != nil {
//		panic(err)
//	}
//	for _, v := range cs {
//		logger.Info("address: ", v.Status.Addr)
//	}
//
//	ns, err := fsouzaCli.ListNetworks()
//	if err != nil {
//		panic(err)
//	}
//	for _, v := range ns {
//		logger.Info("address: ", v.Name)
//	}
//}
package main_test

import (
	docker2 "github.com/fsouza/go-dockerclient"
	"testing"
	"fmt"

	"strings"
	"io/ioutil"
	"github.com/ghodss/yaml"
	"os"
	"net/url"
	"github.com/docker/docker/api/types/swarm"
)

func Test_createservice(t *testing.T) {
	host := "http://192.168.9.82:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	fsouzaCli.CreateService(docker2.CreateServiceOptions{
		ServiceSpec: swarm.ServiceSpec{
			TaskTemplate: swarm.TaskSpec{
				ContainerSpec: &swarm.ContainerSpec{
					Image: "nginx:alpine",
					TTY:   true,
				},
			},
			Annotations: swarm.Annotations{
				Name: "test",
			},
			Networks: []swarm.NetworkAttachmentConfig{
				swarm.NetworkAttachmentConfig{Target: "wasabi-overlay"},
			},
			EndpointSpec: &swarm.EndpointSpec{Mode: swarm.ResolutionModeVIP, Ports: []swarm.PortConfig{
				swarm.PortConfig{
					Name:       "InterPort",
					TargetPort: uint32(7000),
				},
			}},
		},
	})
}

func Test_info(t *testing.T) {
	host := "http://192.168.9.87:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	info, err := fsouzaCli.Info()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", info.Swarm)
	fmt.Println(fsouzaCli.Endpoint())
	u, _ := url.Parse(fsouzaCli.Endpoint())
	fmt.Println(u.Hostname())
}

//list swarm集群容器
func Test_listnodes(t *testing.T) {
	host := "http://192.168.9.82:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	cs, err := fsouzaCli.ListNodes(docker2.ListNodesOptions{
	})
	if err != nil {
		panic(err)
	}
	for _, v := range cs {
		fmt.Println("address: ", v.Status.Addr)
	}
}

//list docker容器
func Test_listnodestatus(t *testing.T) {
	host := "http://192.168.9.87:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	filter := make(map[string][]string)
	filter["name"] = []string{"/peer-0-baas1"}
	cs, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
		Filters: filter,
		All:     true,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range cs {
		fmt.Println("address: ", v.State, v.Names)
	}
}

func Test_listservicestatus(t *testing.T) {
	host := "http://192.168.9.87:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	filter := make(map[string][]string)
	filter["name"] = []string{"peer-0-baas1"}
	cs, err := fsouzaCli.ListServices(docker2.ListServicesOptions{
		Filters: filter,
	})
	if err != nil {
		panic(err)
	}
	for _, v := range cs {
		fmt.Printf("v: %#v\n", int(*v.Spec.Mode.Replicated.Replicas))
	}

	detailService, err := fsouzaCli.InspectService(cs[0].ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("v: \n", int(*detailService.Spec.Mode.Replicated.Replicas))
}

//func (sm *SwarmClient) GetTaskStatus(org, nodename string) (string, error) {
//	value := fmt.Sprintf("nodename=%s", nodename)
//	client, err := GetRandomManagerClient()
//	if err != nil {
//		return "", err
//	}
//	tasks, err := client.ListTasks(dc.ListTasksOptions{
//		Filters: map[string][]string{
//			"label": []string{value},
//		},
//	})
//	if err != nil {
//		return "", fmt.Errorf("Error GetTaskList, err: %s", err)
//	}
//	sort.Sort(sortByTimeForSwarm(tasks))
//	if len(tasks) <= 0 {
//		return "down", nil
//	}
//	return strings.ToLower(string(tasks[0].Status.State)), nil
//}

func Test_Exec(t *testing.T) {
	host := "http://192.168.9.87:2375"
	fsouzaCli, err := docker2.NewVersionedClient(host, "1.37")
	if err != nil {
		panic(err)
	}
	// 获取所有的容器
	allContainers, err := fsouzaCli.ListContainers(docker2.ListContainersOptions{
		All: true,
	})
	if err != nil {
		panic(err)
		return
	}

	yamlFormat := "%s-container.yaml"
	savePath := "conf/docker-yaml-file/"
	// 每一个orderer容器进行stop start操作
	for _, c := range allContainers {
		//fmt.Println(c.Names)
		if strings.Contains(c.Names[0], "orderer-0") {
			cName := strings.Split(c.Names[0], "/")
			fileName := fmt.Sprintf(yamlFormat, cName[1])
			fmt.Println("read file: ", savePath+fileName)
			opt, err := readContainerFromYaml(savePath + fileName)
			if err != nil {
				panic(err)
				return
			}
			opt.HostConfig.ExtraHosts = append(opt.HostConfig.ExtraHosts, fmt.Sprintf("%s:%s", "666", "abc"))

			command, err := fsouzaCli.CreateExec(docker2.CreateExecOptions{
				Container:  c.ID,
				Cmd:        []string{"sh", "-c", fmt.Sprintf("echo '%s' >> /etc/hosts", fmt.Sprintf("%s %s", "666", "aaa"))},
				Privileged: true,
				Tty:        true,
			})
			if err != nil {
				panic(err)
			}
			err = fsouzaCli.StartExec(command.ID, docker2.StartExecOptions{})
			if err != nil {
				panic(err)
			}

			err = saveDockerContainerYaml(*opt, fileName)
			if err != nil {
				panic(err)
			}
		}
	}
}

func readContainerFromYaml(yamlFile string) (*docker2.CreateContainerOptions, error) {
	data, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		panic(err)
		return nil, err
	}
	opt := &docker2.CreateContainerOptions{}
	err = yaml.Unmarshal(data, opt)
	if err != nil {
		panic(err)
		return nil, err
	}

	opt.Name = "aaa"

	return opt, nil
}

func saveDockerContainerYaml(conConfig docker2.CreateContainerOptions, fileName string) error {
	fmt.Println("in saveDockerContainerYaml")
	savePath := "conf/docker-yaml-file/"

	fileContent, err := yaml.Marshal(conConfig)
	if err != nil {
		return fmt.Errorf("yaml [%s] marshal failed %s", conConfig.Name, err.Error())
	}
	fmt.Println(os.Getwd())
	_, err = os.Stat(savePath)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(savePath, 0777)
		if err != nil {
			return fmt.Errorf("Create dir [conf/docker-yaml-file] failed %s", fileName, err.Error())
		}
	} else if err != nil {
		return fmt.Errorf("Create dir [conf/docker-yaml-file] failed %s", fileName, err.Error())
	}
	file, err := os.OpenFile(savePath+fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Create yaml file [%s] failed %s", fileName, err.Error())
	}
	_, err = file.Write(fileContent)
	if err != nil {
		return fmt.Errorf("Write yaml file [%s] failed %s", fileName, err.Error())
	}
	fmt.Printf("Create yaml file [%s] Success!", fileName)
	return nil
}
