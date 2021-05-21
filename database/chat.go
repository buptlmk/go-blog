package database

import (
	"blog/database/mysql"
	"blog/log"
	"blog/utils"
)

func GetAllChatRoom() ([]utils.Chat, error) {

	tx, err := mysql.Begin()
	if err != nil {
		return nil, err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `select id,name,person_number from chat_room`

	rows, err := tx.Query(sqlCmd)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}

	rooms := make([]utils.Chat, 0, 16)
	for rows.Next() {
		room := utils.Chat{}
		err = rows.Scan(&room.Id, &room.Name, &room.PersonNumber)
		if err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}
