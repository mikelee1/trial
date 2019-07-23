package main

import (
	"context"
	"fmt"
	docker2 "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/sdk"
	"github.com/op/go-logging"
	"myproj/try/dockerclient-create-peerorderer/fabric"
	"net/url"
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

func prepare() {
	org = "baas1"
	mspDir = "msp/"
	gm = true

	//orgCA, err := service.GetCA(path.Join(mspDir, org), org)
	//if err != nil {
	//	logger.Error("Error getting peer ca", err)
	//}
}

// Initiator ...
type Initiator struct {
	org     string
	orgMSP  string
	orgCA   *sdk.CA
	support service.InitiatorSupport

	client *sdk.Client
}

//func NewInitiator() *Initiator {
//	fabricClient, err := sdk.NewClient(orgCA.AdminCommonName(), org, orgCA.AdminMSPDir(), gm)
//	if err != nil {
//		logger.Error(err)
//		return nil
//	}
//	return &Initiator{
//		org:     org,
//		orgMSP:  org,
//		orgCA:   orgCA,
//		client:fabricClient,
//	}
//}

func main() {
	logger.Info("start main")
	prepare()
	//initiator := NewInitiator()

	//生成docker句柄
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	fmt.Println(ctx)
	cli, err := NewDockerClient([]string{host})
	if err != nil {
		panic(err)
	}
	logger.Info("start ContainerList")
	//获取容器列表
	conlist, err := cli.client.ListContainers(docker2.ListContainersOptions{
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
			err = cli.client.RemoveContainer(docker2.RemoveContainerOptions{
				ID:    value.ID,
				Force: true,
			})
			if err != nil {
				logger.Error(err)
				return
			}
		}
	}

	peers := []service.PeerPorts{
		service.PeerPorts{
			Main:      30031,
			Chaincode: 30032,
			Event:     30033,
		},
		service.PeerPorts{
			Main:      30034,
			Chaincode: 30035,
			Event:     30036,
		},
	}
	orderer0 := service.OrdererPorts{
		Main:  30020,
		Raft:  30021,
		Debug: 30022,
	}
	setupReq := service.SetupRequest{
		OrgName:      "baas1",
		PeerPorts:    peers,
		OrdererPorts: []service.OrdererPorts{orderer0},
		Consensus:    "kafka",
	}
	bi := service.GetBuildInfo(&setupReq)
	//logger.Info(bi.String())
	initiator, err := createInitiator()
	if err != nil {
		logger.Error(err)
		return
	}
	block, err := service.CreateGenesisBlockData(initiator.orgMSP, initiator.orgCA, bi, kafkas, bi.Consensus)
	logger.Info(string(block))

	resp, err := cli.client.CreateContainer(docker2.CreateContainerOptions{
		Config: &docker2.Config{
			ExposedPorts: map[docker2.Port]struct{}{
				docker2.Port("17050"): struct{}{},
			},
			Tty:   true,
			Image: "hyperledger/fabric-peer:amd64-latest",
			//Cmd:   []string{"/bin/sh","-c","ln -s & peer node start"},
			Cmd: []string{"sh"},
		},
		HostConfig: &docker2.HostConfig{
			PortBindings: map[docker2.Port][]docker2.PortBinding{
				"17050/tcp": []docker2.PortBinding{{
					HostIP:   "0.0.0.0",
					HostPort: "7050",
				}},
			},
			Binds: []string{"/var/lib/docker/volumes/a92f8bd649169bf8da0e45eed8e395ea1f812a5cf3cec447984543bc9d438bda/_data:/tmp"},
			Mounts: []docker2.HostMount{
				docker2.HostMount{
					Type:     "bind",
					Source:   "/var/run",
					Target:   "/host/var/run",
					ReadOnly: false,
				},
				docker2.HostMount{
					Type:     "volume",
					Source:   "todoooooooooooo",
					Target:   ContainerMountPath,
					ReadOnly: false,
				},
			},
			UsernsMode: "host",
			ExtraHosts: []string{"host1:192.168.9.82"},
		},
		NetworkingConfig: &docker2.NetworkingConfig{
			EndpointsConfig: map[string]*docker2.EndpointConfig{},
		},
		Name: "peer-0-baas2",
	})
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("start container")
	if err := cli.client.StartContainer(resp.ID, &docker2.HostConfig{}); err != nil {
		panic(err)
		return
	}
	fmt.Println("awit container")
	statusCh, errCh := cli.client.WaitContainer(resp.ID)
	fmt.Println("after waite")
	fmt.Println(statusCh, errCh)

	err = cli.client.Logs(docker2.LogsOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(err)

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

//type raftmember struct {
//	ID       string
//	RaftPort uint32
//}

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
