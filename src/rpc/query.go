package rpc

import (
	"encoding/json"
	"net/http"
	"strconv"

	"com.tuntun.rangers/service/chaindata/src/middleware/mysql"
)

func count(w http.ResponseWriter, r *http.Request) {
	input := getDataFromUrl(r)
	addr := input.addr
	chainId := input.chainId
	if 0 == len(addr) || 0 == len(chainId) {
		rpcLogger.Errorf("fail to get input, %s", r.RemoteAddr)
		w.Write(failResult("fail to get input"))
		return
	}

	w.Write(successResult(mysql.Count(addr, chainId)))
}

func query(w http.ResponseWriter, r *http.Request) {
	input := getDataFromUrl(r)
	addr := input.addr
	to := input.to
	chainId := input.chainId
	pageSize := input.pageSize

	if (0 == len(addr) && 0 == len(to)) || 0 == len(chainId) || 0 == pageSize {
		rpcLogger.Errorf("fail to get input, %s", r.RemoteAddr)
		w.Write(failResult("fail to get input"))
		return
	}

	w.Write(successResult(mysql.Query(addr, to, chainId, input.page, pageSize)))
}

func queryAdv(w http.ResponseWriter, r *http.Request) {
	input := getDataFromUrl(r)
	req := &mysql.QueryReq{
		From:        input.from,
		To:          input.to,
		Contract:    input.contract,
		StartNumber: input.startNumber,
		EndNumber:   input.endNumber,
	}

	if !req.Validate() {
		rpcLogger.Error("fail to get input:", req)
		w.Write(failResult("fail to get queryAdv input"))
		return
	}

	w.Write(successResult(mysql.QueryAdv(req)))
}

func getDataFromUrl(r *http.Request) *InputData {
	values := r.URL.Query()
	rpcLogger.Debugf("receive: %s, remote: %s", r.URL, r.RemoteAddr)

	var data InputData
	object, ok := values["addr"]
	if ok {
		data.addr = object[0]
	}
	///
	object, ok = values["from"]
	if ok {
		data.from = object[0]
	}
	object, ok = values["to"]
	if ok {
		data.to = object[0]
	}
	object, ok = values["contract"]
	if ok {
		data.contract = object[0]
	}
	///
	object, ok = values["chainId"]
	if ok {
		data.chainId = object[0]
	}

	object, ok = values["page"]
	if ok {
		result, err := strconv.ParseUint(object[0], 10, 32)
		if nil != err {
			rpcLogger.Errorf("fail to parse page, %s", object[0])
			return nil
		}
		data.page = result
	}

	object, ok = values["pageSize"]
	if ok {
		result, err := strconv.ParseUint(object[0], 10, 32)
		if nil != err {
			rpcLogger.Errorf("fail to parse pageSize, %s", object[0])
			return nil
		}
		data.pageSize = result
	}
	///
	object, ok = values["startNumber"]
	if ok {
		result, err := strconv.ParseInt(object[0], 10, 32)
		if nil != err {
			rpcLogger.Errorf("fail to parse page, %s", object[0])
			return nil
		}
		data.startNumber = result
	}
	object, ok = values["endNumber"]
	if ok {
		result, err := strconv.ParseInt(object[0], 10, 32)
		if nil != err {
			rpcLogger.Errorf("fail to parse page, %s", object[0])
			return nil
		}
		data.endNumber = result
	}
	return &data
}

func failResult(s string) []byte {
	result := Response{Status: -1, Result: s}
	data, _ := json.Marshal(result)

	return data
}

func successResult(s interface{}) []byte {
	result := Response{Status: 0, Result: s}
	data, _ := json.Marshal(result)

	return data
}
