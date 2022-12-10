package gosupport

import (
	"fmt"
	"testing"
)

func TestNewRedisGroupManage(t *testing.T) {
	redisCfgManage := NewRedisGroupManage()
	redisCfgManage.SetMap("common", map[string]interface{}{
		"host":     "127.0.0.1",
		"port":     "6379",
		"password": "abc123",
		"db":       "9",
		"prefix":   "mobile_api:common:",
		"desc":     "无法归类的业务模块",
	})
	fmt.Println("common组key前缀为：", redisCfgManage.GetPrefix("common"))

	redisCfgManage = NewRedisGroupManage() //验证多次调用也不影响原先注入的配置
	redisCfgManage.SetMap("user", map[string]interface{}{
		"host":     "127.0.0.2",
		"port":     "6380",
		"password": "abc123",
		"db":       "9",
		"prefix":   "mobile_api:user:",
		"desc":     "用户业务相关模块",
	})
	fmt.Println("user组key前缀为：", redisCfgManage.GetPrefix("user"))
	redisCfgManage.SetMap("sms_code", map[string]interface{}{
		"host":     "127.0.0.3",
		"port":     "6381",
		"password": "abc123",
		"db":       "9",
		"prefix":   "mobile_api:smscode:",
		"desc":     "缓存短信验证码",
	})

	if redisInfo, ok := redisCfgManage.Get("common"); ok == nil {
		fmt.Println("common host:", redisInfo.GetHost())
		fmt.Println(redisInfo.ToString())
	}

	if redisInfo, ok := redisCfgManage.Get("user"); ok == nil {
		fmt.Println("user host:", redisInfo.GetHost())
		fmt.Println(redisInfo.ToString())
	}

	xxxRedisInfo := NewRedisInfo()
	xxxRedisInfo.SetRedisInfo(map[string]interface{}{
		"host":     "10.1.0.119",
		"port":     "6379",
		"password": "abc123",
		"db":       "9",
		"prefix":   "xxx_service:userinfo:",
		"desc":     "xxx专用缓存",
	})
	redisCfgManage.Set("xxx", xxxRedisInfo)
	if redisInfo, ok := redisCfgManage.Get("xxx"); ok == nil {
		fmt.Println("xxx host:", redisInfo.GetHost())
		fmt.Println(redisInfo.ToString())
	}

}
