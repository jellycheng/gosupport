package gosupport

//通用常量定义

// 版本常量
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
	SignTypeRSA        = "RSA"
	SignTypeRSA2       = "RSA2"

	PrivateFileMode = 0600
)

// 协议相关常量
type Scheme string

const (
	HTTP  Scheme = "http"
	HTTPS Scheme = "https"
	FTP   Scheme = "ftp"
)

// 字符及字符串相关常量
const (
	CharsetStr1 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetStr2 = "0123456789"
	// 去掉了i、j、l、o、u、v、z、I、J、L、O、U、V、Z、0、1、2字符
	CharsetStr3 = "abcdefghkmnpqrstwxyABCDEFGHKMNPQRSTWXY3456789"
)

// 日期时间相关常量
const (
	TIME_FORMAT       = "2006-01-02 15:04:05"
	TimeFormat        = "2006-01-02 15:04:05"
	DateFormat        = "2006-01-02"
	TFormat           = "15:04:05"
	DTNoSecondsFormat = "2006-01-02 15:04"
	Shanghai          = "Asia/Shanghai"  // 上海
	Chongqing         = "Asia/Chongqing" // 重庆
)

const (
	PHP_EOL = "\n"
	GO_EOL  = "\n"
)

// 前景色即文字颜色
const (
	FgBlack   uint8 = iota + 30 // 黑色
	FgRed                       // 红色
	FgGreen                     //绿色
	FgYellow                    //黄色
	FgBlue                      //蓝色
	FgFuchsia                   // 紫红色
	FgCyan                      // 青蓝色
	FgWhite                     // 白色
)

// 背景色
const (
	BgBlack uint8 = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgFuchsia
	BgCyan
	BgWhite
)
