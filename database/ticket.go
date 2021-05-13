package database

import (
	"blog/database/mysql"
	"blog/log"
	"time"
)

func SetTicket(name string, totalNumber int64) (err error) {

	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `insert into ticket(name,total_number,res_number,created_time) value(?,?,?,?)`

	if _, err = tx.Exec(sqlCmd, name, totalNumber, totalNumber, time.Now()); err != nil {

		log.Logger.Error(err.Error())
		return err
	}
	return nil
}

func SetResTicket(id int64, number int64) (err error) {

	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)
	sqlCmd := `update ticket set res_number = ? where id=?`

	_, err = tx.Exec(sqlCmd, number, id)

	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	return

}
