package dbutils

//https://dev.mysql.com/doc/refman/5.6/en/select.html

//SELECT [字段] FROM [表名] [inner|left|right JOIN 子句] [WHERE 子句] [GROUP BY 子句] [HAVING 子句] [ORDER BY 子句] [LIMIT 子句]

type SQLBuilderSelect struct {
	//表名
	table string
	join string
	where string
	groupBy string
	having  string
	orderBy string
	limit string
	offset string
	//条件参数值
	whereParams []interface{}
	havingParams []interface{}
	limitParams  []interface{}
	joinParams   []interface{}

}

func (sqlb *SQLBuilderSelect) SetTable(tbl string) *SQLBuilderSelect {
	sqlb.table = tbl
	return sqlb
}


func NewSQLBuilderSelect() *SQLBuilderSelect {
	return &SQLBuilderSelect{}
}

