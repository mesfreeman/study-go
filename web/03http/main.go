/*
 * @Author: Hequanxi
 * @Email: hequanxi@bianfeng.com
 * @Date: 2020-12-15 10:57:46
 * @Logs:
 */
package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // 解析参数，默认不解析
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key is ", k)
		fmt.Println("value is ", v)
	}
	fmt.Fprintf(w, "hello end.")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
