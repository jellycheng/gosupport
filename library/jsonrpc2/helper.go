package jsonrpc2

type Method struct {
	methodStr string
}

func (m *Method)GetMethod() string {
	return m.methodStr
}

//拼接请求Method值: JoinMethodOneToStr("User\\Account", "getUserInfo")
func (m *Method)JoinMethodOneToStr(prefix string, suffix string) *Method {
	if prefix != "" {
		m.methodStr = prefix + "." + suffix;
	} else {
		m.methodStr = suffix;
	}
	return m
}

//解析Method
func (m *Method)ParseMethodOne() map[string]string {
	ret := make(map[string]string)
	ret["prefix"] = "";
	ret["suffix"] = "";
	//methodVal = m.methodStr
	//分隔
	
	return ret
}


