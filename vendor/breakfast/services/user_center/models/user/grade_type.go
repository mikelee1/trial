package user

// 类型
type GradeType int

const (
	GradeType__Yi GradeType = iota + 1
	GradeType__Er
)

var grades = []string{
	"yi",
	"er",
}

func (t GradeType) String() string {
	return grades[t-1]
}
