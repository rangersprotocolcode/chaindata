package block

import (
	"com.tuntun.rangers/service/chaindata/src/common"
	"com.tuntun.rangers/service/chaindata/src/middleware/log"
	"strings"
	"time"
)

const (
	chains       = "chains"
	chainsHeight = "height"
	contracts    = "contracts"
)

var (
	chainClient map[string]*ethModule
	closed      = false
)

func Init() {
	chainClient := make(map[string]*ethModule, 4)

	chainRpcList := common.GlobalConf.GetStrings(chains)
	if 0 == len(chainRpcList) {
		return
	}

	chainContracts := common.GlobalConf.GetStrings(contracts)
	for chain, rpc := range chainRpcList {
		key := strings.ToLower(chain)
		rpcList := strings.Split(rpc, ",")
		module := ethModule{
			rpcList: rpcList,
		}
		chainClient[key] = &module

		module.start(key, chainContracts[key])
	}

	logLoop()
}

func Close() {
	closed = true
	for _, module := range chainClient {
		module.close()
	}
}

func logLoop() {
	ticker := time.NewTicker(time.Second * 10)
	monitorLogger := log.GetLoggerByIndex(log.MonitorLogConfig, "")

	go func() {
		for _ = range ticker.C {
			if closed {
				return
			}

			for _, module := range chainClient {
				monitorLogger.Info(module.info())
			}
		}
	}()
}
