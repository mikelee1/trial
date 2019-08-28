package models

type SetupRequest struct {
	// 网络名称
	NetworkName string `json:"networkName"`
	// 共识模式
	Consensus string `json:"consensus"`
	// 组织名称
	OrgName string `json:"orgName"`
	// 组织别名
	OrgAlias string `json:"orgAlias"`
	// peer 端口
	PeerPorts []PeerPorts `json:"peerPorts"`
	// orderer 端口
	OrdererPorts []OrdererPorts `json:"ordererPorts"`
	// 公司名称
	Company string `json:"company"`
	// 证书是否自动生成 现在只支持true
	AutoGeneratedCerts bool `json:"autoGeneratedCerts"`
}


type PeerPorts struct {
	Main      uint32 `json:"main"`
	Chaincode uint32 `json:"chaincode"`
}

type OrdererPorts struct {
	Main  uint32 `json:"main"`
	Debug uint32 `json:"debug"`
}


type IndirectInvokeRequest struct {
	Userid       string
	Org          string //直参机构的org
	ChannelName  string
	CcName       string
	Args         []string
	//PeerNodes    []*ServiceNode
	//OrdererNodes []*ServiceNode
}

type InitChannelRequest struct {
	Orgs        []string
	ChannelName string
	Peers       []string
}

type CreateIndirectRequest struct {
	Orgname  string `json:"orgname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	IsAdmin bool   `json:"isAdmin"`
	UserId  string `json:"userId"`

	HasCluster            bool   `json:"hasCluster"`
	HasCreatedNetwork     bool   `json:"hasCreatedNetwork"`
	HasJoinedNetwork      bool   `json:"hasJoinedNetwork"`
	HasArtifactsGenerated bool   `json:"hasArtifactsGenerated"`
	UserName              string `json:"userName"`
	OrgName               string `json:"orgName"`
}


type InstallAndInstantiateChainCodeRequest struct {
	CcName      string   `json:"cc_name"`
	CcVersion   string   `json:"cc_version"`
	Args        []string `json:"args"`
	ChannelName string   `json:"channel_name"`
	PeerNodes   []string `json:"peer_nodes"`
	CcPath      string   `json:"cc_path"`
	CcHash      string   `json:"cc_hash"`
	Org         string
}

type AuthUpdate struct {
	Indirectid string
	Items      map[string][]string
}