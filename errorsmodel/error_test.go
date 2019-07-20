package errorsmodel_test

import (
	"fmt"
	"testing"

	"myproj/try/errorsmodel"
)

func TestErrorInt_Code(t *testing.T) {
	a := errorsmodel.FirstError
	fmt.Println(a.String())
}
