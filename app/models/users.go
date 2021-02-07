package models

import (
	"log"
	"time"
)

//User ユーザーの構造体
type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	PassWord  string
	CreatedAt time.Time
}

//CreateUser ユーザーを作るためのメソッド
func (u *User) CreateUser() (err error) {
	cmd := `insert into users(
			uuid,
			name,
			email,
			password,
			created_at) values (?,?,?,?,?)`

	_, err = Db.Exec(cmd,
		createUUID(),
		u.Name,
		u.Email,
		Encrypt(u.PassWord),
		time.Now())

	if err != nil {
		log.Fatal(err)
	}
	return err
}

//GetUser IDを指定して、Userを得る関数
func GetUser(id int) (user User, err error) {
	user = User{}
	cmd := `SELECT id, uuid, name, email, password, created_at
	FROM users where id = ?`
	err = Db.QueryRow(cmd, id).Scan(
		&user.ID,
		&user.UUID,
		&user.Name,
		&user.Email,
		&user.PassWord,
		&user.CreatedAt,
	)

	return user, err

}

//UpdateUser Userのポインタを渡して、User情報をアップデートする
func (u *User) UpdateUser() (err error) {
	cmd := `UPDATE users set name = ?, email = ? where id = ?`

	Db.Exec(cmd, u.Name, u.Email, u.ID)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

//DeleteUser Userのポインタを渡して、User情報を削除する
func (u *User) DeleteUser() (err error) {
	cmd := `DELETE FROM users where id = ?`
	_, err = Db.Exec(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
