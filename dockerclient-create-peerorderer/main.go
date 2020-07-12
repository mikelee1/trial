package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	docker2 "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/sdk"
	"github.com/op/go-logging"
	"myproj.lee/try/dockerclient-create-peerorderer/fabric"
	"net/url"
	"os"
	"path"
	_ "path"
	"time"
)

const (
	NFS                = "nfs"
	LOCAL              = "local"
	ContainerMountPath = "/var/hyperledger"
	host               = "http://192.168.9.82:2375"
	endpoint           = "http://192.168.9.82:2375"
)

var logger *logging.Logger

func init() {
	logger = logging.MustGetLogger("docker-client")
}

var (
	org    string
	mspDir string
	gm     bool
	orgCA  *sdk.CA
	kafkas = []string{"192.168.9.82:9092", "192.168.9.82:9192", "192.168.9.82:9292", "192.168.9.82:9392"}
)

// Initiator ...
type Initiator struct {
	org     string
	orgMSP  string
	orgCA   *sdk.CA
	support service.InitiatorSupport

	client *sdk.Client
}

func main() {
	logger.Info("start main")

	//生成docker句柄
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cli, err := client.NewClientWithOpts(client.WithHost(host))
	if err != nil {
		panic(err)
	}
	cli.NegotiateAPIVersion(ctx)

	logger.Info("start ContainerList")
	//获取容器列表
	conlist, err := cli.ContainerList(ctx, types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	//删除容器
	for _, value := range conlist {
		//删除test容器
		if value.Names[0] == "/peer-0-baas2" {
			err = cli.ContainerRemove(ctx, value.ID, types.ContainerRemoveOptions{
				Force: true,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}


	resp, err := cli.ContainerCreate(ctx, &container.Config{
		//Env:peerEnv,
		Tty:   true,
		Image: "hyperledger/fabric-peer:amd64-latest",
		Cmd:   []string{"/bin/sh", "-c", "ln -s & peer node start"},
	},
		&container.HostConfig{
			Binds: []string{"/var/lib/docker/volumes/a92f8bd649169bf8da0e45eed8e395ea1f812a5cf3cec447984543bc9d438bda/_data:/tmp"},
			Mounts: []mount.Mount{
				mount.Mount{
					Type:     "bind",
					Source:   "/var/run",
					Target:   "/host/var/run",
					ReadOnly: false,
				},
				mount.Mount{
					Type:     "volume",
					Source:   "todoooooooooooo",
					Target:   ContainerMountPath,
					ReadOnly: false,
				},
			},
			UsernsMode: "host",
			ExtraHosts: []string{"host1:192.168.9.82"},
		}, &network.NetworkingConfig{}, "peer-0-baas2")
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("start container")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
		return
	}

	//exec,err := cli.ContainerExecCreate(ctx,resp.ID,types.ExecConfig{
	//	Cmd:[]string{"mkdir /lee & ls"},
	//	AttachStdin:true,
	//	Tty:true,
	//})
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}
	//err = cli.ContainerExecStart(ctx,exec.ID,types.ExecStartCheck{})
	//if err != nil {
	//	logger.Error(err)
	//	return
	//}

	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		Timestamps: true,
		Details:    true,
		Follow:     true,
	})
	if err != nil {
		panic(err)
	}

	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

}

// NewInitiator ...
func NewInitiator(org string, orgMSP string, orgCA *sdk.CA, support service.InitiatorSupport, gm bool) (*Initiator, error) {
	client, err := sdk.NewClient(orgCA.AdminCommonName(), orgMSP, orgCA.AdminMSPDir(), gm)
	if err != nil {
		logger.Info("Error creating client for org", err)
		return nil, err
	}

	return &Initiator{
		org:    org,
		orgMSP: orgMSP,
		orgCA:  orgCA,
		//support: support,
		client: client,
	}, nil

}

func createInitiator() (*Initiator, error) {

	orgCA, err := service.GetCA(path.Join(mspDir, org), org)
	if err != nil {
		logger.Error("Error getting peer ca", err)
		return nil, err
	}
	//todo k8s 仅供测试 后续选择底层可能是前端传入

	var client *Docker
	client, err = NewDockerClient([]string{endpoint})

	if err != nil {
		logger.Error("Error getting constructor", err)
	}
	support := struct {
		Docker
		service.Storge
	}{*client, nil}

	return NewInitiator(org, org, orgCA, support, gm)
}

type Docker struct {
	client    *docker2.Client
	svcs      []string
	confs     []string
	endpoints []string
}

type raftmember struct {
	ID       string
	RaftPort uint32
}

// NewDockerClient ...
func NewDockerClient(endpoints []string) (*Docker, error) {
	client, err := docker2.NewClient(endpoints[0])
	if err != nil {
		logger.Error("Error creating docker client", err)
		return nil, err
	}
	var eps []string
	for _, ep := range endpoints {
		u, err := url.Parse(ep)
		if err != nil {
			logger.Error("Error parsing docker endpoint", err)
			return nil, err
		}
		eps = append(eps, u.Hostname())

	}
	return &Docker{
		client:    client,
		endpoints: eps,
	}, nil
}
