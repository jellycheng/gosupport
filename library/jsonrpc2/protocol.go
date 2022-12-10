package jsonrpc2

import (
	"encoding/json"
)

//https://www.jsonrpc.org/specification

const (
	JsonrpcVersion     = "2.0"  //版本
	CodeSuccess        = 0      //成功
	CodeParseError     = -32700 //分析错误
	CodeInvalidRequest = -32600 //无效请求
	CodeMethodNotFound = -32601 //找不到方法
	CodeInvalidParams  = -32602 //无效的方法参数
	CodeInternalError  = -32603 //Internal error内部错误	JSON-RPC内部错误

)

//请求对象
type RPCRequest struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
	Id      interface{} `json:"id"` //字符串、int、null
}

func (req RPCRequest) ToJson() string {
	return ToJson(req)
}

func NewRPCRequest() *RPCRequest {
	return &RPCRequest{Jsonrpc: JsonrpcVersion}
}

//响应对象
type RPCResponse struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	Id      interface{} `json:"id"`
}

func (rep RPCResponse) ToJson() string {
	return ToJson(rep)
}

type JsonrpcResponse struct {
	Jsonrpc string           `json:"jsonrpc"`
	Result  *json.RawMessage `json:"result,omitempty"`
	Error   *json.RawMessage `json:"error,omitempty"`
	Id      interface{}      `json:"id"`
}

//错误对象
type RPCError struct {
	Code    int         `json:"code"` //整数
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (e RPCError) ToJson() string {
	return ToJson(e)
}

func (e RPCError) Error() string {
	return e.Message
}
