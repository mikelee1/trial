package models

import "time"

type ChaincodeInfo struct {
	Id                  int
	UserId              string //用户
	Name                string `orm:"index"` //名称
	Version             string //版本
	Channel             string //所属链
	NodeChaincodeStatus string //base64 peer节点的chiancode状态 已上传->(->安装失败)已安装->(->实例化失败)已实例化
	Status              string //合约状态   running 部署完成运行中 /  unDeployed 未部署 /install 已安装
	Createdtime         int64  //创建时间
	Createdorg          string //创建组织
	Hash                string //合约文件的hash
	CcPath              string //合约安装路径
}

type Auth struct {
	Id        int `orm:"pk;auto"`
	Userid    string
	Channel   string
	Chaincode string
}

type OrgChannel struct {
	Id         int
	Channel    string `orm:"unique"`
	Createorg  string
	Createtime string
	Currentorg string
}

type Org struct {
	Id           int
	UserId       int
	MspId        string
	GenesisBlock string `orm:"type(text)"`
	PublicBlock  string `orm:"type(text)"`
	TlsCert      string `orm:"type(text)"`
	RootCert     string `orm:"type(text)"`
	AdminCert    string `orm:"type(text)"`
	AnchorPeers  string
	Consensus    string
	Namespace    string
}

type Inform struct {
	InformDetail
	AcceptGroup string //已同意的组织集合 strings join with ","
	RejectGroup string //已拒绝的组织集合 strings join with ","
	UnDoGroup   string //未操作的组织集合 strings join with ","
}
type InformDetail struct {
	Id          int
	Title       string    //标题
	Content     string    //内容
	Route       string    //路由
	SenderId    string    //发件人
	ReceiverId  string    //收件人，操作时也就是signer
	InviterId   string    //邀请人
	InviteeId   string    //被邀请人
	ChainId     string    //chainid
	Unread      int       //是否已读/未读/操作 0:已读,1:未读,2:操作
	Time        time.Time //时间
	InformId    string    //消息uuid
	Informdata  string    //关联数据结构体base64编码后数据
	InformType  int       //消息类型 0:只读消息, 1:agree类型, 2: confirm类型
	OperateType int       //操作类型 0:addOrgType 1:delOrgType
	TypeUniqId  int       //每个消息类型下都维护一个自增Id，用于比新，分为加入邀请和签名两种
	SignerId    string
}