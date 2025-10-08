package jsonrpc2

import (
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/uuid"
	"io"
	"net/http"
	"strings"
	"time"
)

type RpcClient struct {
	url        string
	httpClient *http.Client
	headers    map[string]string
	timeout    int64

	//rpc的id
	id string
	//rpc响应原始内容
	responseRawContent string
}

func NewRpcClient(url string, timeout int64) *RpcClient {
	rpcCli := &RpcClient{
		url:        url,
		httpClient: &http.Client{},
		headers:    make(map[string]string),
		timeout:    timeout,
	}
	rpcCli.headers["Content-Type"] = "application/json"

	return rpcCli
}

func (m *RpcClient) SetTimeout(timeout int64) *RpcClient {
	m.timeout = timeout
	return m
}

func (m *RpcClient) SetId(traceid string) *RpcClient {
	m.id = traceid
	return m
}

func (m *RpcClient) AddHeader(header, val string) *RpcClient {
	m.headers[header] = val
	return m
}

func (m *RpcClient) AddHeaders(header map[string]string) *RpcClient {
	for k, v := range header {
		m.headers[k] = v
	}
	return m
}

// 发起单个调用
func (m *RpcClient) Call(method string, params ...interface{}) *RpcClient {
	url := m.url
	id := m.id
	if id == "" {
		id = gosupport.FromatUUIDString(uuid.GenerateUUID(time.Now()))
	}
	rpcReqObj := RPCRequest{
		Jsonrpc: JsonrpcVersion,
		Method:  method,
		Id:      id,
		Params:  params,
	}
	payload := strings.NewReader(rpcReqObj.ToJson())
	//fmt.Println(url, payload,rpcReqObj.ToJson())
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Set("content-type", "application/json")
	//处理header,追加头，存在不修改
	for k, v := range m.headers {
		req.Header.Add(k, v)
	}

	cliObj := &http.Client{
		Timeout: time.Duration(m.timeout),
	}
	res, err := cliObj.Do(req)
	if err != nil {
		return m
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	m.responseRawContent = string(body)
	return m
}

// 获取结果
func (m *RpcClient) GetResult() (string, error) {
	ret := m.responseRawContent
	return ret, nil
}

//发起批量调用
