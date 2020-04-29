package jsonrpc2

import "net/http"

type RpcClient struct {
	url         string
	httpClient  *http.Client
	headers     map[string]string
	timeout     int64
}

func NewRpcClient(url string, timeout int64) *RpcClient {
	rpcCli := &RpcClient{
				url:url,
				httpClient:&http.Client{},
				headers:make(map[string]string),
				timeout:timeout,
	}
	rpcCli.headers["Content-Type"] = "application/json"

	return rpcCli
}

func (client *RpcClient)SetTimeout(timeout int64) *RpcClient {
	client.timeout = timeout
	return client
}

//发起单个调用
func (client *RpcClient)Call(method string, params ...interface{}) *RpcClient  {

	return client
}

//获取结果
func (client *RpcClient)GetResult() (string, error) {
	ret := ""

	return ret, nil
}


//发起批量调用
