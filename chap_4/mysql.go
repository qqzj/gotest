package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "用户名:密码@/test")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	rows, _ := db.Query("SELECT NOW() as time;")
	column, _ := rows.Columns()
	fmt.Println(column)
	var time string
	for rows.Next() {
		rows.Scan(&time)
		fmt.Println(time)
	}
	defer rows.Close()
	stmt, _ := db.Prepare("INSERT INTO `test`.`ay_app` (`uid`, `caid`, `app_name`, `app_description`, `app_key`, `app_platform`, `app_language`, `app_country`, `status`, `create_time`, `update_time`, `delete_time`) VALUES ('1', '1', ?, '一款休闲游戏', '10000004', 'Android', 'zh-cn', 'China', '1', '1438308391', '1438308391', '0');")
	for i := 0; i < 100; i++ {
		stmt.Exec("hello")
	}
	defer stmt.Close()
}
