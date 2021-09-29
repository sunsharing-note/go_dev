package main

import (
	"encoding/xml"
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
    v:=&Users{Users: []User{wanger,huazi,qiaoke,xialaoshi,zhengge,jiangzong,dongdong}}

	result, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(string(result))
	fileName := "users.xml"
	err = ioutil.WriteFile(fileName, result, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}
}
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users []User `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Age int64 `xml:"age"`
	Address string `xml:"address,attr"`
	Name string `xml:"name"`
	Social Social `xml:"social"`
}
type Social struct {
	XMLName xml.Name `xml:"social"`
	Mobile string `xml:"mobile"`
	Email string `xml:"email"`
}
