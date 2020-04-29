package jsonrpc2

import (
	"fmt"
	"testing"
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
	fmt.Println(jsonRpcStr01)

	//json字符串转结构体对象
	var reqRpcObj RPCRequest = JsonToRPCRequestStruct(jsonRpcStr01)
	fmt.Println(reqRpcObj.Method)


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
	fmt.Println(jsonRpcStr4Batch)

}

