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

func (sqlb *SQLBuilderDelete) OrderBy(order string) *SQLBuilderDelete {
	sqlb.orderBy = order
	return sqlb
}

//考虑重复调用,operator=AND、OR, condition="=,!="
func (sqlb *SQLBuilderDelete) ConditionWhere(operator string, field string, condition string, value interface{}) *SQLBuilderDelete {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")  //AND、OR
	}
	buf.WriteString(field + " " + condition + " ? ")
	sqlb.where = buf.String()
	sqlb.whereParams = append(sqlb.whereParams, value)

	return sqlb
}


func (sqlb *SQLBuilderDelete) Where(field string, condition string, value interface{}) *SQLBuilderDelete {
	return sqlb.ConditionWhere("AND", field, condition, value)
}

func (sqlb *SQLBuilderDelete) OrWhere(field string, condition string, value interface{}) *SQLBuilderDelete {
	return sqlb.ConditionWhere("OR", field, condition, value)
}

func (sqlb *SQLBuilderDelete) WhereIn(field string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.conditionIn("AND", field, "IN", values)
}

func (sqlb *SQLBuilderDelete) WhereNotIn(field string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.conditionIn("AND", field, "NOT IN", values)
}

func (sqlb *SQLBuilderDelete) OrWhereIn(field string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.conditionIn("OR",  field,"IN", values)
}

func (sqlb *SQLBuilderDelete) OrWhereNotIn(field string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.conditionIn("OR",  field,"NOT IN", values)
}

func (sqlb *SQLBuilderDelete) conditionIn(operator string, field string,condition string, values []interface{}) *SQLBuilderDelete {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")
	}
	s,_ := PinConditionIn(field, condition, values)
	buf.WriteString(s)

	sqlb.where = buf.String()

	for _, value := range values {
		sqlb.whereParams = append(sqlb.whereParams, value)
	}

	return sqlb
}

//WhereRaw("`title` = ?", "hello")
func (sqlb *SQLBuilderDelete) WhereRaw(raw string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.Raw("AND", raw, values)
}

//OrWhereRaw("(`age` = ? OR `sex` = ?) AND `class` = ?", 7, 1, "一年级2班")
func (sqlb *SQLBuilderDelete) OrWhereRaw(raw string, values ...interface{}) *SQLBuilderDelete {
	return sqlb.Raw("OR", raw, values)
}

func (sqlb *SQLBuilderDelete) Raw(operator string, raw string, values []interface{}) *SQLBuilderDelete {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")
	}

	buf.WriteString(raw)
	sqlb.where = buf.String()

	for _, value := range values {
		sqlb.whereParams = append(sqlb.whereParams, value)
	}

	return sqlb
}

func (sqlb *SQLBuilderDelete)GetWhereParamValues() []interface{} {
	return sqlb.whereParams
}

func (sqlb *SQLBuilderDelete) GetSQL() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}

	var buf strings.Builder
	buf.WriteString("DELETE FROM " + sqlb.table)
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

