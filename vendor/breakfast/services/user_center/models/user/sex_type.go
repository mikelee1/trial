package user

// 类型
type SexType int

const (
	SexType__Nan SexType = iota + 1
	SexType__Nv
	SexType__Baomi
)

var sexs = []string{
	"1",
	"2",
	"3",
}

func (t SexType) String() string {
	return sexs[t-1]
}
