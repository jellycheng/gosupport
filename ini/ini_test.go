package ini

import (
	"fmt"
	"testing"
)

func TestNewIniConfig(t *testing.T) {
	file := "./cjs.ini"
	obj := NewIniConfig(file)
	fmt.Println(obj.GetIniFile())
	if err := obj.ParseIniFile();err == nil {
		if val,err := obj.GetValue("author", "NAME");err == nil {
			fmt.Println(val)
		} else {
			fmt.Println(err.Error())
		}

		fmt.Println(obj.MustGetValue("author", "E-MAIL"))
		obj.DelValue("author", "E-MAIL")
		fmt.Println(obj.GetConfigData())

	}



}

func TestCleanComment(t *testing.T) {
	if v, ok := GetComment([]byte("hello#2;world"));ok {
		fmt.Println(string(v))
	} else {
		fmt.Println("没有注释")
	}

	fmt.Println(string(GetCleanComment([]byte("how"))))

	fmt.Println(string(GetCleanComment([]byte("good#a yes "))))

}
