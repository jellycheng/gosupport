package jsonrpc2

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type RpcClient struct {
	url         string
	httpClient  *http.Client
	headers     map[string]string
	timeout     int64

	id string
	//rpc响应原始内容
	responseRawContent string
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
	url := client.url
	id := client.id
	if id=="" {
		id = gosupport.FromatUUIDString(uuid.GenerateUUID(time.Now()))
	}
	rpcReqObj := RPCRequest{
							Jsonrpc:JsonrpcVersion,
							Method:method,
							Id:id,
							Params:params,
						}
	payload := strings.NewReader(rpcReqObj.ToJson())
	fmt.Println(url, payload,rpcReqObj.ToJson())
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	client.responseRawContent = string(body)
	return client
}

//获取结果
func (client *RpcClient)GetResult() (string, error) {
	ret := client.responseRawContent

	return ret, nil
}


//发起批量调用
