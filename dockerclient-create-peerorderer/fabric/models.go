package service

import (
	"encoding/json"
	"fmt"
	"bytes"
)

// Constructor ...
type Constructor interface {
	//ConstructBlockchain(orgMSP string, orgCA *sdk.CA, info *BuildInfo, genesisBlock []byte) (*BuildResult, error)
	//CreatePeers(orgMSP string, orgCA *sdk.CA, peersInfo []*PeerInfo) (*BuildResult, error)
	//CreateSeparatePeers(orgMSP string, orgCA *sdk.CA, peersInfo []*PeerInfo) (*BuildResult, error)
	//RemoveAll()
}

// Storge ...
type Storge interface {
	SaveBuildResult(*ResultPack) error
	SaveIdentityResult(*IdentityPack) error
	QueryPeersByOrg() ([]string, error)
	//SavePeerInfo(*ResultPack) error
}

// BuildResult contains node info after building
type BuildResult struct {
	MSPID        string
	TLSCert      []byte
	RootCert     []byte
	AdminCert    []byte
	PeerNodes    []*ServiceNode
	OrdererNodes []*ServiceNode
	AnchorPeers  []string
	Consensus    string
}

// BuildInfo ...
type BuildInfo struct {
	SeparatePeerInfo []*SeparatePeerInfo
	PeerInfo         []*PeerInfo
	OrdererInfo      []*OrdererInfo
	InviteCode       []byte
	Consensus        string
}

func (conf *BuildInfo) String() string {
	b, err := json.Marshal(*conf)
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "    ")
	if err != nil {
		return fmt.Sprintf("%+v", *conf)
	}
	return out.String()
}

// Orderers ...
func (bi *BuildInfo) Orderers() (orderers []string) {
	for _, info := range bi.OrdererInfo {
		orderers = append(orderers, info.MainPort.ExternalEndpoint)
	}
	return
}

// PeerInfo ...
type PeerInfo struct {
	ID            string
	Labels        map[string]string
	MainPort      PortMapping
	ChaincodePort PortMapping
	EventPort     PortMapping
	BlockPort     PortMapping
	Public        bool
	// Exist         bool //exist 表示外部已经创建的peer，不需要baas部署和管理
}
type SeparatePeerInfo struct {
	ID            string
	Labels        map[string]string
	MainPort      PortMapping
	ChaincodePort PortMapping
	EventPort     PortMapping
	BlockPort     PortMapping
	// ExternalIp    string
	Public bool
}
// OrdererInfo ...
type OrdererInfo struct {
	ID              string
	Labels          map[string]string
	MainPort        PortMapping
	RaftPort        PortMapping
	Debugserverport PortMapping
	RaftId          uint32
}
// PortMapping ...
type PortMapping struct {
	Port             uint32
	ExternalEndpoint string
}


// ResultPack ...
type ResultPack struct {
	BuildRes     *BuildResult
	GenesisBlock []byte
	PublicBlock  []byte
}

// ServiceNode ...
type ServiceNode struct {
	ID                string //存放的是peername，不是id。。
	Endpoint          string
	ExternalEndpoint  string
	ChaincodeEndpoint string
	EventEndpoint     string
	RaftEndpoint      string
	DebugEndpoint     string
	RaftId            uint32
	Public            bool
}


// InitiatorSupport ...
type InitiatorSupport interface {
	Constructor
	Storge
}



// IdentityPack ...
type IdentityPack struct {
	Org         string
	PeerInfo    string
	OrdererInfo string
	Consensus   string
}