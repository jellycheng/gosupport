package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
)

// dsn := GetDsn(map[string]interface{}{"username":"root","password":"88888888"})
func GetDsn(dbConfig map[string]interface{}) string {
	user_name, ok := dbConfig["username"]
	if !ok {
		user_name = "root"
	}
	password, ok := dbConfig["password"]
	if !ok {
		password = ""
	}
	host, ok := dbConfig["host"]
	if !ok {
		host = "localhost"
	}
	port, ok := dbConfig["port"]
	if !ok {
		port = "3306"
	}
	dbname, ok := dbConfig["dbname"]
	if !ok {
		dbname = ""
	}
	charset, ok := dbConfig["charset"]
	if !ok {
		charset = "utf8"
	}

	param, ok := dbConfig["extparam"]
	if ok && param != "" {
		param = fmt.Sprintf("&%s", param)
	} else {
		param = "";
	}
	dbDsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s%s", user_name, password, host, port, dbname, charset,param)
	return dbDsn
}


// 获取表字段信息
func GetTableFields(connect *sql.DB, tbl string) ([]map[string]string, error) {
	//结果
	ret := []map[string]string{}
	sqlStr := fmt.Sprintf("show full fields from %s;", tbl);
	rows, err:=connect.Query(sqlStr)
	if err !=nil {
		return ret, err
	}
	defer rows.Close()
	cols,err := rows.Columns() //获取所有字段
	//一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));
	//填充一行数据
	scans := make([]interface{}, len(cols));
	for k, _ := range vals {
		scans[k] = &vals[k];
	}
	for rows.Next() {
		//每行数据
		row := make(map[string]string)
		if err := rows.Scan(scans...);err ==nil {
			for k, v := range vals { //遍历结果
				key := cols[k]; //字段名
				//把[]byte数据转成string,放入结果集
				row[key] = string(v)
			}
			ret = append(ret, row)

		} else {
			return ret, err
		}
	}
	return ret, nil

}


/*
  执行插入记录sql
  示例：insSql := "insert INTO t_user_token_1(user_id,user_token) values(?,?)"
		id,err := InsertSql(connect, insSql, "1000", "token123")
*/
func InsertSql(connect *sql.DB, insSql string, args ...interface{}) (id int64, err error)  {
	id = 0
	res, err := connect.Exec(insSql, args...)
	if err !=nil {
		return
	}
	id, err = res.LastInsertId()
	return
}

/*
  更新记录sql
  示例：upSql := "update t_user_token_1 set user_token=?,update_time=? where user_id=? order by id desc limit 1"
	id,err := UpdateSql(connect, upSql,  "token123abc222", time.Now().Unix(),"1000")
*/
func UpdateSql(connect *sql.DB, insSql string, args ...interface{}) (num int64, err error)  {
	num = 0
	res, err := connect.Exec(insSql, args...)
	if err !=nil {
		return
	}
	num, err = res.RowsAffected()
	return
}

// 删除记录sql
func DeleteSql(connect *sql.DB, insSql string, args ...interface{}) (num int64, err error)  {
	num = 0
	res, err := connect.Exec(insSql, args...)
	if err !=nil {
		return
	}
	num, err = res.RowsAffected()
	return
}


// 查询单条记录
func SelectOne(connect *sql.DB, sqlStr string, args ...interface{}) (map[string]string, error)  {
	//结果
	ret := make(map[string]string)
	rows,err := connect.Query(sqlStr, args...)
	if err !=nil {
		return ret, err
	}
	defer rows.Close()
	//查询到的字段名，返回的是一个string数组切片
	cols, err := rows.Columns()
	//一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));
	//填充一行数据
	scans := make([]interface{}, len(cols));
	for k, _ := range vals {
		scans[k] = &vals[k];
	}

	for rows.Next() {
		if err := rows.Scan(scans...);err ==nil {
			for k, v := range vals { //遍历结果
				key := cols[k]; //字段名
				//把[]byte数据转成string,放入结果集
				ret[key] = string(v)
			}
			break
		} else {
			return ret, err
		}
	}
	return ret, nil

}

// 查询多条记录
func SelectRows(connect *sql.DB, sqlStr string, args ...interface{}) ([]map[string]string, error)  {
	//结果
	ret := []map[string]string{}

	rows,err := connect.Query(sqlStr, args...)
	if err !=nil {
		return ret, err
	}
	defer rows.Close()
	//查询到的字段名，返回的是一个string数组切片
	cols, err := rows.Columns()
	//一行所有列的值，用[]byte表示
	vals := make([][]byte, len(cols));
	//填充一行数据
	scans := make([]interface{}, len(cols));
	for k, _ := range vals {
		scans[k] = &vals[k];
	}

	for rows.Next() {
		//每行数据
		row := make(map[string]string)
		if err := rows.Scan(scans...);err ==nil {
			for k, v := range vals { //遍历结果
				key := cols[k]; //字段名
				//把[]byte数据转成string,放入结果集
				row[key] = string(v)
			}
			ret = append(ret, row)

		} else {
			return ret, err
		}
	}
	return ret, nil
}

// 连接db
func DbConnect(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn) //返回*sql.DB结构体指针类型对象
	if err != nil {
		return nil, errors.New("db连接发生错误")
	}
	return db, nil
}

var dbObj = make(map[string]*sql.DB)

func GetDbConnect(dbnameCode, dsn string) (*sql.DB, error)  {
	db,ok := dbObj[dbnameCode]
	if ok {
		return db,nil
	}
	if db, err := DbConnect(dsn); err==nil{
		dbObj[dbnameCode] = db
		return db, nil
	} else {
		return db,err
	}
}


// 获取所有表名
func ShowTables(connect *sql.DB) ([]string) {
	ret := []string{}
	fieldName := ""
	sql := "show tables;"
	tables, err := SelectRows(connect, sql)
	if err == nil {
		for _,v := range tables {
			for field, _ := range v {
				fieldName = field
				break
			}
			if fieldName != "" {
				break
			}
		}

		for _,v := range tables {
			tblName,ok := v[fieldName]
			if ok {
				ret = append(ret, tblName)
			}
		}
	}
	return ret
}

// 获取创建表sql语句
func ShowCreateTable(connect *sql.DB, tblName string) (string, error) {
	ret := ""
	sqlStr := fmt.Sprintf("show create table `%s`;", tblName)
	res,err := SelectOne(connect, sqlStr)
	if err != nil {
		return ret, err
	}
	tblSql,ok := res["Create Table"]
	if ok {
		return tblSql, nil
	}
	return ret, errors.New("获取表结构失败")
}
