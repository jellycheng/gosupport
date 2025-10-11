package dbutils

import "strings"

//https://dev.mysql.com/doc/refman/5.6/en/update.html

type SQLBuilderUpdate struct {
	//表名
	table   string
	set     string
	where   string
	orderBy string
	limit   string
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

func (sqlb *SQLBuilderUpdate) OrderBy(order string) *SQLBuilderUpdate {
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
		buf.WriteString(field)
		buf.WriteString(" = ? ")
		if k != fieldLen-1 {
			buf.WriteString(",")
		}
	}

	sqlb.set = buf.String()

	for k, value := range values {
		if k >= fieldLen {
			break
		}
		sqlb.setParams = append(sqlb.setParams, value)
		sqlb.allParams = append(sqlb.allParams, value)
	}

	return sqlb
}

// 考虑重复调用,operator=AND、OR,condition=<、<=、=、!=、>
func (sqlb *SQLBuilderUpdate) ConditionWhere(operator string, field string, condition string, value interface{}) *SQLBuilderUpdate {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")
	}
	buf.WriteString(field)
	buf.WriteString(" " + condition + " ? ")
	sqlb.where = buf.String()

	sqlb.whereParams = append(sqlb.whereParams, value)
	sqlb.allParams = append(sqlb.allParams, value)

	return sqlb
}

func (sqlb *SQLBuilderUpdate) Where(field string, condition string, value interface{}) *SQLBuilderUpdate {
	return sqlb.ConditionWhere("AND", field, condition, value)
}

func (sqlb *SQLBuilderUpdate) OrWhere(field string, condition string, value interface{}) *SQLBuilderUpdate {
	return sqlb.ConditionWhere("OR", field, condition, value)
}

func (sqlb *SQLBuilderUpdate) WhereIn(field string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.conditionIn("AND", field, "IN", values)
}

func (sqlb *SQLBuilderUpdate) WhereNotIn(field string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.conditionIn("AND", field, "NOT IN", values)
}

func (sqlb *SQLBuilderUpdate) OrWhereIn(field string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.conditionIn("OR", field, "IN", values)
}

func (sqlb *SQLBuilderUpdate) OrWhereNotIn(field string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.conditionIn("OR", field, "NOT IN", values)
}

func (sqlb *SQLBuilderUpdate) conditionIn(operator string, field string, condition string, values []interface{}) *SQLBuilderUpdate {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")
	}
	s, _ := PinConditionIn(field, condition, values)
	buf.WriteString(s)

	sqlb.where = buf.String()

	for _, value := range values {
		sqlb.whereParams = append(sqlb.whereParams, value)
	}

	return sqlb
}

// WhereRaw("`title` = ?", "hello")
func (sqlb *SQLBuilderUpdate) WhereRaw(raw string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.Raw("AND", raw, values)
}

// OrWhereRaw("(`age` = ? OR `sex` = ?) AND `class` = ?", 7, 1, "一年级2班")
func (sqlb *SQLBuilderUpdate) OrWhereRaw(raw string, values ...interface{}) *SQLBuilderUpdate {
	return sqlb.Raw("OR", raw, values)
}

func (sqlb *SQLBuilderUpdate) Raw(operator string, raw string, values []interface{}) *SQLBuilderUpdate {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ")
	}

	buf.WriteString(raw)
	sqlb.where = buf.String()

	for _, value := range values {
		sqlb.whereParams = append(sqlb.whereParams, value)
		sqlb.allParams = append(sqlb.allParams, value)
	}

	return sqlb
}

func (sqlb *SQLBuilderUpdate) GetParamValues() []interface{} {
	return sqlb.allParams
}

func (sqlb *SQLBuilderUpdate) GetSetParamValues() []interface{} {
	return sqlb.setParams
}

func (sqlb *SQLBuilderUpdate) GetWhereParamValues() []interface{} {
	return sqlb.whereParams
}

// UPDATE [表名] [SET 子句] [WHERE 子句] [ORDER BY ...] [LIMIT row_count]
func (sqlb *SQLBuilderUpdate) GetSQL() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}

	if sqlb.set == "" {
		return "", ErrUpdateEmpty
	}

	var buf strings.Builder

	buf.WriteString("UPDATE ")
	buf.WriteString(sqlb.table)
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
