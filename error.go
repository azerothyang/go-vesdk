package go_vesdk

// VeRspError 当维易调用异常的时候返回error
type VeRspError struct {
	Err  string      `json:"error"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (vErr VeRspError) Error() string {
	return toJson(vErr)
}
