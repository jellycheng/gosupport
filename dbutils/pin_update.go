package dbutils

import "strings"

//https://dev.mysql.com/doc/refman/5.6/en/update.html

type SQLBuilderUpdate struct {
	//表名
	table string

	set string
	where string
	orderBy string
	limit string
	//更新参数值
	setParams []interface{}
	//条件参数值
	whereParams []interface{}
	//set + where参数
	allParams []interface{}
}

func (sqlb *SQLBuilderUpdate) SetTable(tbl string) *SQLBuilderUpdate {
	sqlb.table = tbl
	return sqlb
}

func (sqlb *SQLBuilderUpdate) SetLimit(rowCount string) *SQLBuilderUpdate {
	sqlb.limit = rowCount
	return sqlb
}

func (sqlb *SQLBuilderUpdate) SetOrderBy(order string) *SQLBuilderUpdate {
	sqlb.orderBy = order
	return sqlb
}

func (sqlb *SQLBuilderUpdate) SetUpdateData(fileds []string, values ...interface{}) *SQLBuilderUpdate {
	fieldLen := len(fileds)
	if fieldLen == 0 {
		return sqlb
	}
	var buf strings.Builder
	for k, field := range fileds {
		buf.WriteString(WrapField(field))
		buf.WriteString(" = ? ")
		if k != fieldLen-1 {
			buf.WriteString(",")
		}
	}

	sqlb.set = buf.String()

	for k, value := range values {
		if k>=fieldLen {
			break
		}
		sqlb.setParams = append(sqlb.setParams, value)
		sqlb.allParams = append(sqlb.allParams, value)
	}

	return sqlb
}

//考虑重复调用
func (sqlb *SQLBuilderUpdate) Where(operator string, field string, condition string, value interface{}) *SQLBuilderUpdate {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" ")
		buf.WriteString(operator) //AND、OR
		buf.WriteString(" ")
	}

	buf.WriteString(WrapField(field))
	//=、!=、in
	buf.WriteString(" " + condition + " ? ")
	sqlb.where = buf.String()

	sqlb.whereParams = append(sqlb.whereParams, value)
	sqlb.allParams = append(sqlb.allParams, value)

	return sqlb
}

func (sqlb *SQLBuilderUpdate) AndWhere(field string, condition string, value interface{}) *SQLBuilderUpdate {
	return sqlb.Where("AND", field, condition, value)
}

func (sqlb *SQLBuilderUpdate) OrWhere(field string, condition string, value interface{}) *SQLBuilderUpdate {
	return sqlb.Where("OR", field, condition, value)
}

func (sqlb *SQLBuilderUpdate)GetParamValues() []interface{} {
	return sqlb.allParams
}

func (sqlb *SQLBuilderUpdate)GetSetParamValues() []interface{} {
	return sqlb.setParams
}

func (sqlb *SQLBuilderUpdate)GetWhereParamValues() []interface{} {
	return sqlb.whereParams
}

//UPDATE [表名] [SET 子句] [WHERE 子句] [ORDER BY ...] [LIMIT row_count]
func (sqlb *SQLBuilderUpdate) GetSQL() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}

	if sqlb.set == "" {
		return "", ErrUpdateEmpty
	}

	var buf strings.Builder

	buf.WriteString("UPDATE ")
	buf.WriteString(WrapTable(sqlb.table))
	buf.WriteString(" SET ")
	buf.WriteString(sqlb.set)
	if sqlb.where != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(sqlb.where)
	}
	if sqlb.orderBy != "" {
		buf.WriteString(" ORDER BY ")
		buf.WriteString(sqlb.orderBy)
	}
	if sqlb.limit != "" {
		buf.WriteString(" LIMIT ")
		buf.WriteString(sqlb.limit)
	}

	return buf.String(), nil
}

func NewSQLBuilderUpdate() *SQLBuilderUpdate {
	return &SQLBuilderUpdate{}
}

