package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"github.com/op/go-logging"
	"time"
	"myproj/try/testbeegoorm/models"
)


var logger1 *logging.Logger



