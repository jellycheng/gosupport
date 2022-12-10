package dbutils

import (
	"fmt"
	"strconv"
	"testing"
)

// go test -run="TestNewSplitDbAndTbl"
func TestNewSplitDbAndTbl(t *testing.T) {

	dt := NewSplitDbAndTbl(1, 16)
	userid := "1000000002"

	dbN, err := dt.GetDbNum(userid)
	fmt.Println("userid:", userid, "总库数:", dt.GetTotalDbNum(), "实际分配的库号：", dbN)
	if err != nil {
		fmt.Println(err.Error())
	}
	if dbN2, err := dt.GetDbNum4int64(1000000002); err == nil {
		fmt.Println("userid:1000000002", "总库数:", dt.GetTotalDbNum(), "实际分配的库号：", dbN2)
	}

	tblN1, err := dt.GetTblNum(userid)
	fmt.Println("userid:", userid, "总表数:", dt.GetTotalTblNum(), "实际分配的表号：", tblN1)
	if err != nil {
		fmt.Println(err.Error())
	}
	if tblN2, err := dt.GetTblNum4int64(1000000002); err == nil {
		fmt.Println("userid:1000000002", "总表数:", dt.GetTotalTblNum(), "实际分配的表号：", tblN2)
	}

	if tblN3, err := dt.GetTblNum4int64(9006923516854095567); err == nil {
		fmt.Println("userid:9006923516854095567", "总表数:", dt.GetTotalTblNum(), "实际分配的表号：", tblN3)
	}

	userid2, _ := strconv.ParseInt(userid, 10, 64)
	fmt.Println(userid2, dt.GetDBBaseByUserid(uint64(userid2)), dt.GetDBBaseByStrUserid(userid))

	strUserid2 := "abc123"
	fmt.Println(strUserid2, dt.GetDBBaseByStrUserid(strUserid2))
	fmt.Println(strUserid2+"的hash值%128=", GetHashOrd127(strUserid2))

	dt2 := NewSplitDbAndTbl(11, 13)
	fmt.Printf("%+v \n", dt2) //{11 13}

	// 购物车使用这种
	dt3 := NewSplitDbAndTbl(1, 16)
	strUserid := "5466935782086561357"
	fmt.Println(strUserid, dt3.GetDBBaseByStrUserid(strUserid))
	fmt.Println(strUserid+"的hash值%128=", GetHashOrd127(strUserid))

}
