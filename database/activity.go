package database

import (
	"blog/database/mysql"
	"blog/log"
	"blog/utils"
)

func GetAllActivity() ([]utils.Activity, error) {

	tx, err := mysql.Begin()
	if err != nil {
		return nil, err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `select id,name,description,time,img,total,res from activity`

	rows, err := tx.Query(sqlCmd)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}

	activities := make([]utils.Activity, 0, 16)
	for rows.Next() {
		activity := utils.Activity{}
		err = rows.Scan(&activity.Id, &activity.Name, &activity.Description, &activity.Time, &activity.Img, &activity.Total, &activity.Res)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func JoinActivity(activityId int) (err error) {
	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `update activity set res = res-1 where id=?`
	_, err = tx.Exec(sqlCmd, activityId)
	return err
}

func AddActivity(name string, description string, img string, time string, total, res int) (id int64, err error) {
	tx, err := mysql.Begin()
	if err != nil {
		return 0, err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `insert into activity(name, description, time, img, total, res) values(?,?,?,?,?,?) `
	result, err := tx.Exec(sqlCmd, name, description, time, img, total, res)
	if err != nil {
		return 0, err
	}
	id, err = result.LastInsertId()
	return id, err
}
