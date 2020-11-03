package jsonrpc2

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	abc := Method{}
	abc.JoinMethodOneToStr(`User\Account`, "getUserInfo")
	rpcObj := RPCRequest{
			Jsonrpc:JsonrpcVersion,
			Method:abc.GetMethod(),
			Id:"1233456",
			Params:"hello world",

	}
	//结构体对象转json字符串
	jsonRpcStr01 := rpcObj.ToJson()
	fmt.Println(jsonRpcStr01) //{"jsonrpc":"2.0","method":"User\\Account.getUserInfo","params":"hello world","id":"1233456"}

	//json字符串转结构体对象
	var reqRpcObj RPCRequest = JsonToRPCRequestStruct(jsonRpcStr01)
	fmt.Println(reqRpcObj.Method) //User\Account.getUserInfo


	//拼装批量协议
	abc2 := Method{}
	abc2.JoinMethodOneToStr(`User\Address`, "getAddressList")
	rpcObj2 := RPCRequest{
		Jsonrpc:JsonrpcVersion,
		Method:abc2.GetMethod(),
		Id:"12334567",
		Params:[]int{10, 20, 50},
	}
	abc3 := Method{}
	abc3.JoinMethodOneToStr(`User\Account`, "getUserInfo2")
	rpcObj3 := RPCRequest{
		Jsonrpc:JsonrpcVersion,
		Method:abc3.GetMethod(),
		Id:"12334568",
		Params:[]string{"uid123", "uid456", "uid789"},
	}
	var batchRpc = []RPCRequest{}
	batchRpc = append(batchRpc, rpcObj2)
	batchRpc = append(batchRpc, rpcObj3)
	jsonRpcStr4Batch := ToJson(batchRpc)
	//[{"jsonrpc":"2.0","method":"User\\Address.getAddressList","params":[10,20,50],"id":"12334567"},
	//{"jsonrpc":"2.0","method":"User\\Account.getUserInfo2","params":["uid123","uid456","uid789"],"id":"12334568"}]
	fmt.Println(jsonRpcStr4Batch)

}

func TestRpc02(t *testing.T)  {
	 url := "http://cart.devci01.s.dev.xxx.com/rpc.php"
	 rpcObj := NewRpcClient(url, int64(15 * time.Second))
	 param := map[string]interface{}{
						 "owner":"3427234",
						 "selected": 1,
						 "cart_type": 0,
						 "owner_type": 1,
					 }
	headers := map[string]string{
		"X-FROM-SERVICE": "community-service",
		"Content-Type":  "application/json",
		"BRANCHNAME":  "master",
		"X-ENTERPRISE-ID":  "10000",
	}
	 content, _ := rpcObj.SetId("abc-123-xyz-66").AddHeaders(headers).Call(`Cart\Search.listGoods`, param).GetResult()
	 fmt.Println(content)
	 var jsonMap map[string]interface{}
	 json.Unmarshal([]byte(content), &jsonMap)
	 fmt.Println(jsonMap)
}

