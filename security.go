package gosupport

import "strings"

var quotesEscaper = strings.NewReplacer(
	`'`, "&#39;",
	`"`, "&#34;",
)

var xssEscaper = strings.NewReplacer(
	`<`, "&lt;",
	`>`, "&gt;",
)

func QuotesToEntity(s string) string {
	return quotesEscaper.Replace(s)
}

// XssToEntity 转义< >，更多字符转义见 html.EscapeString(s) 方法
func XssToEntity(s string) string {
	return xssEscaper.Replace(s)
}
