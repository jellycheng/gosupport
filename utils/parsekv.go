package utils

import (
	"bytes"
	"regexp"
	"strings"
)

//分析注释、env、ini配置行的key-value值
type ParseKV struct {
	commentFlag   string //支持的注释标识符，如 #、;、
	content       string //原内容
	resultKey     string //分析结果key名
	resultVal     string //分析结果值
	resultComment string //注释内容
}

func (p *ParseKV) SetCommnetFlag(flag string) *ParseKV {
	tmpFlag := strings.TrimSpace(flag)
	p.commentFlag = tmpFlag
	return p
}
func (p *ParseKV) GetCommnetFlag() string {
	return p.commentFlag
}

func (p *ParseKV) SetContent(con string) *ParseKV {
	p.content = con
	return p
}

func (p *ParseKV) GetContent() string {
	return p.content
}

func (p *ParseKV) GetResultKey() string {
	return p.resultKey
}

func (p *ParseKV) GetResultVal() string {
	return p.resultVal
}

func (p *ParseKV) GetResultComment() string {
	return p.resultComment
}

//把内容和注释分开,返回内容、注释、是否存在注释
func (p *ParseKV) SplitContent4Flag(in []byte) ([]byte, []byte, bool) {
	i := bytes.IndexAny(in, p.commentFlag) //"#;"
	if i == -1 {
		return in, nil, false
	}
	return in[:i], in[i+1:], true
}

func (p *ParseKV) Parse() *ParseKV {
	envReg := regexp.MustCompile(`(\w+)=('(.*)'|"(.*)"|(.*))`)
	pcontent := strings.TrimSpace(p.content)
	contentByte, commentByte, isComment := p.SplitContent4Flag([]byte(pcontent))
	if isComment == true {
		p.resultComment = string(commentByte)
	}
	content := strings.TrimSpace(string(contentByte))
	if content == "" {
		return p
	}
	res := envReg.FindAllStringSubmatch(content, -1)
	for _, param := range res {
		p.resultKey = param[1]
		val := param[3]
		if val == "" && param[4] != "" {
			val = param[4]
		}
		if val == "" && param[5] != "" {
			val = param[5]
		}
		p.resultVal = val
		break
	}
	return p
}

func NewParseKV() *ParseKV {
	obj := &ParseKV{
		commentFlag: "#",
	}
	return obj
}
