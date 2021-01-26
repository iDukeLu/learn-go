package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func main() {
	InitDB()
}

//数据库配置
const (
	userName = "fangbao_adb_prod"
	password = "66f3utI1PpiQp1D2"
	ip       = "81.71.16.200"
	port     = "33061"
	dbName   = "yx_mc"
)

//注意方法名大写，就是public
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, err := sql.Open("mysql", path)
	if err != nil {
		fmt.Println("opon database fail" + err.Error())
		return
	}

	if err = db.Ping(); err != nil {
		fmt.Println("opon database fail" + err.Error())
		return
	}
	fmt.Println("connnect success")
	QueryOne(db)
}

//查询单行
func QueryOne(DB *sql.DB) {
	defer DB.Close()
	user := new(Visitor) //用new()函数初始化一个结构体对象
	row := DB.QueryRow("select id,open_id,nickname from im_visitor where id=?", 597923011)
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.Id, &user.OpenId, &user.Nickname); err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Println("Single row data:", *user)
}

type Visitor struct {
	Id       string
	OpenId   string
	Nickname string
}
