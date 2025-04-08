package main

import (
	"database/sql"
	"fmt"
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
	dbFile string
}

type Task struct {
	id        int
	text      string
	completed bool
}

func (tsql TaskSQL) CreateTable() error {
	db, err := sql.Open("sqlite3", tsql.dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(CreateTable)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) PutTask(task string) error {
	db, err := sql.Open("sqlite3", tsql.dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(InsertTask, task)
	if err != nil {
		return err
	}
	return nil
}



func (tsql TaskSQL) DeleteTask() error {
	db, err := sql.Open("sqlite3", tsql.dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(DeleteWhereCompleted, true)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) UpdateTaskToCompleted(id int) error {
	db, err := sql.Open("sqlite3", tsql.dbFile)
	if err != nil {
		return err
	}
	defer db.Close()
	_, err = db.Exec(UpdateTaskCompletedWereID, true, id)
	if err != nil {
		return err
	}
	return nil
}


func (tsql TaskSQL) GetTask() ([]Task, error) {
	tasks := []Task{}
	db, err := sql.Open("sqlite3", tsql.dbFile)
	if err != nil {
		return nil,err
	}
	defer db.Close()
	rows, err := db.Query(SelectAllTasks)
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

func (tsql TaskSQL) printTask(tasks []Task) {
	for i, t := range tasks {
		fmt.Printf("\n%v ] %v {%v}\n", i, t.text, t.completed)
	}
}
