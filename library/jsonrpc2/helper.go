package jsonrpc2

import (
	"encoding/json"
	"strings"
)

type Method struct {
	methodStr string
}

func (m *Method)GetMethod() string {
	return m.methodStr
}

//拼接请求Method值: JoinMethodOneToStr("User\\Account", "getUserInfo") 或者 JoinMethodOneToStr(`User\Account`, "getUserInfo")
func (m *Method)JoinMethodOneToStr(prefix string, suffix string) *Method {
	prefix = strings.TrimSpace(prefix)
	suffix = strings.TrimSpace(suffix)
	if prefix != "" && suffix != "" {
		m.methodStr = prefix + "." + suffix;
	} else if prefix != "" {
		m.methodStr = prefix;
	} else if suffix != "" {
		m.methodStr = suffix;
	}
	return m
}

//解析Method
func (m *Method)ParseMethodOne() map[string]string {
	ret := make(map[string]string)
	ret["prefix"] = "";
	ret["suffix"] = "";
	methodVal := m.methodStr
	//分隔
	methodSplit := strings.SplitN(methodVal, ".",  2)
	if len(methodSplit) == 1 {
		ret["prefix"] = methodVal
	} else {
		ret["prefix"] = methodSplit[0]
		ret["suffix"] = methodSplit[1]
	}
	return ret
}


//转成json字符串
func ToJson(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

func JsonToRPCRequestStruct(str string) RPCRequest {
	var reqRpcObj RPCRequest
	err := json.Unmarshal([]byte(str), &reqRpcObj)
	if err!=nil {
		return RPCRequest{}
	}
	return reqRpcObj
}

func JsonToRPCResponseStruct(str string) RPCResponse {
	var repRpcObj RPCResponse
	err := json.Unmarshal([]byte(str), &repRpcObj)
	if err!=nil {
		return RPCResponse{}
	}
	return repRpcObj
}

func JsonToRPCErrorStruct(str string) RPCError {
	var reeRpcObj RPCError
	err := json.Unmarshal([]byte(str), &reeRpcObj)
	if err!=nil {
		return RPCError{}
	}
	return reeRpcObj
}

