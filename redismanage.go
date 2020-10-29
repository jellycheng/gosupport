package gosupport

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
)

//redis配置管理

type redisinfo struct {
	host     string   //redis host
	port     int      //redis端口
	password string   //redis密码
	db int            //redis db号
	prefix string     //redis key前缀
	desc string       //描述，缓存用户相关信息
}

func (r *redisinfo)GetHost() string  {
	host := r.host
	if host=="" {
		host = "127.0.0.1"
	}
	return host
}

func (r *redisinfo)SetHost(host string) *redisinfo  {
	r.host = host
	return r
}

func (r *redisinfo)GetPort() int  {
	port := r.port
	if port==0 {
		port = 6379
	}
	return port
}

func (r *redisinfo)SetPort(port int) *redisinfo  {
	r.port = port
	return r
}

func (r *redisinfo)GetPassword() string  {
	return r.password
}

func (r *redisinfo)SetPassword(password string) *redisinfo  {
	r.password = password
	return r
}

func (r *redisinfo)GetDb() int  {
	dbNum := r.db
	return dbNum
}

func (r *redisinfo)SetDb(dbNum int) *redisinfo  {
	r.db = dbNum
	return r
}

func (r *redisinfo)GetPrefix() string  {
	return r.prefix
}

func (r *redisinfo)SetPrefix(prefix string) *redisinfo  {
	r.prefix = prefix
	return r
}

func (r *redisinfo)GetDesc() string  {
	return r.desc
}

func (r *redisinfo)SetDesc(desc string) *redisinfo  {
	r.desc = desc
	return r
}

func (r *redisinfo)GetAddr() string  {
	return fmt.Sprintf("%s:%d", r.host, r.port)
}

func (r *redisinfo)SetRedisInfo(info map[string]interface{}) *redisinfo  {
	for k,v := range info {
		switch k {
		case "host":
			if val,ok:=v.(string);ok{
				r.host = val
			}

		case "port":
			if val,ok:=v.(int);ok{
				r.port = val
			} else if val2,ok2 := v.(string);ok2 {
				v2,_ := strconv.Atoi(val2)
				r.port = v2
			}

		case "password":
			if val,ok:=v.(string);ok{
				r.password = val
			}
		case "prefix":
			if val,ok:=v.(string);ok{
				r.prefix = val
			}

		case "db":
			if val,ok:=v.(int);ok{
				r.db = val
			} else if val2,ok2 := v.(string);ok2 {
				v2,_ := strconv.Atoi(val2)
				r.db = v2
			}

		case "desc":
			if val,ok:=v.(string);ok{
				r.desc = val
			}

		}
	}
	return r
}

func (r *redisinfo)ToString() string  {
	return fmt.Sprintf("host=%s,port=%d,password=%s,db=%d,prefix=%s,desc=%s",
						r.GetHost(),
						r.GetPort(),
						r.GetPassword(),
						r.GetDb(),
						r.GetPrefix(),
						r.GetDesc() )
}

func NewRedisInfo() redisinfo {
	return redisinfo{}
}


//redis组管理
type RedisGroupManage struct {
	redisGroup map[string]redisinfo
}

func (this *RedisGroupManage)Set(groupName string, cfg redisinfo)  {
	if this.redisGroup == nil {
		this.redisGroup = make(map[string]redisinfo)
	}
	this.redisGroup[groupName] = cfg
}

func (this *RedisGroupManage)SetMap(groupName string, info map[string]interface{})  {
	if this.redisGroup == nil {
		this.redisGroup = make(map[string]redisinfo)
	}
	redisinfoObj := NewRedisInfo()
	redisinfoObj.SetRedisInfo(info)
	this.redisGroup[groupName] = redisinfoObj
}

func (this RedisGroupManage)Get(groupName string) (redisinfo, error) {
	if ret, ok := this.redisGroup[groupName];ok {
		return ret, nil
	}
	return redisinfo{},errors.New(fmt.Sprintf("redis组%s未配置", groupName))
}
//获取当前redis组的配置key前缀
func (this RedisGroupManage)GetPrefix(groupName string) string  {
	if ret, ok := this.redisGroup[groupName];ok {
		return ret.GetPrefix()
	}
	return ""
}

//单例
var muRedisgroup_1 sync.Mutex
var redisGroupManage_1 *RedisGroupManage
var redisGroupManageInit_1 uint32
func NewRedisGroupManage() *RedisGroupManage  {
	if atomic.LoadUint32(&redisGroupManageInit_1) == 1 {
		return redisGroupManage_1
	}
	muRedisgroup_1.Lock()
	defer muRedisgroup_1.Unlock()
	if redisGroupManageInit_1 == 0 { //未初始化
		redisGroupManage_1 = new(RedisGroupManage)
		atomic.StoreUint32(&redisGroupManageInit_1, 1)
	}
	return redisGroupManage_1
}

