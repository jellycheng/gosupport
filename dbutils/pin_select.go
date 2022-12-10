package dbutils

import "strings"

//https://dev.mysql.com/doc/refman/5.6/en/select.html

//SELECT [字段] FROM [表名] [inner|left|right JOIN 子句] [WHERE 子句] [GROUP BY 子句] [HAVING 子句] [ORDER BY 子句] [LIMIT 子句]

type SQLBuilderSelect struct {
	//select字段
	selectField string
	//表名
	table   string
	join    string
	where   string
	groupBy string
	having  string
	orderBy string
	limit   string
	offset  string
	//条件参数值
	whereParams  []interface{}
	havingParams []interface{}
	limitParams  []interface{}
	joinParams   []interface{}
}

//Table("`t_user_ext` as t1")
func (sqlb *SQLBuilderSelect) SetTable(tbl string) *SQLBuilderSelect {
	sqlb.table = tbl
	return sqlb
}

func (sqlb *SQLBuilderSelect) Select(fields ...string) *SQLBuilderSelect {
	var buf strings.Builder
	fieldLen := len(fields)
	for k, field := range fields {
		buf.WriteString(field)
		if k != fieldLen-1 {
			buf.WriteString(",")
		}
	}
	sqlb.selectField = buf.String()
	return sqlb
}

//考虑重复调用,operator=AND、OR, condition="=,!="
func (sqlb *SQLBuilderSelect) ConditionWhere(operator string, field string, condition string, value interface{}) *SQLBuilderSelect {
	var buf strings.Builder
	buf.WriteString(sqlb.where)
	if buf.Len() != 0 {
		buf.WriteString(" " + operator + " ") //AND、OR
	}
	buf.WriteString(field + " " + condition + " ? ")
	sqlb.where = buf.String()
	sqlb.whereParams = append(sqlb.whereParams, value)

	return sqlb
}

func (sqlb *SQLBuilderSelect) Where(field string, condition string, value interface{}) *SQLBuilderSelect {
	return sqlb.ConditionWhere("AND", field, condition, value)
}

func (sqlb *SQLBuilderSelect) OrWhere(field string, condition string, value interface{}) *SQLBuilderSelect {
	return sqlb.ConditionWhere("OR", field, condition, value)
}

func (sqlb *SQLBuilderSelect) WhereIn(field string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.conditionIn("AND", field, "IN", values)
}

func (sqlb *SQLBuilderSelect) WhereNotIn(field string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.conditionIn("AND", field, "NOT IN", values)
}

func (sqlb *SQLBuilderSelect) OrWhereIn(field string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.conditionIn("OR", field, "IN", values)
}

func (sqlb *SQLBuilderSelect) OrWhereNotIn(field string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.conditionIn("OR", field, "NOT IN", values)
}

func (sqlb *SQLBuilderSelect) conditionIn(operator string, field string, condition string, values []interface{}) *SQLBuilderSelect {
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

//WhereRaw("`title` = ?", "hello")
func (sqlb *SQLBuilderSelect) WhereRaw(raw string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.Raw("AND", raw, values)
}

//OrWhereRaw("(`age` = ? OR `sex` = ?) AND `class` = ?", 7, 1, "一年级2班")
func (sqlb *SQLBuilderSelect) OrWhereRaw(raw string, values ...interface{}) *SQLBuilderSelect {
	return sqlb.Raw("OR", raw, values)
}

func (sqlb *SQLBuilderSelect) Raw(operator string, raw string, values []interface{}) *SQLBuilderSelect {
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

//JoinRaw("LEFT JOIN `t_user` as `t2` ON `t1`.`userid` = `t2`.`userid`").
func (sqlb *SQLBuilderSelect) JoinRaw(join string, values ...interface{}) *SQLBuilderSelect {
	var buf strings.Builder
	buf.WriteString(sqlb.join)
	if buf.Len() != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(join)
	sqlb.join = buf.String()
	for _, value := range values {
		sqlb.joinParams = append(sqlb.joinParams, value)
	}

	return sqlb
}

//GroupBy("`age`", "`id`")
func (sqlb *SQLBuilderSelect) GroupBy(fields ...string) *SQLBuilderSelect {
	var buf strings.Builder
	fieldLen := len(fields)
	for k, field := range fields {
		buf.WriteString(field)
		if k != fieldLen-1 {
			buf.WriteString(",")
		}
	}
	sqlb.groupBy = buf.String()
	return sqlb
}

func (sqlb *SQLBuilderSelect) OrderBy(order string) *SQLBuilderSelect {
	sqlb.orderBy = order
	return sqlb
}

func (sqlb *SQLBuilderSelect) GetWhereParamValues() []interface{} {
	return sqlb.whereParams
}

func (sqlb *SQLBuilderSelect) GetSQL() (string, error) {
	if sqlb.table == "" {
		return "", ErrTableEmpty
	}
	var buf strings.Builder

	buf.WriteString("SELECT ")
	if sqlb.selectField != "" {
		buf.WriteString(sqlb.selectField)
	} else {
		buf.WriteString("*")
	}
	buf.WriteString(" FROM ")
	buf.WriteString(sqlb.table)
	if sqlb.join != "" {
		buf.WriteString(" ")
		buf.WriteString(sqlb.join)
	}
	if sqlb.where != "" {
		buf.WriteString(" WHERE ")
		buf.WriteString(sqlb.where)
	}
	if sqlb.groupBy != "" {
		buf.WriteString(" GROUP BY ")
		buf.WriteString(sqlb.groupBy)
	}
	if sqlb.having != "" {
		buf.WriteString(" HAVING ")
		buf.WriteString(sqlb.having)
	}
	if sqlb.orderBy != "" {
		buf.WriteString(" ORDER BY ")
		buf.WriteString(sqlb.orderBy)
	}
	if sqlb.limit != "" {
		buf.WriteString(" LIMIT ")
		buf.WriteString(sqlb.limit)
	}
	if sqlb.offset != "" {
		buf.WriteString(" OFFSET ")
		buf.WriteString(sqlb.offset)
	}

	return buf.String(), nil
}

func NewSQLBuilderSelect() *SQLBuilderSelect {
	return &SQLBuilderSelect{}
}
