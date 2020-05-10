package dbutils

import (
	"strconv"
)

//分库分表算法


//分库分表数量结构体
type SplitDbAndTbl struct {
	totalDbNum int64 //总库数,必须大于0
	totalTblNum int64 //总表数,必须大于0
}

func NewSplitDbAndTbl(dbnum, tblnum int64) SplitDbAndTbl {
	if dbnum <=0 {
		dbnum = 1
	}
	if tblnum <=0 {
		tblnum = 1
	}
	return SplitDbAndTbl{
		totalDbNum:dbnum,
		totalTblNum:tblnum,
	}
}

func (t SplitDbAndTbl)GetTotalDbNum() int64 {
	return t.totalDbNum
}

func (t SplitDbAndTbl)GetTotalTblNum() int64 {
	return t.totalTblNum
}

func (t SplitDbAndTbl)GetDbNum(sVal string) (dbnum int64, err error) {
	ival, err := strconv.ParseInt(sVal, 10, 64)
	if err != nil {
		return 0, err
	} else {
		dbnum = (ival/t.totalDbNum) % t.totalDbNum + 1
	}
	return
}

func (t SplitDbAndTbl)GetTblNum(sVal string) (tblnum int64, err error) {
	ival, err := strconv.ParseInt(sVal, 10, 64)
	if err != nil {
		return 0, err
	} else {
		tblnum = ival % t.totalTblNum + 1
	}
	return
}


func (t SplitDbAndTbl)GetDBBaseByUserid(userid uint64) map[string]int64 {
	var ret map[string]int64 = map[string]int64 {
		"db_sn":1,
		"tbl_sn":1,
	}
	dbN ,err:=t.GetDbNum(strconv.FormatUint(userid, 10))
	if err==nil {
		ret["db_sn"] = dbN
	}

	tbN ,err:=t.GetTblNum(strconv.FormatUint(userid, 10))
	if err==nil {
		ret["tbl_sn"] = tbN
	}

	return ret
}

//字符串型的用户ID
func (t SplitDbAndTbl)GetDBBaseByStrUserid(userid string) map[string]int64 {
	var ret map[string]int64 = map[string]int64 {
		"db_sn":1,
		"tbl_sn":1,
	}

	newUserid := GetHashOrd(userid)
	//int64到string
	strUserid :=strconv.FormatInt(newUserid,10)
	dbN ,err:=t.GetDbNum(strUserid)
	if err==nil {
		ret["db_sn"] = dbN
	}

	tbN ,err:=t.GetTblNum(strUserid)
	if err==nil {
		ret["tbl_sn"] = tbN
	}

	return ret
}
