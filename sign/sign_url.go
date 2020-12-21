package sign

import (
	"crypto/sha1"
	"fmt"
	"github.com/jellycheng/gosupport"
	"io"
	"sort"
)

type UrlGetSign struct {
	timestamp string //时间戳
	nonce string  //随机字符串
	secret string //公共密钥，所有内部服务使用一样的值
}

func (u *UrlGetSign)SetTimestamp(s_timestamp string) *UrlGetSign  {
	u.timestamp = s_timestamp
	return u
}

func (u *UrlGetSign)GetTimestamp() string {
	return u.timestamp
}

func (u *UrlGetSign)SetNonce(s_nonce string) *UrlGetSign {
	u.nonce = s_nonce
	return u
}

func (u *UrlGetSign)GetNonce() string {
	return u.nonce
}

func (u *UrlGetSign)SetSecret(s string) *UrlGetSign {
	u.secret = s
	return u
}

func (u *UrlGetSign)GetSecret() string {
	return u.secret
}

//自动生成拼的url参数
func (u *UrlGetSign)AutoSpellUrlParam() string  {
	getStrFormat := "s_timestamp=%s&s_nonce=%s&s_sign=%s"
	var t,n,s string
	if u.timestamp!="" {
		t = u.timestamp
	} else {
		t = gosupport.ToStr(gosupport.Time())
		u.timestamp = t
	}
	if u.nonce!="" {
		n = u.nonce
	} else {
		n = gosupport.GetRandString(8)
		u.nonce = n
	}
	s = u.GetSign()
	getStr := fmt.Sprintf(getStrFormat, t, n, s)

	return getStr
}

func (u *UrlGetSign)GetSign() string {
	var params = []string{
		u.timestamp,
		u.nonce,
		u.secret,
	}
	sort.Strings(params)
	h := sha1.New()
	for _, s := range params {
		_, _ = io.WriteString(h, s)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

//sign.NewUrlGetSign().SetSecret("cjsJellySecret123456").AutoSpellUrlParam
func NewUrlGetSign() *UrlGetSign  {
	return &UrlGetSign{}
}

func Checks2sSign(timestamp, nonce, sign, secret string) bool {
	var meSign = NewUrlGetSign().SetSecret(secret).SetNonce(nonce).SetTimestamp(timestamp).GetSign()

	if meSign == sign {
		return true
	} else {
		return false
	}
}


/**
php算法：
function checks2sSign($s_sign, $s_timestamp, $s_nonce, $secret) {
    $tmpArr = array($secret, $s_timestamp, $s_nonce);
    sort($tmpArr, SORT_STRING);
    $tmpStr = implode($tmpArr);
    $tmpStr = sha1($tmpStr);

    if( $tmpStr == $s_sign ){
        return true;
    }else{
        return false;
    }
}
 */

