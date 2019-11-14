package main_test

import (
	"testing"
	"time"
	"github.com/op/go-logging"
	"math"
	"fmt"
)

var logger = logging.MustGetLogger("test")

func Test_TranDuration(t *testing.T) {
	duration := - time.Minute
	h := math.Floor(duration.Hours())

	newduration := duration - time.Duration(h)*time.Hour
	m := math.Floor(newduration.Minutes())
	if h == 0 {
		fmt.Sprintf("%d分", int(m))
	} else {
		fmt.Sprintf("%d时%d分", int(h), int(m))
	}
	logger.Info(h, m)
}
