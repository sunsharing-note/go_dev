package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	wanger:=User{Address:"beijing",Name:"wanger",Age:24,Social:Social{Email:"wanger@163.com",Mobile:"111111111111"}}
	huazi:=User{Address:"shenzhen",Name:"huazai",Age:28,Social:Social{Email:"huazai@163.com",Mobile:"111122211111"}}
	qiaoke:=User{Address:"chongqing",Name:"qiaoke",Age:30,Social:Social{Email:"qiaoke@163.com",Mobile:"13332211111"}}
	xialaoshi:=User{Address:"chengdu",Name:"夏老师",Age:29,Social:Social{Email:"xialaoshi@163.com",Mobile:"11144445411111"}}
	jiangzong:=User{Address:"shanghai",Name:"姜总",Age:25,Social:Social{Email:"jiangzong@163.com",Mobile:"111222445211111"}}
	dongdong:=User{Address:"chengdu",Name:"冬哥",Age:30,Social:Social{Email:"dongdong@163.com",Mobile:"1155555211111"}}
	zhengge:=User{Address:"beijing",Name:"郑哥",Age:24,Social:Social{Email:"zhengge@163.com",Mobile:"1112224566211111"}}
	result:=Users{Users: []User{wanger,huazi,jiangzong,xialaoshi,qiaoke,dongdong,zhengge}}
	bytearray,err:=json.Marshal(result)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytearray))
	fileName := "user.json"
	err = ioutil.WriteFile(fileName, bytearray, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
}
type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"Age"`
	Social Social `json:"social"`
}
type Social struct {
	Mobile string `json:"mobile"`
	Email string `json:"email"`
}
