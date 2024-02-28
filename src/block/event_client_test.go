package block

import (
	"com.tuntun.rangers/service/chaindata/src/common"
	"com.tuntun.rangers/service/chaindata/src/middleware/log"
	"com.tuntun.rangers/service/chaindata/src/middleware/mysql"
	"os"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	defer func() {
		log.Close()
		os.RemoveAll("logs")
	}()

	common.InitConf("chain.ini")
	mysql.InitMySql()

	Init()

	time.Sleep(10 * time.Hour)
}
