package main

import (
	"context"
	"fmt"
	"github.com/fsouza/go-dockerclient"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	dockercli, err := docker.NewClient("unix:///var/run/docker.sock")
	if err != nil {
		fmt.Println(err)
		return
	}

	preDelete(dockercli)

	contaier, err := dockercli.CreateContainer(docker.CreateContainerOptions{
		Name: "testetcd",
		Config: &docker.Config{
			Cmd: []string{
				"/usr/local/bin/etcd",
				"--name", "etcd1",
				"--data-dir", "/etcd-data",
				"--listen-client-urls", "http://0.0.0.0:2379",
				"--advertise-client-urls", "http://127.0.0.1:2379",
				"--listen-peer-urls", "http://0.0.0.0:2380",
				"--initial-advertise-peer-urls", "http://127.0.0.1:2380",
				"--initial-cluster", "etcd1=http://127.0.0.1:2380",
				"--initial-cluster-token", "etcd-cluster-1",
				"--initial-cluster-state", "new",
			},
			Image:        "etcd:v3.3.9",
			ExposedPorts: map[docker.Port]struct{}{"2379/tcp": {}, "2380/tcp": {}},
		},
		HostConfig: &docker.HostConfig{
			PortBindings: map[docker.Port][]docker.PortBinding{
				"2379/tcp": []docker.PortBinding{
					docker.PortBinding{
						HostIP:   "localhost",
						HostPort: "2379",
					},
				},

				"2380/tcp": []docker.PortBinding{
					docker.PortBinding{
						HostIP:   "localhost",
						HostPort: "2380",
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = dockercli.StartContainer(contaier.ID, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(contaier.ID)

	config := clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"},
		DialTimeout: 10 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	fmt.Println(client.Status(context.Background(), "http://localhost:2379"))
	kv := clientv3.NewKV(client)
	fmt.Println(client.Endpoints())
	ctx, cancleFunc := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancleFunc()
	fmt.Println("startput")
	putResp, err := kv.Put(ctx, "/job/v3", "push the box3", clientv3.WithPrevKV()) //withPrevKV()是为了获取操作前已经有的key-value
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("stopput")
	fmt.Printf("%v\n", putResp.PrevKv)
	putResp, err = kv.Put(ctx, "/job/v4", "push the box4", clientv3.WithPrevKV()) //withPrevKV()是为了获取操作前已经有的key-value
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", putResp.PrevKv)
	getResp, err := kv.Get(ctx, "/job/", clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", getResp.Kvs)

	wc := client.Watch(context.Background(), "/job/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	for v := range wc {
		if v.Err() != nil {
			panic(err)
		}
		for _, e := range v.Events {
			fmt.Printf("type:%v\n kv:%v  prevKey:%v \n", e.Type, e.Kv, e.PrevKv)
		}
	}
}

func preDelete(cli *docker.Client) {
	fmt.Println("start preDelete")
	conlist, err := cli.ListContainers(docker.ListContainersOptions{
		All: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(conlist)
	//删除容器
	for _, value := range conlist {
		//删除test容器
		if value.Names[0] == "/testetcd" {
			err = cli.RemoveContainer(docker.RemoveContainerOptions{
				ID:    value.ID,
				Force: true,
			})
			if err != nil {
				return
			}
			fmt.Println("remove good")
		}
	}
}
