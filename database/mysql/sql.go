package mysql

import (
	"blog/config"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

var DB *sql.DB

func InitDB() (err error) {
	//config.Load("config.json")
	SQLConfig := config.Settings.SQLSet
	//fmt.Println(SQLConfig)
	path := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", SQLConfig.User, SQLConfig.Password, SQLConfig.Ip, strconv.Itoa(SQLConfig.Port), SQLConfig.DBName)
	//path := SQLConfig.User + ":" + SQLConfig.Password + "@tcp(" + SQLConfig.Ip + ":" + SQLConfig.Port + ")/" + SQLConfig.DBName + "?charset=utf8"
	//log.Logger.Info(path)
	// 这里必须用=，不然会变为局部变量，隐藏的bug
	DB, err = sql.Open("mysql", path)

	if err != nil {
		return err
	}

	DB.SetConnMaxLifetime(time.Minute)
	DB.SetMaxOpenConns(200)
	DB.SetMaxIdleConns(50)
	//log.Logger.Info(path)
	return DB.Ping()
}
func Close() error {
	err := DB.Close()
	return err
}
func Begin() (*sql.Tx, error) {
	return DB.Begin()
}

func GraceCommit(course *sql.Tx, err error) error {
	if err != nil {
		err = course.Rollback()
		return err
	}
	err = course.Commit()
	return err
}

func testDB() {
	_ = InitDB()
	tx, err := Begin()
	if err != nil {
		fmt.Println("0----")
		panic(err)
	}
	var insert = `update data.user set password='23' where password=? limit 1;`
	//name := "lmk"
	//psd := "6542"
	result, err := tx.Exec(insert, "123")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(result.RowsAffected())
	_ = GraceCommit(tx, err)
}
