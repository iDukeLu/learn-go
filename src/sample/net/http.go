package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		method := r.Method
		name := r.FormValue("name")

		log.Println("Starting v1 server ...")
		w.Write([]byte("httpserver v1"))
		log.Printf("path: /, method: %v, param: %v", method, name)
		_, _ = w.Write([]byte("httpserver v1"))
	})
	http.HandleFunc("/bye", sayBye)
	http.HandleFunc("/post", post)
	log.Println("Starting v1 server ...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye bye ,this is v1 httpServer"))
}

func post(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength    // 获取请求实体长度
	body := make([]byte, len) // 创建存放请求实体的字节切片
	r.Body.Read(body)         // 调用 Read 方法读取请求实体并将返回内容存放到上面创建的字节切片
	// io.WriteString(w, string(body))
	post := Post{}
	json.Unmarshal(body, &post)   // 对读取的 JSON 数据进行解析
	fmt.Fprintf(w, "%#v\n", post) // 格式化输出结果
}

type Post struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}
