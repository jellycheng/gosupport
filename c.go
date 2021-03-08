package gosupport

//通用常量定义

//版本常量
const (
	_ byte = iota
	V1
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	V9
	V10
)

const (
	// 空字符串
	Empty = ""
	
	// 签名方式
	SignTypeMD5        = "MD5"
	SignTypeHmacSHA256 = "HMAC-SHA256"

	PrivateFileMode = 0600
)

type Scheme string
const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
	FTP   Scheme = "ftp"
)

const (
	CharsetStr1 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetStr2 = "0123456789"
	// 去掉了i、j、l、o、u、v、z、I、J、L、O、U、V、Z、0、1、2字符
	CharsetStr3 = "abcdefghkmnpqrstwxyABCDEFGHKMNPQRSTWXY3456789"
)

const (
	TIME_FORMAT = "2006-01-02 15:04:05"
	DateFormat  = "2006-01-02"
	TFormat     = "15:04:05"
)
