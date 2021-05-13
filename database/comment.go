package database

import (
	"blog/database/mysql"
	"blog/log"
	"blog/utils"
	"time"
)

func AddComment(userId int64, articleId int64, content string) (err error) {

	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `insert into comment(content,article_id,user_id,created_time) value(?,?,?,?)`

	if _, err = tx.Exec(sqlCmd, content, articleId, userId, time.Now()); err != nil {

		log.Logger.Error(err.Error())
		return err
	}
	return nil
}

func GetComments(articleId int64) (comments []utils.Comment, err error) {

	tx, err := mysql.Begin()
	if err != nil {
		return nil, err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `select a.user_id,b.name,a.article_id,a.content,a.created_time from comment as a, user as b where a.article_id=? and a.user_id=b.id`

	rows, err := tx.Query(sqlCmd, articleId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		comment := utils.Comment{}
		err = rows.Scan(&comment.UserId, &comment.User, &comment.ArticleId, &comment.Content, &comment.CreatedTime)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
