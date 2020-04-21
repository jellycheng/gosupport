package jsonrpc2

//https://www.jsonrpc.org/specification

const (
	JsonrpcVersion = "2.0" //版本
	CodeSuccess = 0  //成功
	CodeParseError = -32700 //分析错误
	CodeInvalidRequest = -32600 //无效请求
	CodeMethodNotFound = -32601 //找不到方法

)

//请求对象
type RPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Id      interface{}      `json:"id"`  //字符串、int、null
}

//响应对象
type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	Id      interface{}      `json:"id"`
}

//错误对象
type RPCError struct {
	Code    int         `json:"code"`    //整数
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}


