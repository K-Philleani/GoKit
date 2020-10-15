package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// sql(mysql) CRUD

var DB *sql.DB
const DSN = "root:123456@tcp(124.70.71.78:3306)/sql_test?charset=utf8mb4&parseTime=True"

type user struct {
	id int
	age int
	name string
}

// 连接数据库库
func initDB() (err error) {
	DB, err = sql.Open("mysql", DSN)
	if err != nil {
		return
	}
	err = DB.Ping()
	if err != nil {
		return
	}
	return nil
}

// 插入数据
func insertRow() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := DB.Exec(sqlStr, "biubiu", 22)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", id)

}

// 单行查询数据
func queryRow() {
	sqlStr := "select id, name, age from user where id = ?"
	var u user
	err := DB.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多行查询
func queryMultiRow() {
	sqlStr := "select id, name, age from user where id > ?"
	rows, err := DB.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 更新数据
func updateRow() {
	sqlStr := "update user set age = ? where id = ?"
	ret, err := DB.Exec(sqlStr, 30, 2)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRow() {
	sqlStr := "delete from user where id = ?"
	ret, err := DB.Exec(sqlStr, 1)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed, err: %v\n", err)
		return
	}
	log.Println("数据库连接成功")
	//insertRow()
	//queryRow()
	//queryMultiRow()
	//updateRow()
	//deleteRow()
}
