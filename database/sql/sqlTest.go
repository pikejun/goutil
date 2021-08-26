package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type data struct {
	city_code    string
	content_type string
	sub_type     string
}

func main() {
	dsn := "my:My@123456@tcp(127.0.0.1:3306)/my_db"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	d := new(data)
	sqlStr := "select city_code,content_type,sub_type from data_raw_content where id =? limit 1"
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	db.QueryRow(sqlStr, 1).Scan(&d.city_code, &d.content_type, &d.sub_type)

	fmt.Printf("city_code:%s content_type:%s sub_type:%s\n", d.city_code, d.content_type, d.sub_type)
}
