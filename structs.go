package gosupport

import (
	"reflect"
	"strconv"
)

//通过结构体的default标签设置结构体默认值,入参为指针类型
func InitStruct4DefaultTag(bean interface{}) {
	configType := reflect.TypeOf(bean) //返回reflect.Type接口类型
	for i := 0; i < configType.Elem().NumField(); i++ {
		field := configType.Elem().Field(i)
		defaultValue := field.Tag.Get("default")
		if defaultValue == "" {
			continue
		}
		setter := reflect.ValueOf(bean).Elem().Field(i)
		switch field.Type.String() {
		case "int", "int8", "int16", "int32", "int64":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "time.Duration":
			intValue, _ := strconv.ParseInt(defaultValue, 10, 64)
			setter.SetInt(intValue)
		case "string":
			setter.SetString(defaultValue)
		case "bool":
			boolValue, _ := strconv.ParseBool(defaultValue)
			setter.SetBool(boolValue)
		case "uint", "uint8", "uint16", "uint32", "uint64":
			uintValue, _ := strconv.ParseUint(defaultValue, 10, 64)
			setter.SetUint(uintValue)
		case "float32":
			float32Value, _ := strconv.ParseFloat(defaultValue, 32)
			setter.SetFloat(float32Value)
		case "float64":
			float64Value, _ := strconv.ParseFloat(defaultValue, 64)
			setter.SetFloat(float64Value)
		}
	}
}

// 返回空结构体
func EmptyStruct() struct{} {
	return struct {
	}{}
}

// 返回空结构体切片，用于json序列化时返回空数组[]
func EmptyStructSlice() []struct{} {
	return []struct {
	}{}
}

// 获取结构体某字段的tag值
func GetStructTagContent(i interface{}, fieldNname string, tagName string) (string, bool) {
	typeOfI := reflect.TypeOf(i)
	for typeOfI.Kind() == reflect.Ptr {
		typeOfI = typeOfI.Elem()
	}
	if iType, ok := typeOfI.FieldByName(fieldNname); ok {
		return iType.Tag.Lookup(tagName)
	}
	return "", false
}
