package main

import (
	"go.etcd.io/etcd/clientv3"
	"time"
	"fmt"
	"context"
)

func main() {
	//设置etcd的config参数
	config := clientv3.Config{
		Endpoints:[]string{"127.0.0.1:2379"},
		DialTimeout:10*time.Second,
	}

	//构建etcd的client
	client,err := clientv3.New(config)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	//获得kv结构体
	kv := clientv3.NewKV(client)
	ctx,cancleFunc:= context.WithTimeout(context.TODO(),5*time.Second)
	defer cancleFunc()

	//存储key,value
	putResp,err := kv.Put(ctx,"/job/v3","push the box3",clientv3.WithPrevKV())  //withPrevKV()是为了获取操作前已经有的key-value
	if err != nil{
		panic(err)
	}
	fmt.Printf("%v\n",putResp.PrevKv)

	//存储key，value
	putResp,err = kv.Put(ctx,"/job/v4","push the box4",clientv3.WithPrevKV())  //withPrevKV()是为了获取操作前已经有的key-value
	if err != nil{
		panic(err)
	}
	fmt.Printf("%v\n",putResp.PrevKv)

	//获取value
	getResp,err := kv.Get(ctx,"/job/",clientv3.WithPrefix()) //withPrefix()是未了获取该key为前缀的所有key-value
	if err != nil{
		panic(err)
	}
	fmt.Printf("%v\n",getResp.Kvs)
}
