package testcover_test

import (
	"testing"
	"myproj/try/testcover"
)

func Test_main(t *testing.T) {
	testcover.Length()
	testcover.Area(true)
}
