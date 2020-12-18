package sign

import (
	"fmt"
	"github.com/jellycheng/gosupport"
	"sort"
	"strings"
)

//签名算法1
type Apisign01 struct {
	secret string
	params map[string]interface{}
	sign_str string
}

func (this *Apisign01) SetSecret(s string) (*Apisign01)  {
	this.secret = s
	return this
}

func (this *Apisign01) SetParams(p map[string]interface{}) (*Apisign01)  {
	/**
	if(this.params == nil) {
		this.params = make(map[string]interface{})
	}
	*/
	this.params = p
	return this
}

func (this *Apisign01) AppendParam(k string, v interface{}) (*Apisign01)  {
	this.params[k] = v
	return this
}

func (this *Apisign01) Md5Sign() (string)  {
	sign := ""
	var keys []string
	for k, v := range this.params {
		if k == "sign" || v == "" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	builder := strings.Builder{}
	if len(keys) != 0 {
		for _, v := range keys {
			//builder.WriteString(v)
			//builder.WriteString("=")
			//builder.WriteString(fmt.Sprint(this.params[v]))
			//builder.WriteString("&")
			builder.WriteString(fmt.Sprintf("%s=%s&", v, fmt.Sprint(this.params[v])))

		}
		builder.WriteString("key=" + this.secret)
	} else {
		builder.WriteString("&key=" + this.secret)
	}
	this.sign_str = builder.String()
	sign = strings.ToUpper(gosupport.Md5V1(this.sign_str))

	return sign
}

func (this Apisign01) GetSignString() string  {
	return this.sign_str
}

func NewApiSign01() (*Apisign01)  {
	apisign := new(Apisign01)
	apisign.params = make(map[string]interface{})
	return apisign
}



