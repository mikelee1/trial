package main

import (
	//"github.com/docker/docker/client"
	dc "github.com/fsouza/go-dockerclient"
	"fmt"
	"github.com/docker/docker/api/types/swarm"
	"github.com/astaxie/beego"
)

var (
	cli  *dc.Client
	host = "tcp://192.168.9.87:2375"
	err  error
)

func init() {
	cli, err = dc.NewClient(host)
	if err != nil {
		panic(err)
	}
}

//docker volume create \
//--driver local \
//--opt type=nfs \
//--opt o=addr=192.168.9.82,nolock,rw,tcp \
//--opt device=192.168.9.82:/home/centos/testnfs/test \
//--name wasabi_nfs
func main() {
	//fmt.Println(cli.Info(context.TODO()))

	resp, err := cli.CreateService(
		dc.CreateServiceOptions{

			ServiceSpec: swarm.ServiceSpec{
				Annotations: swarm.Annotations{
					Name: "aaa2",
				},
				EndpointSpec: &swarm.EndpointSpec{
					Mode: swarm.ResolutionModeVIP,
					Ports: []swarm.PortConfig{
						swarm.PortConfig{
							Name:          "MainPort",
							TargetPort:    uint32(1113),
							PublishedPort: uint32(1112),
							PublishMode:   swarm.PortConfigPublishModeHost,
						},
					},
				},
				TaskTemplate: swarm.TaskSpec{
					ContainerSpec: &swarm.ContainerSpec{
						TTY:   true,
						Image: "alpine",
						Env: []string{
							fmt.Sprintf("COUCHDB_USER=%s", beego.AppConfig.String("CouchDBUsername")),
							fmt.Sprintf("COUCHDB_PASSWORD=%s", beego.AppConfig.String("CouchDBPassword")),
						},
					},
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.ID)
}
