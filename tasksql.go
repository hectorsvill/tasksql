package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName = "tasks"
	CreateTable = "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, completed BOOLEAN DEFAULT FALSE);"
	DeleteWhereCompleted = "DELETE FROM tasks WHERE completed = ?;"
	UpdateTaskCompletedWereID = "UPDATE tasks SET completed = ? WHERE id = ?;"
	InsertTask = "INSERT INTO tasks (text) VALUES (?);"
	SelectAllTasks = "SELECT * FROM Tasks;"
)

type TaskSQL struct{
	db *sql.DB
}

type Task struct {
	id        int
	text      string
	completed bool
}

func (tsql TaskSQL) CreateTable() error {
	_, err := tsql.db.Exec(CreateTable)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) PostTask(task string) error {
	_, err := tsql.db.Exec(InsertTask, task)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) DeleteTask() error {
	_, err := tsql.db.Exec(DeleteWhereCompleted, true)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) UpdateTaskToCompleted(id int) error {
	_, err := tsql.db.Exec(UpdateTaskCompletedWereID, true, id)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) GetTask() ([]Task, error) {
	tasks := []Task{}
	rows, err := tsql.db.Query(SelectAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.id, &t.text, &t.completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

