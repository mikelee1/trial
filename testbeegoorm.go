package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/op/go-logging"
	"time"
)

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

var Oconnect orm.Ormer
var dbtype = "mysql"
var dbname1 = "wasabi"
var dbuser = "yunphant"
var dbpasswd = "123456"
var dbip = "192.168.9.83"
var dbport = "38255"
var dbcharset = "utf8"

var logger1 *logging.Logger

func init() {

	orm.RegisterModel(new(Auth))
	orm.RegisterModel(new(ChaincodeInfo))
	orm.RegisterModel(new(OrgChannel))
	orm.RegisterModel(new(Inform))
	orm.RegisterModel(new(Org))
	connectstr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", dbuser, dbpasswd, dbip, dbport, dbname1, dbcharset)
	// 数据库密码明文去除
	err := orm.RegisterDriver(dbtype, orm.DRMySQL)
	if err != nil {
		panic(err)
	}

	err = orm.RegisterDataBase("default", dbtype, connectstr)
	if err != nil {
		panic(err)
	}

	Oconnect = orm.NewOrm()
	orm.RunSyncdb("default", false, false)
	logger1 = logging.MustGetLogger("main")
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


//修改orgchannel表
func queryOrg() {
	tmp := &Org{
		UserId: 1,
	}
	tmp1 := &Org{
		Id:22,
		UserId: 1,
	}
	if _,err := Oconnect.Insert(tmp1); err != nil {
		fmt.Println(err)
		return
	}

	if err := Oconnect.QueryTable("org").One(tmp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tmp.Id)
}

func main() {

	queryOrg()
	//addOrgChannel()
	//tmpdata := []interface{}{}
	//Oconnect.Begin()
	//ccAmount, err := Oconnect.Raw("select a.name from chaincode_info a left join auth b on b.channel = a.channel" +
	//	" and b.chaincode = a.name" +
	//	" where a.channel = ?","channel1").QueryRows(&tmpdata)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//logger1.Info("ccAmount: ",ccAmount)

	//defer func() {
	//	a,err := Oconnect.Insert(&OrgChannel{
	//		Channel: "channel10",
	//		Createorg:"22",
	//		Currentorg: "3",
	//		Createtime:time.Now().Format("2006-01-02 15:04:05")})
	//	if err != nil{
	//		logger1.Info("err:",err)
	//		return
	//	}
	//	Oconnect.Commit()
	//	logger1.Info("a:",a)
	//}()
}

//修改orgchannel表
func addOrgChannel() {
	chname := "channel2"
	tmp := &OrgChannel{}
	if err := Oconnect.QueryTable("org_channel").Filter("channel", chname).One(tmp); err != nil {
		fmt.Println("insert")
		_, err := Oconnect.Insert(&OrgChannel{
			Channel:    chname,
			Createorg:  "baas1",
			Currentorg: "baas3",
		})

		fmt.Println(err)
		return
	}
	fmt.Println("update")
	fmt.Println(tmp.Id)

	a, err := Oconnect.Update(&OrgChannel{
		Id:         tmp.Id,
		Channel:    chname,
		Createorg:  "baas11",
		Currentorg: "baas33",
	})
	fmt.Println(a, err)
}

func addInform() {
	tmp := &Inform{InformDetail: InformDetail{Title: "快乐圣诞节开发", Time: time.Now(), Informdata: "S"}}
	if _, err := Oconnect.Insert(tmp); err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println(tmp.Id)
}
