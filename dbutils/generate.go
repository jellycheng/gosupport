package dbutils

import (
	"database/sql"
	"fmt"
	"strings"
)

// 生成db模型代码，如gorm的model结构体

const (
	//fieldSqlFormat = "show FULL COLUMNS from %s;"
	tblInfoSqlFormat4Db = "SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema='%s';"
	tblInfoSqlFormat4DbAndTbl = "SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES WHERE table_schema='%s' AND TABLE_NAME IN (%s);"

	jsonTagFormat = `json:"%s"`

)

// 生成gorm的model结构体
func GenerateGormModel(connect *sql.DB, dbName string, tableNames []string, cfg map[string]string) map[string]string {
	// 表名 => gorm文件内容
	ret := make(map[string]string)
	// 获取要生成的表
	tblsInfo := GetSimpleTableInfo(connect, dbName, tableNames)
	for _, val := range tblsInfo {
		tblName := val["Name"]  //表名
		//获取字段信息
		fieldsInfo,_ := GetTableFields(connect, tblName)
		ret[tblName] = GormModelFormat(val, fieldsInfo, cfg)

	}
	return ret
}


func GetSimpleTableInfo(connect *sql.DB, dbName string, tableNames []string) []map[string]string {
	var ret []map[string]string
	sqlStr := ""
	if len(tableNames) == 0 {
		sqlStr = fmt.Sprintf(tblInfoSqlFormat4Db, dbName)
	} else {
		sqlStr = fmt.Sprintf(tblInfoSqlFormat4DbAndTbl, dbName, PinInStr(tableNames))
	}
	if tablesInfo, err := SelectRows(connect, sqlStr);err == nil {
		ret = tablesInfo
	}
	return ret
}

// table=["Name":"表名", "Comment":"表注解"], cfg=["packageName":"包名", "ignoreField":"id,is_delete,create_time,update_time,delete_time,modify_time","appendStructCode":"GormCommonField","trimTblPrefix":"t_","structNameSuffix":"Model"]
func GormModelFormat(table map[string]string, fields []map[string]string, cfg map[string]string) string {
	content := ""
	if packageName, ok := cfg["packageName"];ok && packageName != "" {
		content = fmt.Sprintf("package %s\n\n", packageName)
	} else {
		content = "package models\n\n"
	}
	//表注释
	if tblComment,ok := table["Comment"];ok && tblComment != ""  {
		content += "// "+tblComment+"\n"
	}
	ignoreField := []string{}
	if tmpIgnoreField,ok := cfg["ignoreField"];ok {
		ignoreField = strings.Split(tmpIgnoreField, ",")
	}
	tblName := table["Name"]
	structName := ""
	if trimTblPrefix, ok := cfg["trimTblPrefix"];ok && trimTblPrefix !="" {
		structName = CamelCase(strings.TrimPrefix(tblName, trimTblPrefix))
	} else {
		structName = CamelCase(tblName)
	}
	if structNameSuffix,ok := cfg["structNameSuffix"]; ok && structNameSuffix != "" {
		structName = structName + structNameSuffix
	}

	content += "type " + structName + " struct {\n"
	if tmpAppendCode,ok := cfg["appendStructCode"];ok && tmpAppendCode != "" {//追加到结构体中的代码
		content += "	" + tmpAppendCode + "\n"
	}

	//生成字段
	for _, field := range fields {
		fieldName := field["Field"] //字段名
		if StrInSlice(fieldName, ignoreField) {//忽略的字段
			continue
		}
		fieldComment := GetFieldComment(field) //字段注解
		fieldJson := fmt.Sprintf(jsonTagFormat, fieldName)
		fieldGoType := FiledType2GoType(field["Type"])
		primaryKeyTag := ""  //primary_key;
		if primaryKeyVal, isKey := field["Key"];isKey && primaryKeyVal == "PRI" {
			primaryKeyTag = "primary_key;"
		}
		gormColumn := fmt.Sprintf(`gorm:"%sColumn:%s"`, primaryKeyTag, fieldName)
		content += "	" + CamelCase(fieldName) + " " + fieldGoType + " `" + gormColumn + " " + fieldJson + "` "+fieldComment+"\n"
	}
	content += "}"
	content += fmt.Sprintf(`

func (%s)TableName() string {
	return "%s"
}
`, structName, tblName)

	return content
}

//获取字段说明
func GetFieldComment(field map[string]string) string{
	if comment,ok := field["Comment"]; ok && comment != "" {
		return "// "+comment
	}
	return ""
}


// 生成dto结构体
func GenerateDto(connect *sql.DB, tableName string, cfg map[string]string) string {
	ret := "package dto\n\n"
	structName := "Dto"
	if tmp,ok := cfg["structName"];ok && tmp != "" {
		structName = tmp
	}
	ret += "type " + structName + " struct {\n"
	fieldsInfo,_ := GetTableFields(connect, tableName)
	ignoreField := []string{}
	if tmpIgnoreField,ok := cfg["ignoreField"];ok {
		ignoreField = strings.Split(tmpIgnoreField, ",")
	}

	//生成字段
	for _, field := range fieldsInfo {
		fieldName := field["Field"] //字段名
		if StrInSlice(fieldName, ignoreField) {//忽略的字段
			continue
		}
		fieldComment := GetFieldComment(field) //字段注解
		fieldJson := fmt.Sprintf(jsonTagFormat, fieldName)
		fieldGoType := FiledType2GoType(field["Type"])
		ret += "	" + CamelCase(fieldName) + " " + fieldGoType + " `" + fieldJson + "` " + fieldComment + "\n"
	}
	ret += "}"

	return ret
}
