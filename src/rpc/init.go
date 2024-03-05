package rpc

import (
	"net/http"
	"strings"

	"com.tuntun.rangers/service/chaindata/src/middleware/log"
)

var rpcLogger log.Logger

func Init(port string) {
	rpcLogger = log.GetLoggerByIndex(log.RpcLogConfig, "")

	mux := http.NewServeMux()
	mux.HandleFunc("/count", count)
	mux.HandleFunc("/query", query)
	mux.HandleFunc("/queryAdv", queryAdv)

	go func() {
		err := http.ListenAndServe(":"+port, cros{next: mux})
		if nil != err {
			rpcLogger.Errorf("fail to start rpc, error: %s", err)
		}
	}()

	rpcLogger.Warnf("start rpc, port: %s", port)
}

type cros struct {
	next *http.ServeMux
}

func (service cros) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "token, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	if 0 == strings.Compare("OPTIONS", strings.ToUpper(r.Method)) {
		w.WriteHeader(200)
		return
	}

	service.next.ServeHTTP(w, r)
}
