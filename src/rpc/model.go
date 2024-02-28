package rpc

type InputData struct {
	addr     string `json:"addr"`
	chainId  string `json:"chainId"`
	page     uint64 `json:"page"`
	pageSize uint64 `json:"pageSize"`
}

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}
