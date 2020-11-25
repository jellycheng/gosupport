package dbutils

import "strings"

//https://dev.mysql.com/doc/refman/5.6/en/delete.html

//DELETE FROM [表名] [WHERE 子句] [ORDER BY ...] [LIMIT row_count]

type SQLBuilderDelete struct {
	//表名
	table string
	where string
	orderBy string
	limit string
	//条件参数值
	whereParams []interface{}
}

func (sqlb *SQLBuilderDelete) SetTable(tbl string) *SQLBuilderDelete {
	sqlb.table = tbl
	return sqlb
}

func (sqlb *SQLBuilderDelete) SetLimit(rowCount string) *SQLBuilderDelete {
	sqlb.limit = rowCount
	return sqlb
}

func (sqlb *SQLBuilderDelete) SetOrderBy(order string) *SQLBuilderDelete {
	sqlb.orderBy = order
	return sqlb
}

//考虑重复调用
func (sqlb *SQLBuilderDelete) Where(operator string, field string, condition string, value interface{}) *SQLBuilderDelete {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")  //AND、OR
	}

	buf.WriteString(WrapField(field))
	buf.WriteString(" " + condition + " ? ")
	sqlb.where = buf.String()
	sqlb.whereParams = append(sqlb.whereParams, value)

	return sqlb
}


func (sqlb *SQLBuilderDelete) AndWhere(field string, condition string, value interface{}) *SQLBuilderDelete {
	return sqlb.Where("AND", field, condition, value)
}

func (sqlb *SQLBuilderDelete) OrWhere(field string, condition string, value interface{}) *SQLBuilderDelete {
	return sqlb.Where("OR", field, condition, value)
}

func (sqlb *SQLBuilderDelete)GetWhereParamValues() []interface{} {
	return sqlb.whereParams
}

func (sqlb *SQLBuilderDelete) GetSQL() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}

	var buf strings.Builder
	buf.WriteString("DELETE FROM ")
	buf.WriteString(WrapTable(sqlb.table))
	if sqlb.where != "" {
		buf.WriteString(" WHERE " + sqlb.where)
	}
	if sqlb.orderBy != "" {
		buf.WriteString(" ORDER BY " + sqlb.orderBy)
	}
	if sqlb.limit != "" {
		buf.WriteString(" LIMIT " + sqlb.limit)
	}

	return buf.String(), nil
}

func NewSQLBuilderDelete() *SQLBuilderDelete {
	return &SQLBuilderDelete{}
}

