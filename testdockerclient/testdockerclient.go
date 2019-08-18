package main

import (
	"context"
	"github.com/docker/docker/client"
	dc "github.com/fsouza/go-dockerclient"
	"github.com/op/go-logging"
	logger2 "myproj/try/common/logger"
	"github.com/docker/docker/api/types/volume"
)

var logger = &logging.Logger{}

var (
	containerName  = "test12"
	//host           = "http://192.168.9.83:2375"
	//standardCli, _ = client.NewClientWithOpts(client.WithHost(host))
	//fsouzaCli, _   = dc.NewClient(host)

	host           = "http://192.168.9.82:2375"
	standardCli, _ = client.NewClientWithOpts(client.WithHost(host),client.WithTLSClientConfig("./testdockerclient/ca.pem", "./testdockerclient/client/cert.pem","./testdockerclient/client/key.pem"))
	fsouzaCli, _   = dc.NewTLSClient(host,"./testdockerclient/client/cert.pem","./testdockerclient/client/key.pem","./testdockerclient/ca.pem")
)

func init() {
	logger = logger2.GetLogger()
}



func main() {
	standardCli.VolumeCreate(context.TODO(),volume.VolumeCreateBody{})
	var err error
	err = fsouzaCli.PullImage(dc.PullImageOptions{
		Repository:"192.168.9.83:5000/busybox",
		Tag:"latest",
		Registry:"http://192.168.9.83:5000",
	},dc.AuthConfiguration{})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("good")

	//获取容器列表
	conlist, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
		All: true,
	})
	if err != nil {
		logger.Error(err)
		return
	}

	_, err = standardCli.Info(context.TODO())
	if err != nil {
		logger.Error(err)
		return
	}

	_,err = fsouzaCli.InspectImage("alpine")
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
	//创建容器
	resp, err := fsouzaCli.CreateContainer(dc.CreateContainerOptions{
		Name: containerName,
		Config: &dc.Config{
			Image:      "alpine",
			Tty:        true,
			Cmd:        []string{"sh", "-c", "wget http://test11:18000 & sh"},
			WorkingDir: "/home",
			Volumes: map[string]struct{}{
				"/var/hyperledger": struct {}{},
				"/home": struct {}{},
			},
			Mounts:[]dc.Mount{
				dc.Mount{
					Source:"lee",
					Destination:"/home",
					RW:true,
					Mode:"rprivate",
					Driver:"rw",
				},
			},
		},
		HostConfig: &dc.HostConfig{
			Binds: []string{
				"/home/centos/go/src/wasabi:/home",               //绑定目录
				"/home/centos/go/src/test.txt:/home/test.txt:rw", //绑定文件,读写权限
			},
			AutoRemove: true,
			UsernsMode: "host",
			ExtraHosts:[]string{"test11:192.168.9.83"},
			//Links:[]string{"test11:test11"},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(resp.ID)

	if err = fsouzaCli.StartContainer(resp.ID, nil); err != nil {
		panic(err)
	}

	////暂停
	//time.Sleep(5 * time.Second)
	//err = fsouzaCli.PauseContainer(containerName)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//logger.Info("pause well")
	//
	////启动
	//time.Sleep(5 * time.Second)
	//err = fsouzaCli.UnpauseContainer(containerName)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//logger.Info("unpause well")


	////容器状态
	//logger.Error(resp.ID)
	//cs, err := fsouzaCli.InspectContainer(resp.ID)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//logger.Info(cs.State.Status)

	//dcInfo,err := fsouzaCli.Info()
	//logger.Info(fmtstruct.String(dcInfo))

	////根据容器名获取容器状态
	//a, err := fsouzaCli.ListContainers(dc.ListContainersOptions{
	//	Filters: map[string][]string{"name": []string{containerName}},
	//})
	//logger.Info(fmtstruct.String(a))
	//
	////根据容器名获取容器状态
	//fsArgs := filters.NewArgs()
	//fsArgs.Add("name", containerName)
	//fcl, err := standardCli.ContainerList(context.TODO(), types.ContainerListOptions{
	//	Filters: fsArgs,
	//})
	//logger.Info(fmtstruct.String(fcl))
}
