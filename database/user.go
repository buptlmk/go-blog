package database

import (
	"blog/database/mysql"
	"blog/log"
	"blog/utils"
)

func GetUserPassword(cardId string) (id int64, passwordHash, saltHash string, name string, err error) {

	db, err := mysql.Begin()
	if err != nil {
		log.Logger.Error(err.Error())
		return 0, "", "", "", err
	}
	defer func() {
		_ = mysql.GraceCommit(db, err)
	}()

	sqlCmd := `select id,password,password_salt,name from data.user where card_id = ?`

	row := db.QueryRow(sqlCmd, cardId)
	err = row.Scan(&id, &passwordHash, &saltHash, &name)

	return

}

func GetUserInfo(cardId string) (user utils.User, err error) {

	db, err := mysql.Begin()
	if err != nil {
		log.Logger.Error(err.Error())
		return user, err
	}
	defer func() {
		_ = mysql.GraceCommit(db, err)
	}()

	sqlCmd := `select id,name,phone,email,github,qq,wechat from data.user where card_id = ?`

	row := db.QueryRow(sqlCmd, cardId)
	err = row.Scan(&user.Id, &user.Name, &user.Phone, &user.Email, &user.Github, &user.QQ, &user.Wechat)

	return user, err

}

func RegisterUser(cardID, name, passwordSalt, salt, phone, email string) (err error) {

	db, err := mysql.Begin()
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	defer func() {
		_ = mysql.GraceCommit(db, err)
	}()

	sqlCmd := `insert into data.user(card_id,name,password,password_salt,phone,email) VALUES (?,?,?,?,?,?);`

	_, err = db.Exec(sqlCmd, cardID, name, passwordSalt, salt, phone, email)

	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}

	return nil
}
