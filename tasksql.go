package tasksql

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName = "tasks"
	createTableIfNotExist = "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, completed BOOLEAN DEFAULT FALSE);"
	deleteWhereTaskCompleted = "DELETE FROM tasks WHERE completed = ?;"
	updateTaskCompletedWereID = "UPDATE tasks SET completed = ? WHERE id = ?;"
	insertTasksValueText = "INSERT INTO tasks (text) VALUES (?);"
	selectAllTasks = "SELECT id,text,completed FROM tasks;"
)

type TaskSQL struct{
	DB *sql.DB
}

type Task struct {
	Id        int
	Text      string
	Completed bool
}

func (tsql TaskSQL) CreateTableIfNotExist() error {
	_, err := tsql.DB.Exec(createTableIfNotExist)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) PostTask(task string) error {
	_, err := tsql.DB.Exec(insertTasksValueText, task)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) DeleteWhereTaskCompleted() error {
	_, err := tsql.DB.Exec(deleteWhereTaskCompleted, true)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) UpdateTaskToCompleted(id int) error {
	_, err := tsql.DB.Exec(updateTaskCompletedWereID, true, id)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) GetTask() ([]Task, error) {
	tasks := []Task{}
	rows, err := tsql.DB.Query(selectAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var t Task
		err = rows.Scan(&t.Id, &t.Text, &t.Completed)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
