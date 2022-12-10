package sign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// https://cloud.tencent.com/document/api/865/35466
type TencentSign struct {
	SecretId             string // 密钥对中的 SecretId
	SecretKey            string // 密钥对中的 SecretKey
	Algorithm            string // TC3-HMAC-SHA256
	Signature            string // 签名值,签名结果
	HeaderTimestamp      int64  // 设置请求头时间戳，utc时间
	ServiceName          string // service服务名 即产品名，如 cvm
	HTTPRequestMethod    string // HTTP 请求方法（GET、POST ）
	CanonicalURI         string // URI 参数，API 3.0 固定为正斜杠（/）
	CanonicalQueryString string // 发起 HTTP 请求 URL 中的查询字符串，对于 POST 请求，固定为空字符串""，对于 GET 请求，则为 URL 中问号（?）后面的字符串内容，例如：Limit=10&Offset=0
	//CredentialScope string // 凭证范围，格式： 年-月-日/service服务名/tc3_request，示例：2019-02-25/cvm/tc3_request
	SignedHeaders        string // 参与签名的头部信息，示例：content-type;host
	CanonicalHeaders     string // 参与签名的请求头和值,key小写，如： content-type:application/json; charset=utf-8\nhost:cvm.tencentcloudapi.com\n
	Host                 string // tiia.tencentcloudapi.com
	HashedRequestPayload string // 请求正文payload，即 body内容
}

// 当前UTC时间戳
func (m TencentSign) GetTimeNow() int64 {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc).Unix()
}

// 时间戳转指定的格式
func (m TencentSign) Timestamp2Date(timestamp int64, format string, timezone ...*time.Location) string {
	var loc *time.Location
	if len(timezone) == 0 {
		loc = time.UTC
	} else {
		loc = timezone[0]
	}
	return time.Unix(timestamp, 0).In(loc).Format(format)
}

// 获取凭证范围，格式： 年-月-日/service服务名/tc3_request，示例：2019-02-25/cvm/tc3_request
func (m TencentSign) GetCredentialScope() string {
	date := m.Timestamp2Date(m.HeaderTimestamp, "2006-01-02")
	s := fmt.Sprintf("%s/%s/tc3_request", date, m.ServiceName)
	return s
}

func (m TencentSign) PingCanonicalHeaders(contentType, host string) string {
	ret := fmt.Sprintf("content-type:%s\nhost:%s\n", contentType, host)
	return ret
}

func (m TencentSign) GetCanonicalRequest() string {
	canonicalRequest := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		m.HTTPRequestMethod,
		m.CanonicalURI,
		m.CanonicalQueryString,
		m.CanonicalHeaders,
		m.SignedHeaders,
		m.Sha256hex(m.HashedRequestPayload))
	return canonicalRequest
}

func (m *TencentSign) GetSign() string {
	sign := ""
	canonicalRequest := m.GetCanonicalRequest()
	signSrc := fmt.Sprintf("%s\n%d\n%s\n%s",
		m.Algorithm,
		m.HeaderTimestamp,
		m.GetCredentialScope(),
		m.Sha256hex(canonicalRequest))
	date := m.Timestamp2Date(m.HeaderTimestamp, "2006-01-02")
	secretDate := m.HmacSha256(date, "TC3"+m.SecretKey)
	secretService := m.HmacSha256(m.ServiceName, secretDate)
	secretDesKey := m.HmacSha256("tc3_request", secretService)
	sign = hex.EncodeToString([]byte(m.HmacSha256(signSrc, secretDesKey)))
	m.Signature = sign
	return sign
}

func (m TencentSign) GetAuthorization() string {
	s := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		m.Algorithm,
		m.SecretId,
		m.GetCredentialScope(),
		m.SignedHeaders,
		m.GetSign())
	return s
}

func (m TencentSign) HmacSha256(src, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(src))
	return string(h.Sum(nil))
}

func (m TencentSign) Sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func NewTencentSign(secretId, secretKey string) TencentSign {
	ts := TencentSign{
		SecretId:  secretId,
		SecretKey: secretKey,
	}
	ts.Algorithm = "TC3-HMAC-SHA256"
	ts.CanonicalURI = "/"
	ts.HeaderTimestamp = ts.GetTimeNow()
	return ts
}
