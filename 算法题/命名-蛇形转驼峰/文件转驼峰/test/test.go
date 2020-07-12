package utils

import (
	"yindeng/constant"
)

//-----------角色-----------
type ReqCreateRole struct {
	RoleId    string             `json:"roleId"`
	Name      string             `json:"name,omitempty"`
	Authority constant.Authority `json:"authority,omitempty"`
}

type ReqUpdateRole struct {
	RoleId    string             `json:"roleId"`
	Name      string             `json:"name,omitempty"`
	Authority constant.Authority `json:"authority"`
}

type ReqGetRoleInfo struct {
	RoleId string `json:"roleId"`
}

type RoleListParam struct {
	PageNo   int    `json:"pageNo,omitempty"`   //从第1页开始
	PageSize int    `json:"pageSize,omitempty"` //每页条数
	Bookmark string `json:"bookmark,omitempty"`  //couchdb中的书签号
}

type RespRole struct {
	Roles    []models.BaseRole `json:"roles"`
	Bookmark string            `json:"bookmark"` //该页最后一条的bookmark
	Count    int32             `json:"count"`    //所有条数，用于分页
}

//-----------客户-----------
type ReqCreateMember struct {
	BaseMember models.BaseMember `json:"baseMember"`
}

type ReqUpdateMember struct {
	BaseMember models.BaseMember `json:"baseMember"`
}

type ReqGetMemberInfo struct {
	CustNo string `json:"custNo"`
}

type ReqUpdateMemberRole struct {
	CustNo string `json:"custNo"` // 客户编号,CN
	Role   string `json:"role"`
}

type MemberListParam struct {
	PageNo   int    `json:"pageNo,omitempty"`   //从第1页开始
	PageSize int    `json:"pageSize,omitempty"` //每页条数
	Bookmark string `json:"bookmark,omitempty"`  //couchdb中的书签号

	MemberCode string `json:"memberCode,omitempty"` //客户编号
	MemberName string `json:"memberName,omitempty"` //客户简称
}

type RespMember struct {
	Members  []models.BasicMember `json:"members"`
	Bookmark string               `json:"bookmark"`
	Count    int32                `json:"count"`
}
type RespMemberProducts struct {
	MemberProduct []models.BasicMemberProduct `json:"memberProduct"`
	Bookmark      string                      `json:"bookmark"`
	Count         int32                       `json:"count"`
}

//-----------资产-产品-----------
type ReqCreateAsset struct {
	AssetPkgCode string             `json:"assetPkgCode"`
	BatchNum     string             `json:"batchNum"`
	AssetType    constant.AssetType `json:"assetType"`
	Payload      string             `json:"payload"` // aa []models.CompositeAsset
}

type ReqAssetInfo struct {
	AssetPkgCode string `json:"assetPkgCode"`
	BatchNum     string `json:"batchNum"`
}

type ReqDeleteProduct struct {
	ProdCode string `json:"prodCode"`
}

type ReqConfirmProduct struct {
	ProdCode string `json:"prodCode"`
}
type ReqAuditProduct struct {
	ProdCode string `json:"prodCode"`
	Pass     bool   `json:"pass,omitempty"`    //是否通过
	Message  string `json:"message,omitempty"` //审核意见
}

type ReqCreateProduct struct {
	Product        models.BaseProduct `json:"product"`
	AgreementTrust *models.ChainTrust `json:"agreementTrust,omitempty"`
}

type ReqUpdateProduct struct {
	Product        models.BaseProduct `json:"product"`
	AgreementTrust *models.ChainTrust `json:"agreementTrust,omitempty"`
}

type ReqGetProductInfo struct {
	ProdCode string `json:"prodCode"`
}

type AssetListParam struct {
	PageNo              int    `json:"pageNo,omitempty"`                 //从第1页开始
	PageSize            int    `json:"pageSize,omitempty"`               //每页条数
	Bookmark            string `json:"bookmark,omitempty"`                //couchdb中的书签号
	ProdCode            string `json:"prodCode,omitempty"`               //产品id
	AssetCode           string `json:"assetCode,omitempty,omitempty"`    //资产编号
	BorrowerName        string `json:"borrowerName,omitempty"`           //借款人/债务人简称
	LoanIssOrgShortName string `json:"loanIssOrgShortName,omitempty"` //贷款发放机构简称
}

type ProductListParam struct {
	PageNo   int    `json:"pageNo,omitempty"`   //从第1页开始
	PageSize int    `json:"pageSize,omitempty"` //每页条数
	Bookmark string `json:"bookmark,omitempty"`  //couchdb中的书签号

	OnlySelf bool `json:"onlySelf"` //区分资产管理里的产品列表和信息查看里的产品列表

	ProdName   string              `json:"prodName,omitempty"`   //资产管理里的产品列表：产品名称，支持模糊查询
	ProdStatus constant.ProdStatus `json:"prodStatus,omitempty"` //资产管理里的产品列表：产品状态

	PkgCode              string            `json:"pkgCode,omitempty"`                //资产包id
	StartTime            string            `json:"startTime,omitempty"`              //发布开始时间
	EndTime              string            `json:"endTime,omitempty"`                //发布结束时间
	AssetsOwnerShortName string            `json:"assetsOwnerShortName,omitempty"` //产品发布机构简称
	ProdType             constant.ProdType `json:"prodType,omitempty"`               //产品类型
}

type RespProducts struct {
	Products []models.ChainProduct `json:"products"`
	Bookmark string                `json:"bookmark"`
	Count    int32                 `json:"count"`
}

type RespProductInfo struct {
	Product        models.DetailProduct `json:"product"`
	AgreementTrust *models.ChainTrust   `json:"agreementTrust,omitempty"`
}

type RespAssets struct {
	Assets   []models.BasicAsset `json:"assets"`
	Bookmark string              `json:"bookmark"`
	Count    int32               `json:"count"`
}

type BasicAsset struct {
	AssetCode           string  `json:"assetCode"`           //资产编号
	Borrower            string  `json:"borrower"`            //借款人名称
	LoanIssOrgShortName string  `json:"loanIssOrgShortName"` //贷款发放机构简称
	LoadBalance         float64 `json:"loadBalance"`         //贷款余额
	LoadPeriod          int     `json:"loadPeriod"`          //贷款期限
	LoanPeriodUnit      string  `json:"loanPeriodUnit"`      //贷款期限单位
}

//-----------交易-----------
type ReqCreateTransaction struct {
	BaseTransaction models.BaseTransaction `json:"baseTransaction"`
}

type ReqGetTransactionInfo struct {
	TxCode string `json:"txCode"`
}

type MyBuyProductListParam struct {
	PageNo           int    `json:"pageNo,omitempty"`            //从第1页开始
	PageSize         int    `json:"pageSize,omitempty"`          //每页条数
	Bookmark         string `json:"bookmark,omitempty"`           //couchdb中的书签号
	ProdName         string `json:"prodName,omitempty"`          //产品名称，支持模糊查询
	ProdCode         string `json:"prodCode,omitempty"`          //产品id
	ProdCreator      string `json:"prodCreator,omitempty"`       //出让方
	FixtureDateStart string `json:"fixtureDateStart,omitempty"` //交易开始时间
	FixtureDateEnd   string `json:"fixtureDateEnd,omitempty"`   //交易结束时间
}

type TransactionListParam struct {
	PageNo    int              `json:"pageNo,omitempty"`     //从第1页开始
	PageSize  int              `json:"pageSize,omitempty"`   //每页条数
	Bookmark  string           `json:"bookmark,omitempty"`    //couchdb中的书签号
	ProdCode  string           `json:"prodCode,omitempty"`   //产品id'
	BuyOrSell constant.BuySell `json:"buyOrSell,omitempty"` //买入还是卖出
}

type RespTransactions struct {
	Transactions []models.BasicTransaction `json:"transactions"`
	Bookmark     string                    `json:"bookmark"`
	Count        int32                     `json:"count"`
}

type SelectorStructInterface struct {
	Selector map[string]interface{} `json:"selector"`
	UseIndex []string               `json:"useIndex"`
}
