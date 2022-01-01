package gosupport

//接口定义文件

//标识性接口
type IdentifyInterface interface {

}

type StrMaper interface {
	// 获取值
	Get(key string) (string, bool)
	// 设置值
	Set(key, value string)
	// 删除
	Del(key string)
	// 克隆
	Clone() StrMaper
}
