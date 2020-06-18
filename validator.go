package gosupport

import "regexp"

//验证


//是否邮箱
func IsMail(mail string) bool {
	isMatch,err := regexp.MatchString("^([.a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((.[a-zA-Z0-9_-]{2,10}){1,3})$", mail)
	if err!=nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

//是否手机号，11位数字
func IsMobile(str string) bool  {
	isMatch,err := regexp.MatchString("^1[0-9]{10}$", str)
	if err!=nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}

//是否座机
func IsPhone(str string) bool  {
	return checkRegexp(str, "^([0-9]{3,4}-)?[0-9]{7,8}$")
}

//是否链接
func IsUrl(str string) bool  {
	return checkRegexp(str, `^http[s]?://.*`)
}

func checkRegexp(val string, reg string) bool {
	isMatch,err := regexp.MatchString(reg, val)
	if err!=nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}


//字符串是否为整数数字字符串
func IsNumber(str string) bool {
	isMatch,err := regexp.MatchString("^[1-9][0-9]*$", str)
	if err!=nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}


//字符串是否为浮点数字符串
func IsFloatNumber(str string) bool {
	isMatch,err := regexp.MatchString("^[0-9]+[.]?[0-9]*$", str)
	if err!=nil {
		return false
	}
	if isMatch {
		return true
	} else {
		return false
	}
}
