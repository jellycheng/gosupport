package dbutils

import (
	"fmt"
)

type MysqlDsn struct {
	host       string
	port       string
	username   string
	password   string
	dbname     string   //库名
	charset    string  //数据库编码
	extparam   string  //扩展参数
}

func (m *MysqlDsn) ToDsn() string {
	var param string  //parseTime=True&loc=Local
	if m.extparam != "" {
		param = fmt.Sprintf("&%s", m.extparam)
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8%s",
						m.username, m.password, m.host, m.port, m.dbname, param)
}

func (m *MysqlDsn) Key() string {
	return fmt.Sprintf("%s:%s:%s", m.host, m.port, m.dbname)
}

func (m *MysqlDsn) GetHost() string  {
	return m.host
}

func (m *MysqlDsn) GetPort() string  {
	return m.port
}

func (m *MysqlDsn) GetUsername() string  {
	return m.username
}

func (m *MysqlDsn) GetPassword() string  {
	return m.password
}

func (m *MysqlDsn) GetDbname() string  {
	return m.dbname
}

func (m *MysqlDsn) GetCharset() string  {
	return m.charset
}

func (m *MysqlDsn) GetEextparam() string  {
	return m.extparam
}

//mysqlDsnObj := dbutils.NewMysqlDsn(map[string]interface{}{"user_name":"root","password":"123456","port":3307})
func NewMysqlDsn(dbConfig map[string]interface{}) *MysqlDsn {
	dsnObj := &MysqlDsn{}

	if h, ok := dbConfig["host"];ok {
		if hStr,ok := h.(string);ok {
			dsnObj.host = hStr
		} else {
			dsnObj.host = "localhost"
		}
	} else {
		dsnObj.host = "localhost"
	}

	if p, ok := dbConfig["port"];ok {
		if pStr,ok := p.(string);ok {
			dsnObj.port = pStr
		} else if p2Str,ok:=p.(int);ok{
			dsnObj.port = fmt.Sprintf("%d",p2Str)
		} else {
			dsnObj.port = "3306"
		}
	} else {
		dsnObj.port = "3306"
	}

	if u, ok := dbConfig["username"];ok {
		if uStr,ok := u.(string);ok {
			dsnObj.username = uStr
		} else {
			dsnObj.username = "root"
		}
	} else {
		dsnObj.username = "root"
	}

	if p, ok := dbConfig["password"];ok {
		if pStr,ok := p.(string);ok {
			dsnObj.password = pStr
		} else {
			dsnObj.password = ""
		}
	} else {
		dsnObj.password = ""
	}

	if d, ok := dbConfig["dbname"];ok {
		if dStr,ok := d.(string);ok {
			dsnObj.dbname = dStr
		} else {
			dsnObj.dbname = ""
		}
	} else {
		dsnObj.dbname = ""
	}

	if c, ok := dbConfig["charset"];ok {
		if cStr,ok := c.(string);ok {
			dsnObj.charset = cStr
		} else {
			dsnObj.charset = "utf8"
		}
	} else {
		dsnObj.charset = "utf8"
	}

	if param, ok := dbConfig["extparam"];ok && param != "" {
		dsnObj.extparam = fmt.Sprintf("%s", param)
	} else {
		dsnObj.extparam = "";
	}

	return dsnObj
}

