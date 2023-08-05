package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type UserInfo struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
	sync.RWMutex
}

func main() {
	var ui UserInfo
	ui.SetUserInfo("Lory", 37)
	userdata, _ := json.MarshalIndent(&ui, "", " ")
	fmt.Printf("User info [%+v]\n", string(userdata))
}

func (ui *UserInfo) SetUserInfo(strName string, nAge int) {
	ui.Lock()
	defer ui.Unlock()
	ui.Age = nAge
	ui.Name = strName
}
