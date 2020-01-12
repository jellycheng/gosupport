package ini

//*.ini配置文件解析

type config struct {
	inifile string                         //ini文件
	conflist []map[string]map[string]string //配置
}

var iniCfg map[string]*config= make(map[string]*config)

func LoadIni(configName, iniFile string)  {
	iniCfg[configName] = new(config)
	iniCfg[configName].inifile = iniFile
	//分析配置文件
	
}

func GetIniGroupCfg(k string) (cfg *config, ok bool)  {
	cfg, ok = iniCfg[k]
	return
}

//ini.GetVal("redis.USER_REDIS_HOST", "127.0.0.1")
func GetVal(k string, defaultVal interface{}) (ret interface{}, err error)  {


	return
}

func SetVal(k string, v interface{}) (err error)  {

	return
}

func DeleteVal(k string) (err error)  {

	return err
}

