package models

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type Post struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Text      string    `json:"text" db:"text"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdateAt  time.Time `db:"updated_at" json:"updated_at"`
}

type PostModel struct {
	DB *sqlx.DB
}

func NewPostModel(db *sqlx.DB) *PostModel {
	return &PostModel{DB: db}
}

func (pm *PostModel) InsertNewPost(title, text string) error {
	const op = "PostModel.Insert"
	query := "INSERT INTO post(title, text, created_at, updated_at) VALUES ($1, $2, time.Now(), time.Now())"
	_, err := pm.DB.Exec(query, title, text)
	if err != nil {
		return errors.New(op + ": " + err.Error())
	}
	return nil
}

func (pm *PostModel) GetAllPosts(page int) ([]Post, int, error) {
	const op = "PostModel.GetAllPosts"
	perPage := 3
	offset := (page - 1) * perPage
	query := "SELECT title, text FROM posts ORDER BY created_at LIMIT $1 OFFSET $2"
	queryTotal := "SELECT COUNT(*) FROM posts"
	rows, err := pm.DB.Query(query, perPage, offset)
	if err != nil {
		return nil, 0, errors.New(op + ": " + err.Error())
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println(err.Error())
		}
	}()
	var total int
	err = pm.DB.QueryRow(queryTotal).Scan(&total)
	if err != nil {
		return nil, 0, errors.New(op + ": " + err.Error())
	}
	var posts []Post
	for rows.Next() {
		var post Post
		if err := rows.Scan(&post.Title, &post.Text); err != nil {
			return nil, 0, errors.New(op + ": " + err.Error())
		}
		posts = append(posts, post)
	}
	return posts, total, nil
}

func (pm *PostModel) CreatePost(title, text string) error {
	const op = "PostModel.CreatePost"
	queryToRestoreIncrement := "SELECT setval('posts_id_seq', (SELECT MAX(id) FROM posts))"
	tx, err := pm.DB.Begin()
	if err != nil {
		return errors.New(op + " failed to begin transaction: " + err.Error())
	}
	defer tx.Rollback()
	_, err = tx.Exec(queryToRestoreIncrement)
	if err != nil {
		return errors.New(op + ": failed to restore increment" + err.Error())
	}
	query := "INSERT INTO posts(title, text) VALUES ($1, $2)"
	_, err = pm.DB.Exec(query, title, text)
	if err != nil {
		return errors.New(op + ": " + err.Error())
	}
	err = tx.Commit()
	if err != nil {
		return errors.New(op + " failed to begin transaction: " + err.Error())
	}
	return nil
}

//func (um *UserModel) GetUserById(id int) (*User, error) {
//	var user User
//	err := um.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)
//	if err != nil {
//		return nil, err
//	}
//	return &user, nil
//}
