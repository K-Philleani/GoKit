package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// sql(mysql) MySQL事务
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

func transaction() {
	tx, err := DB.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "update user set age = 30 where id =?"
	ret1, err := tx.Exec(sqlStr1, 5)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr := "update user set age = 40 where id = ?"
	ret2, err := tx.Exec(sqlStr, 4)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}
	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}
	fmt.Println("exec trans success!")
}

func main() {
	initDB()
	transaction()
}
