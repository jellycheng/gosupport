package dbutils

import (
	"fmt"
	"testing"
)

func TestNewSplitDbAndTbl(t *testing.T) {

	dt := NewSplitDbAndTbl(8, 10)
	dt2 := NewSplitDbAndTbl(11, 13)
	fmt.Println(dt2)  //{11 13}

	userid := "1000000002"
	dbN ,err:=dt.GetDbNum(userid)
	fmt.Println("userid:", userid, "totaldbnum:", dt.GetTotalDbNum(),"库号：", dbN)
	if err != nil {
		fmt.Println(err.Error())
	}
	dbT ,err:=dt.GetTblNum(userid)
	fmt.Println("userid:", userid, "totaltblnum:", dt.GetTotalTblNum(),"表号：", dbT)
	if err != nil {
		fmt.Println(err.Error())
	}

}
