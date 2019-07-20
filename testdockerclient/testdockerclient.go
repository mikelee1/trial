package main

import (
	dc "github.com/fsouza/go-dockerclient"
	"github.com/op/go-logging"
	"os"
	"myproj/try/common/fmtstruct"
)

var logger = &logging.Logger{}
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)
var (
	containerName = "test11"
	host = "http://192.168.9.82:2375"
)

func init() {
	logger = logging.MustGetLogger("testdockerclient")
	backend1 := logging.NewLogBackend(os.Stderr, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)

	//设置正常的输出format走backend2
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	//设置异常的输出走backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.ERROR, "")

	logging.SetBackend(backend1Leveled,backend2Formatter)
}

func main() {
	cli, err := dc.NewClient(host)
	if err != nil {
		panic(err)
		return
	}

	//获取容器列表
	conlist, err := cli.ListContainers(dc.ListContainersOptions{
		All: true,
	})
	if err != nil {
		logger.Error(err)
		return
	}

	//删除容器
	for _, value := range conlist {
		//删除test容器
		if value.Names[0] == "/"+containerName {
			err = cli.RemoveContainer(dc.RemoveContainerOptions{
				Force: true,
				ID:    value.ID,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}

	resp, err := cli.CreateContainer(dc.CreateContainerOptions{
		Name: containerName,
		Config: &dc.Config{
			Image: "alpine",
			Tty:   true,
			Cmd:   []string{"sh"},
		},
		HostConfig: &dc.HostConfig{
			Binds: []string{
				"/home/centos/go/src/wasabi:/home",               //绑定目录
				"/home/centos/go/src/test.txt:/home/test.txt:rw", //绑定文件,读写权限
			},
			AutoRemove: true,
			UsernsMode: "host",
			ExtraHosts: []string{"host1:192.168.9.82"},
		},
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(resp.ID)

	if err = cli.StartContainer(resp.ID, nil); err != nil {
		panic(err)
	}

	////暂停
	//time.Sleep(5 * time.Second)
	//err = cli.PauseContainer(containerName)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//logger.Info("pause well")
	//
	////启动
	//time.Sleep(5 * time.Second)
	//err = cli.UnpauseContainer(containerName)
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//logger.Info("unpause well")



	dcInfo,err := cli.Info()
	logger.Info(fmtstruct.String(dcInfo))

}


