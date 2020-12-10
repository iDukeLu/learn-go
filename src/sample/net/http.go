package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var (
	m = map[string]func(http.ResponseWriter, *http.Request){
		"/bye": sayBye,
		"post": post,
	}
)

type User struct {
	Name string
}

type UserService interface {
	getName() string
}

func main() {
	userService := User{"11111"}
	userService.getName()
	fmt.Print(userService.getName())
}

func (user User) getName() string {
	return user.Name
}

func httpDemo() {
	//for paths, handler := range m {
	//	for path := range paths {
	//		http.HandleFunc(path, handler)
	//	}
	//
	//}
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

func a() {
	adder(Post{"", ""})
	b()
}

func b() func() {
	return func() {

	}
}

func adder(p Post) func(int) int {
	sum := 0

	g := GetMapping{}
	method := g.method
	return func(x int) int {
		sum += x
		return sum
	}
}

type RequestMapping struct {
	path   []string
	method string
}

type GetMapping struct {
	RequestMapping
}
