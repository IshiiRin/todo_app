package models

import (
	"log"
	"time"
)

//Todo Todoの構造体
type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

//CreateTodo Userに対して、ToDoを一つ作成するメソッド
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (content,
			user_id,
			created_at) values(?,?,?)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//GetTodo IDが一致するTodoを取得する関数
func GetTodo(id int) (todo Todo, err error) {
	cmd := `SELECT id,content,user_id,created_at from todos
	where id =?`
	todo = Todo{}

	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt,
	)
	return todo, err
}

//GetTodos 全てのTodosを取得する関数
func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id,content,user_id,created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//GetTodosByUser Userの全てのtodoを抽出する
func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id,content,user_id,created_at FROM todos
				WHERE user_id = ?`
	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()
	return todos, err
}

//UpdateTodo todoをアップデートする
func (t *Todo) UpdateTodo() error {
	cmd := `UPDATE todos set content = ?, user_id = ?
	WHERE id = ?`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//DeleteTodo Todoを削除する
func (t *Todo) DeleteTodo() error {
	cmd := `DELETE FROM todos WHERE id = ?`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
