package gosupport

// 常用正则常量调用示例：b := gosupport.RegexpVerify("123", gosupport.NumericRegexStr)

const (
	// 是否包含html标签
	HTMLRegexStr = `<[/]?([a-zA-Z]+).*?>`
	// 纯字母
	LetterRegexStr = "^[a-zA-Z]+$"
	// 纯数字
	NumberRegexStr = "^[0-9]+$"
	// 字母 + 数字
	LetterNumericRegexStr = "^[a-zA-Z0-9]+$"
	// 正负浮点数、正负整数
	NumericRegexStr = "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
	// 以 #;开头的字符就算是注释行
	CommentLineRegexStr = "^\\s*[#;]+"

	WwVerify = "^WW_verify_([0-9a-zA-Z]{16}).txt$"
	MpVerify = "^MP_verify_([0-9a-zA-Z]{16}).txt$"

)
