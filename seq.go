package gosupport

import (
	"fmt"
	"strconv"
	"time"
)

// Sequence seq字符串格式：%s%s%06s%03s 分别是 前缀+（6位年月日*100000+5位时间差（当前时间戳-今天凌晨时间戳））+ (步长%999999=前导补零的6位串) +（用户ID%128）
type Sequence struct {
	prefix   string // 前缀
	uid      int64
	stepSize int64 //步长
}

func (m *Sequence) SetPrefix(prefixVal string) *Sequence {
	m.prefix = prefixVal
	return m
}

func (m *Sequence) SetUid(uidVal int64) *Sequence {
	m.uid = uidVal
	return m
}

func (m *Sequence) GetUidMod() int64 {
	ret := m.uid % 128
	return ret
}

func (m *Sequence) SetStepSize(stepSizeVal int64) *Sequence {
	if stepSizeVal <= 0 {
		stepSizeVal = 1
	}
	m.stepSize = stepSizeVal
	return m
}

func (m Sequence) GetSeq(uidModVal ...int64) string {
	seq := strconv.FormatInt(m.stepSize%999999, 10)
	var uidMod int64 = 0
	if len(uidModVal) > 0 {
		uidMod = uidModVal[0]
	} else {
		uidMod = m.GetUidMod()
	}
	timeStr := m.getTimestampSeq()
	ret := fmt.Sprintf("%s%s%06s%03s", m.prefix, timeStr, seq, strconv.FormatInt(uidMod, 10))

	return ret
}

func (m Sequence) getTimestampSeq() string {
	ret := ""
	loc := GetShanghaiTimezone()
	today := time.Now().In(loc).Format("2006-01-02")
	start, _ := time.ParseInLocation("2006-01-02 15:04:05", today+" 00:00:00", loc)
	startTime := start.Unix()
	endTime := time.Now().In(loc).Unix()
	second := endTime - startTime
	d := StrTo(time.Now().In(loc).Format("060102")).MustInt64() * 100000
	ret = ToStr(d + second)
	return ret
}

func NewSequence() *Sequence {
	return &Sequence{}
}

type SeqString string

func (m SeqString) GetUidMod4Seq() int64 {
	var ret int64 = 0
	if v, err := strconv.ParseInt(m.GetUidMod4Seq2Str(), 10, 64); err == nil {
		ret = v
	}
	return ret
}

func (m SeqString) GetUidMod4Seq2Str() string {
	ret := ""
	rs := []rune(m)
	l := len(rs)
	if l <= 3 {
		ret = string(rs[:])
		return ret
	}
	s := l - 3
	ret = string(rs[s:l])
	return ret
}
