package database

import (
	"blog/database/mysql"
	"blog/log"
	"blog/utils"
	"time"
)

func GetArticle(id int64) (article utils.Article, err error) {

	sqlCmd := `select a.id,a.author_id,b.name,a.title,a.content,a.visited_number,a.star_number,a.like_number,a.created_time from article as a, user as b where a.id = ?`

	tx, err := mysql.Begin()
	if err != nil {
		return article, err
	}
	defer mysql.GraceCommit(tx, err)

	row := tx.QueryRow(sqlCmd, id)
	err = row.Scan(&article.Id, &article.AuthorId, &article.Author, &article.Title, &article.Content, &article.VisitedNumber, &article.StarNumber, &article.LikeNumber, &article.CreatedTime)
	if err != nil {
		log.Logger.Error(err.Error())
		return article, err
	}

	return article, nil

}

func StarArticle(id int64) (err error) {
	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)
	sqlCmd := `update article set star_number = star_number+1 where id=?`

	_, err = tx.Exec(sqlCmd, id)
	return err
}

func GetRecentArticle(limit int, offset int) (articles []utils.Article, err error) {
	tx, err := mysql.Begin()
	if err != nil {
		return nil, err
	}
	defer mysql.GraceCommit(tx, err)
	sqlCmd := `select a.id,a.author_id,b.name,a.title,a.content,a.visited_number,a.star_number,a.like_number,a.created_time,
				a.comment_number from article as a, user as b order by a.created_time desc limit ? offset ?`

	rows, err := tx.Query(sqlCmd, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		article := utils.Article{}
		err = rows.Scan(&article.Id, &article.AuthorId, &article.Author, &article.Title, &article.Content, &article.VisitedNumber,
			&article.StarNumber, &article.LikeNumber, &article.CreatedTime, &article.CommentNumber)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil

}

func AddArticleVisitedNumber(id int64, number int64) (err error) {

	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)
	sqlCmd := `update article set visited_number = visited_number + ? where id=?`

	_, err = tx.Exec(sqlCmd, number, id)

	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	return

}

func AddArticle(authorId int, title string, content string) (err error) {
	tx, err := mysql.Begin()
	if err != nil {
		return err
	}
	defer mysql.GraceCommit(tx, err)

	sqlCmd := `insert into article(title,author_id,content,created_time) values (?,?,?,?)`

	_, err = tx.Exec(sqlCmd, title, authorId, content, time.Now())
	return err

}
