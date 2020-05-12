package gosupport

import (
	"net/http"
	"strconv"
	"strings"
	"time"
)

//cors配置结构体

type CorsConfig struct {
	//是否开启跨域，true是，false否
	IsOpenCors bool
	//允许的源,默认*，赋值给 Access-Control-Allow-Origin
	AllowOrigins string
	//允许的请求方式，默认*，如GET、POST，赋值给Access-Control-Allow-Methods
	AllowMethods []string
	//允许的请求头，赋值给 Access-Control-Allow-Headers
	AllowHeaders []string
	//值为 true、false，赋值给Access-Control-Allow-Credentials
	AllowCredentials bool
	//赋值给 Access-Control-Max-Age
	MaxAge time.Duration

}
//设置允许的请求方式
func (c *CorsConfig) AddAllowMethods(methods ...string) {
	c.AllowMethods = append(c.AllowMethods, methods...)
}

//设置允许的请求头
func (c *CorsConfig) AddAllowHeaders(headers ...string) {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
}

//生成http类型头,类型：type Header map[string][]string
func (c *CorsConfig) GenerateHttpHeaders() http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	headers.Set("Access-Control-Allow-Origin", c.GetAllowOrigins())
	if len(c.AllowMethods) > 0 {
		allowMethods := convert(normalize(c.AllowMethods), strings.ToUpper)
		value := strings.Join(allowMethods, ",")
		headers.Set("Access-Control-Allow-Methods", value)
	}
	if len(c.AllowHeaders) > 0 {
		allowHeaders := convert(normalize(c.AllowHeaders), http.CanonicalHeaderKey)
		value := strings.Join(allowHeaders, ",")
		headers.Set("Access-Control-Allow-Headers", value)
	}
	if c.MaxAge > time.Duration(0) {
		value := strconv.FormatInt(int64(c.MaxAge/time.Second), 10)
		headers.Set("Access-Control-Max-Age", value)
	}

	return headers
}

//设置允许的源
func (c *CorsConfig) SetAllowOrigins(origin string) {
	c.AllowOrigins = origin
}

//获取允许的源
func (c *CorsConfig) GetAllowOrigins() string {
	var ret string
	if c.AllowOrigins == "" {
		ret = "*"
	} else {
		ret = c.AllowOrigins
	}
	return ret
}

//corsConfig := gosupport.DefaultCorsConfig()
func DefaultCorsConfig() CorsConfig {
	return CorsConfig{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type","token","app-type","app-v","*"},
		AllowCredentials: true,
		MaxAge:           10 * time.Hour,
		IsOpenCors: true,
	}
}

type ConverterFunc func(string) string

func convert(s []string, c ConverterFunc) []string {
	var out []string
	for _, i := range s {
		out = append(out, c(i))
	}
	return out
}

func normalize(values []string) []string {
	if values == nil {
		return nil
	}
	distinctMap := make(map[string]bool, len(values))
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.ToLower(value)
		if _, seen := distinctMap[value]; !seen {
			normalized = append(normalized, value)
			distinctMap[value] = true
		}
	}
	return normalized
}

