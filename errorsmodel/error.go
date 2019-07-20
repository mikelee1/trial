package errorsmodel

type errorInt int

const (
	DefaultError errorInt = iota
	FirstError
)

var errorString = []string{
	"DefaultError",
	"FirstError",
}

func (e errorInt) Code() int {
	return int(e)
}

func (e errorInt) String() string {
	return errorString[e]
}
