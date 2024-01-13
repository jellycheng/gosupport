package gosupport

import (
	"net/http"
	"regexp"
	"strings"
)

func ParseRoutePath(reg, p string) (bool, map[string]string, []string) {
	isMatch := false
	placeholder := map[string]string{}
	numStr := []string{}
	tags := GetRegTags(reg)
	tagsNum := len(tags)

	regStr := BuildPattern(reg)
	if regObj, err := regexp.Compile(regStr); err == nil {
		res := regObj.FindStringSubmatch(p)
		if len(res) > 0 {
			isMatch = true
			for kIndex, vSlice := range res {
				if kIndex >= 1 {
					numStr = append(numStr, vSlice)
					if tagsNum > 0 && kIndex-1 <= tagsNum {
						tmpName := tags[kIndex-1]
						placeholder[tmpName] = vSlice
					}
				}

			}
		}

	}

	return isMatch, placeholder, numStr
}

// 把含有占位符的串生成正则串，^/<wx>/<type>/getinfo => ^/([^/]+)/([^/]+)/getinfo
func BuildPattern(reg string) string {
	ret := ""
	if regObj, err := regexp.Compile(`<.+?>`); err == nil {
		ret = regObj.ReplaceAllString(reg, `([^/]+)`)
	}
	return ret
}

// 提取占位符: fmt.Println(GetRegTags("^/callback/<wechat>/<type>/xxx"))
func GetRegTags(reg string) []string {
	tags := []string{}
	// `[/]*<(.+?)>[/]*`  "[^/]+"  "[^.]+"
	if regObj, err := regexp.Compile(`<(.+?)>`); err == nil {
		res := regObj.FindAllStringSubmatch(reg, -1)
		if len(res) > 0 {
			for _, vSlice := range res {
				if len(vSlice) > 1 {
					tags = append(tags, vSlice[1])
				}
			}
		}
	}
	return tags
}

type MyHandlerFunc func(http.ResponseWriter, *http.Request, map[string]string, []string)
type MyRoute struct {
	// 路由正则=》处理函数
	getRouter  map[string]MyHandlerFunc
	postRouter map[string]MyHandlerFunc
}

func (m *MyRoute) AddRoute(method string, pattern string, handler MyHandlerFunc) {
	method = strings.ToUpper(method)
	if method == "GET" {
		m.getRouter[pattern] = handler
	} else if method == "POST" {
		m.postRouter[pattern] = handler
	}
}

func (m *MyRoute) GET(pattern string, handler MyHandlerFunc) {
	m.AddRoute("GET", pattern, handler)
}

func (m *MyRoute) POST(pattern string, handler MyHandlerFunc) {
	m.AddRoute("POST", pattern, handler)
}

func (m *MyRoute) Any(pattern string, handler MyHandlerFunc) {
	m.AddRoute("GET", pattern, handler)
	m.AddRoute("POST", pattern, handler)
}

func (m *MyRoute) MatchRoute(method, p string) (bool, MyHandlerFunc, map[string]string, []string) {
	isMatch := false
	placeholder := map[string]string{}
	numStr := []string{}
	method = strings.ToUpper(method)
	if method == "GET" {
		for reg, handler := range m.getRouter {
			isMatch, placeholder, numStr = ParseRoutePath(reg, p)
			if isMatch {
				return isMatch, handler, placeholder, numStr
			}
		}
	} else if method == "POST" {
		for reg, handler := range m.postRouter {
			isMatch, placeholder, numStr = ParseRoutePath(reg, p)
			if isMatch {
				return isMatch, handler, placeholder, numStr
			}
		}
	}
	return isMatch, nil, placeholder, numStr

}

func (m *MyRoute) GetRouter() map[string]MyHandlerFunc {
	return m.getRouter
}

func (m *MyRoute) PostRouter() map[string]MyHandlerFunc {
	return m.postRouter
}

func NewMyRoute() *MyRoute {
	obj := new(MyRoute)
	obj.getRouter = make(map[string]MyHandlerFunc)
	obj.postRouter = make(map[string]MyHandlerFunc)
	return obj
}
