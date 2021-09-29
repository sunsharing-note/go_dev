package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)
func main() {
	jsonFile,err:=os.Open("user.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	defer jsonFile.Close()
	byteValue,_:=ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue,&users)
	for i :=0;i<len(users.Users);i++ {
		fmt.Println("User Type: "+ users.Users[i].Address)
		fmt.Println("User Age: "+strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: "+users.Users[i].Name)
		fmt.Println("User Email: "+users.Users[i].Social.Email)
	}
	var result Users
	json.Unmarshal(byteValue,&result)
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
