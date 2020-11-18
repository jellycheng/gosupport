package ini

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

/**
ini文件格式示例
# 注释1
; 注释2
[database]
hostname = localhost
username = root
password = 123456
port=3306

[redis]
host = 127.0.0.1
port = 6379
db = 9

[default]
app_env = dev
app_anme = order-service

 */


type Config struct {
	iniFile string                         //ini文件
	configList map[string]map[string]string //配置
}

func (c *Config) GetIniFile() string {
	return c.iniFile
}

func (c *Config) GetConfigData() map[string]map[string]string {
	return c.configList
}

func (c *Config) ParseIniFile() (error){
	var (
		groupName = DefaultGroupName                          //组名
		data      = make(map[string]map[string]string) //组下的key-value配置
	)
	data[groupName] = make(map[string]string)  //默认组

	fileObj, err := os.Open(c.iniFile)
	if err != nil {
		return err
	}
	defer fileObj.Close()

	buf := bufio.NewReader(fileObj)
	for {
		lineByte,_,err := buf.ReadLine()
		if err==io.EOF {
			break
		}
		if err!=nil && err!=io.EOF {//读取内容发生错误
			break
		}
		line := strings.TrimSpace(string(lineByte))
		switch {
		case len(line) == 0:
		case string(line[0]) == ";":
		case string(line[0]) == "#":	//增加配置文件备注
		case line[0] == '[' && line[len(line)-1] == ']': //分组
			groupName = strings.TrimSpace(line[1 : len(line)-1])
			if _,ok :=data[groupName];ok == false {
				data[groupName] = make(map[string]string)
			}

		default:
			i := strings.IndexAny(line, "=")
			if i == -1 {
				continue
			}
			value := strings.TrimSpace(string(GetCleanComment([]byte(line[i+1 : len(line)]))))
			data[groupName][strings.TrimSpace(line[0:i])] = value

		}

	}
	c.configList = data
	return nil
}

func (c *Config) CheckGroupExists(groupName string) bool {
	for k, _ := range c.configList {
		if k == groupName {
			return true
		}
	}
	return false
}

func (c *Config) GetValue(groupName, keyName string) (string,error) {
	conf := c.configList
	if retGroupVal,ok :=conf[groupName];ok {
		if retV,ok :=retGroupVal[keyName];ok {
			return retV, nil
		}
		return "",errors.New(fmt.Sprintf("配置%s组下的%s key不存在", groupName, keyName))
	} else {
		return "", errors.New(fmt.Sprintf("%s配置组不存在", groupName))
	}

}

func (c *Config) MustGetValue(groupName, keyName string) string  {
	val,_ := c.GetValue(groupName, keyName)
	return val
}

func (c *Config) DelValue(groupName, keyName string) bool {
	data := c.configList
	if retGroupVal,ok :=data[groupName];ok {
		if _,ok :=retGroupVal[keyName];ok {
			delete(c.configList[groupName], keyName)
			return true
		}
	}
	return false
}

func NewIniConfig(iniFile string) *Config {
	c := new(Config)
	c.iniFile = iniFile
	return c
}

