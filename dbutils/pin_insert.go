package dbutils

import "strings"

//https://dev.mysql.com/doc/refman/5.6/en/insert.html

type SQLBuilderInsert struct {
	//表名
	table string
	//格式：([字段名]) VALUES ([?,?,?]) 示例：(`email`,`username`,`qq`) VALUES (?,?,?)
	insert string
	//插入参数值
	insertParams []interface{}
}

func (sqlb *SQLBuilderInsert) SetTable(tbl string) *SQLBuilderInsert {
	sqlb.table = tbl
	return sqlb
}

//SetInsertData([]string{"`name`", "`age`"}, "tom", 18)
func (sqlb *SQLBuilderInsert) SetInsertData(fileds []string, values ...interface{}) *SQLBuilderInsert {
	fieldLen := len(fileds)
	if fieldLen == 0 {
		return sqlb
	}
	var buf strings.Builder
	wenHao := []string{}
	buf.WriteString("(")
	for k, field := range fileds {
		wenHao = append(wenHao, "?")
		buf.WriteString(field)
		if k != fieldLen-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(") VALUES (")
	buf.WriteString(strings.Join(wenHao, ","))
	buf.WriteString(")")

	sqlb.insert = buf.String()

	for k, value := range values {
		if k >= fieldLen {
			break
		}
		sqlb.insertParams = append(sqlb.insertParams, value)
	}

	return sqlb
}

func (sqlb *SQLBuilderInsert) GetParamValues() []interface{} {
	return sqlb.insertParams
}

//INSERT INTO [表名] ([字段名]) VALUES ([要插入的值]): INSERT INTO `bbs_user`(`email`,`username`,`qq`) VALUES (?,?,?)
func (sqlb *SQLBuilderInsert) GetSql() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}
	if sqlb.insert == "" {
		return "", ErrInsertEmpty
	}

	var buf strings.Builder
	buf.WriteString("INSERT INTO ")
	buf.WriteString(sqlb.table)
	buf.WriteString(" ")
	buf.WriteString(sqlb.insert)

	return buf.String(), nil
}

func NewSQLBuilderInsert() *SQLBuilderInsert {
	return &SQLBuilderInsert{}
}
