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
	//空字符串
	Empty = ""
	
	//签名方式
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
)
