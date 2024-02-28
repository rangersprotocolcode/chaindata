package cli

import (
	"com.tuntun.rangers/service/chaindata/src/block"
	"com.tuntun.rangers/service/chaindata/src/common"
	"com.tuntun.rangers/service/chaindata/src/middleware/mysql"
	"com.tuntun.rangers/service/chaindata/src/rpc"
	"testing"
	"time"
)

func TestGX_Run(t *testing.T) {
	common.InitConf("chain.ini")
	mysql.InitMySql()
	block.Init()
	rpc.Init("8888")

	time.Sleep(10 * time.Hour)
}
