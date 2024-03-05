package rpc

type InputData struct {
	addr        string `json:"addr"`
	from        string `json:"from"`
	to          string `json:"to"`
	contract    string `json:"contract"`
	chainId     string `json:"chainId"`
	page        uint64 `json:"page"`
	pageSize    uint64 `json:"pageSize"`
	startNumber int64  `json:"startNumber"`
	endNumber   int64  `json:"endNumber"`
}

type Response struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}
