package main
import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

// sql(mysql) MySQL预处理

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

// 预处理查询示例
func prepareQuery() {
	sqlStr := "select id, name, age from user where id > ?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入示例
func perpareInsert() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("小王子", 23)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("晚熟的人", 18)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	fmt.Println("insert success.")

}

func main() {
	initDB()
	prepareQuery()
	//perpareInsert()
}
