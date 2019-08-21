package user

// 类型
type NowStatusType int

const (
	NowStatusType__Bukaolv NowStatusType = iota + 1
	NowStatusType__Xianghuan
	NowStatusType__Yilizhi
	NowStatusType__Baomi
)

var nowstatus = []string{
	"zaizhi,bukaolv",
	"zaizhi,xianghuan",
	"yilizhi",
	"baomi",
}

func (t NowStatusType) String() string {
	return nowstatus[t-1]
}
