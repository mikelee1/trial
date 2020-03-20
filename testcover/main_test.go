package testcover_test

import (
	"testing"
	"myproj/try/testcover"
)

//go test -coverprofile=a.out
//
//go tool cover -html=a.out

func Test_main(t *testing.T) {
	testcover.Length()
	testcover.Area(true)
}
