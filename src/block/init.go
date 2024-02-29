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
	chainClients map[string]*ethModule
	closed       = false
)

func Init() {
	chainClients = make(map[string]*ethModule, 4)

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
		chainClients[key] = &module

		module.start(key, chainContracts[key])
	}

	logLoop()
}

func Close() {
	closed = true
	for _, module := range chainClients {
		module.close()
	}
}

func logLoop() {
	monitorLogger := log.GetLoggerByIndex(log.MonitorLogConfig, "")

	go func() {
		for range time.Tick(time.Second * 10) {
			if closed {
				return
			}

			for _, module := range chainClients {
				monitorLogger.Info(module.info())
			}
		}
	}()
}
