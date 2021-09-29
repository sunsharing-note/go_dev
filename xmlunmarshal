package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)
func main() {

	xmlFile,err:=os.Open("users.xml")
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println("successfully opened users.xml")
	defer xmlFile.Close()
	byteValue,_:=ioutil.ReadAll(xmlFile)
	var users Users
	xml.Unmarshal(byteValue,&users)
	for i :=0;i<len(users.Users);i++ {
		fmt.Println("User Address: "+users.Users[i].Address)
		fmt.Println("User Name: "+users.Users[i].Name)
		fmt.Println("Facebook Url: "+users.Users[i].Social.Email)
	}
}
type Users struct {
	XMLName xml.Name `xml:"users"`
	Users []User `xml:"user"`
}

type User struct {
	XMLName xml.Name `xml:"user"`
	Address string `xml:"address,attr"`
	Name string `xml:"name"`
	Social Social `xml:"social"`
}
type Social struct {
	XMLName xml.Name `xml:"social"`
	Mobile string `xml:"mobile"`
	Email string `xml:"email"`
}
