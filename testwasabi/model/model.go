package model

import "github.com/hyperledger/fabric/sdk"

const (
	AcceptAllPolicy             = "OutOf(0, 'None.member')"
)

type OrgInfo struct {
	OrgName      string
	MspID        string
	OrgMSP       string
	OrgCA        *sdk.CA
	Client       *sdk.Client
	PeerNodes    []*ServiceNode
	OrdererNodes []*ServiceNode
}

type NewCreateChannelRequest struct {
	Orgs        []*OrgInfo
	ChannelName string
}

type InvokeRequest struct {
	Invoke       bool
	Org          string
	ChannelName  string
	CcName       string
	Args         [][]byte
	PeerNodes    []*ServiceNode
	OrdererNodes []*ServiceNode
}



type JoinChannelRequest struct {
	Orgs        []*OrgInfo
	ChannelName string
}

type IdentityRequest struct {
	Orgs []*OrgInfo
}

type InstallChainCodeRequest struct {
	Org       string
	CcTarPath string
	CcPath    string
	CcName    string
	CcVersion string
	PeerNodes []*ServiceNode
}

type ServiceNode struct {
	ID               string
	Endpoint         string
	ExternalEndpoint string
	Public           bool
}


type InstantiateChainCodeRequest struct {
	Org          string
	ChannelName  string
	CcName       string
	CcVersion    string
	Policy       string
	Args         [][]byte
	PeerNodes    []*ServiceNode
	OrdererNodes []*ServiceNode
}

type IndirectInvokeRequest struct {

	Org          string //直参机构的org
	ChannelName  string
	CcName       string
	Args         []string
	PeerNodes    []*ServiceNode
	OrdererNodes []*ServiceNode
}