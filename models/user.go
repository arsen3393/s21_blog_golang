package models

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdateAt  time.Time `db:"updated_at" json:"updated_at"`
}

type UserModel struct {
	DB *sqlx.DB
}

func NewUserModel(db *sqlx.DB) *UserModel {
	return &UserModel{DB: db}
}

func (um *UserModel) GetUserById(id int) (*User, error) {
	var user User
	err := um.DB.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (um *UserModel) GetUserByName(name string) (*User, error) {
	var user User
	err := um.DB.Get(&user, "SELECT * FROM users WHERE name = $1", name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//func (um *UserModel) Authenticate(name, password string) (*User, error) {
//	row := um.DB.QueryRow("SELECT * FROM users WHERE name = $1", name)
//	user := &User{}
//	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return nil, errors.New("user not found")
//		}
//		return nil, err
//	}
//}
