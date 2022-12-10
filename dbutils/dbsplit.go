package dbutils

import (
	"strconv"
)

//分库分表算法

//分库分表数量结构体
type SplitDbAndTbl struct {
	totalDbNum  int64 //总库数,必须大于0，即最小1
	totalTblNum int64 //总表数,必须大于0，即最小1
}

func NewSplitDbAndTbl(dbnum, tblnum int64) SplitDbAndTbl {
	if dbnum <= 0 {
		dbnum = 1
	}
	if tblnum <= 0 {
		tblnum = 1
	}
	return SplitDbAndTbl{
		totalDbNum:  dbnum,
		totalTblNum: tblnum,
	}
}

func (t SplitDbAndTbl) GetTotalDbNum() int64 {
	return t.totalDbNum
}

func (t SplitDbAndTbl) GetTotalTblNum() int64 {
	return t.totalTblNum
}

// 入参sVal值是可以转为int64类型的字符串，获取分配的库号
func (t SplitDbAndTbl) GetDbNum(sVal string) (dbnum int64, err error) {
	ival, err := strconv.ParseInt(sVal, 10, 64)
	if err != nil {
		return 0, err
	} else {
		dbnum = (ival/t.totalDbNum)%t.totalDbNum + 1
	}
	return
}

func (t SplitDbAndTbl) GetDbNum4int64(ival int64) (dbnum int64, err error) {
	dbnum = (ival/t.totalDbNum)%t.totalDbNum + 1
	return
}

// 入参sVal值是可以转为int64类型的字符串，获取分配的表号
func (t SplitDbAndTbl) GetTblNum(sVal string) (tblnum int64, err error) {
	ival, err := strconv.ParseInt(sVal, 10, 64)
	if err != nil {
		return 0, err
	} else {
		tblnum = ival%t.totalTblNum + 1
	}
	return
}

func (t SplitDbAndTbl) GetTblNum4int64(ival int64) (tblnum int64, err error) {
	tblnum = ival%t.totalTblNum + 1
	return
}

// 返回分配的库号和表号
func (t SplitDbAndTbl) GetDBBaseByUserid(userid uint64) map[string]int64 {
	var ret = map[string]int64{
		"db_sn":  1,
		"tbl_sn": 1,
	}
	if dbN, err := t.GetDbNum(strconv.FormatUint(userid, 10)); err == nil {
		ret["db_sn"] = dbN
	}

	if tbN, err := t.GetTblNum(strconv.FormatUint(userid, 10)); err == nil {
		ret["tbl_sn"] = tbN
	}

	return ret
}

// 字符串型的用户ID，返回分配的库号和表号
func (t SplitDbAndTbl) GetDBBaseByStrUserid(userid string) map[string]int64 {
	var ret = map[string]int64{
		"db_sn":  1,
		"tbl_sn": 1,
	}

	newUserid := GetHashOrd(userid)
	// int64到string
	strUserid := strconv.FormatInt(newUserid%128, 10)
	if dbN, err := t.GetDbNum(strUserid); err == nil {
		ret["db_sn"] = dbN
	}

	if tbN, err := t.GetTblNum(strUserid); err == nil {
		ret["tbl_sn"] = tbN
	}

	return ret
}
