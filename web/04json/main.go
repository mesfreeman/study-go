package main

import (
	"encoding/json"
	"fmt"
)

// Server 服务对象，定义服务名称及IP
type Server struct {
	ServerName string
	serverIP   string
}

// ServerList 服务对象列表
type ServerList struct {
	Servers []Server
}

func main() {
	jsonStr := `{"servers": [{"serverName": "s1", "serverIP": "192.168.1.1"}, {"serverName": "s2", "serverIP": "192.168.1.2"}]}`

	var s ServerList
	err := json.Unmarshal([]byte(jsonStr), &s)
	if err != nil {
		fmt.Println("json convert err", err)
	}
	fmt.Println(s)
	fmt.Printf("s type is %T \n", s)
	fmt.Printf("s type is %T \n", s.Servers)

	for k, v := range s.Servers {
		fmt.Println(k, v.ServerName)
	}
}
