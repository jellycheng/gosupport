package sensitiveword

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

//敏感词过滤 - 简易写法

type FilterSensitiveWord struct {
	sensitiveWordDict string //敏感词，多个用|分隔
	replaceSymbol     string //替换符，默认*

}

//获取敏感词
func (f *FilterSensitiveWord) SensitiveWordDict() string {
	return f.sensitiveWordDict
}
//设置敏感词，会覆盖以前设置的
func (f *FilterSensitiveWord) SetSensitiveWordDict(sensitiveWordDict string) *FilterSensitiveWord {
	f.sensitiveWordDict = strings.TrimSpace(sensitiveWordDict)
	return f
}

//追加敏感词
func (f *FilterSensitiveWord) AppendSensitiveWordDict(sensitiveWordDict string) *FilterSensitiveWord {
	if f.sensitiveWordDict == "" {
		f.sensitiveWordDict = strings.TrimSpace(sensitiveWordDict)
	} else {
		f.sensitiveWordDict = fmt.Sprintf("%s|%s", f.sensitiveWordDict,strings.TrimSpace(sensitiveWordDict))
	}
	return f
}

//获取替换符
func (f *FilterSensitiveWord) ReplaceSymbol() string {
	return f.replaceSymbol
}

//设置替换符
func (f *FilterSensitiveWord) SetReplaceSymbol(replaceSymbol string) *FilterSensitiveWord  {
	f.replaceSymbol = replaceSymbol
	return f
}
//过滤敏感词
func (f *FilterSensitiveWord) Filter(src string) (string, error)  {
	pattern01 := f.sensitiveWordDict //敏感词过滤, 内容如 `涉黄|X大片`
	re, err := regexp.Compile(pattern01)
	if err!=nil {
		return src, errors.New("正则表达式错误")
	} else {
		newStr := re.ReplaceAllStringFunc(src, func(s string) string {
			i := utf8.RuneCountInString(s)
			return strings.Repeat(f.replaceSymbol, i)
		})
		return newStr, nil
	}
}

func NewFilterSensitiveWord(sensitiveWordDict string, replaceSymbol string) *FilterSensitiveWord {
	if replaceSymbol == "" { //默认*，但可以调用设置方法重新更改，甚至替换为空
		replaceSymbol = "*"
	}
	return &FilterSensitiveWord{sensitiveWordDict: sensitiveWordDict, replaceSymbol: replaceSymbol}
}


