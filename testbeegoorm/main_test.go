package main_test

import (
	"testing"
	"myproj.lee/try/testbeegoorm/models"
	"fmt"
	"time"
)

//修改orgchannel表
func Test_UpdateOrgChannel(t *testing.T) {
	const (
		PeerMainPort       = "peer main port"
		PeerDebugBlockPort = "peer block debug port"
		PeerChainCodePort  = "peer chaincode port"
		OrderMainPort      = "order main port"
		CouchDBMainPort    = "couchdb main port"

		K8sMinNodePort = 30000
		K8sMaxNodePort = 32767
	)
	tmp := &models.ChaincodeInfo{
		Id: 1,
	}
	fmt.Println("update")
	fmt.Println(tmp.Id)
	tmp.Channel = "channel3"
	a, err := models.GetDBClient().Update(tmp, "Channel")
	fmt.Println(a, err)
}

//修改orgchannel表
func Test_queryOrg(t *testing.T) {
	tmp := &models.Org{
		UserId: 1,
	}
	tmp1 := &models.Org{
		Id:     22,
		UserId: 1,
	}
	if _, err := models.GetDBClient().Insert(tmp1); err != nil {
		fmt.Println(err)
		return
	}

	if err := models.GetDBClient().QueryTable("org").One(tmp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmp.Id)
}

func Test_addInform(t *testing.T) {
	tmp := &models.Inform{InformDetail: models.InformDetail{Title: "快乐圣诞节开发", Time: time.Now(), Informdata: "S"}}
	if _, err := models.GetDBClient().Insert(tmp); err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(tmp.Id)
}

func Test_querynodestatus(t *testing.T) {
	if models.GetDBClient().QueryTable("node").Filter("name", "peer-0-baas3").Filter("is_peer", 1).Filter("status", models.NodeRunningStatus).Exist() {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}

func Test_findNode(t *testing.T) {
	node := &models.Node{}
	node.Name = "peer-0-baas1"
	if err := models.GetDBClient().Read(node, "name"); err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("good")
}

func Test_GetNodesFromLocalDB(t *testing.T) {
	nodes, _ := GetNodesFromLocalDB()
	for _, node := range nodes {
		fmt.Println(node.Name)
	}
	return
}

func GetNodesFromLocalDB() ([]*models.Node, error) {
	nodes := []*models.Node{}
	_, err := models.GetDBClient().Raw("select * from node").QueryRows(&nodes)
	if err != nil {
		return nil, err
	}
	return nodes, nil
}
