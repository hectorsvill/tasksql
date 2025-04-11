package tasksql

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName = "tasks"
	createTableIfNotExist = "CREATE TABLE IF NOT EXISTS {table} (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, deleted BOOLEAN DEFAULT FALSE);"
	deleteWhereDeletedTrue = "DELETE FROM {table} WHERE completed = ?;"
	updateDeletedTrueWereID = "UPDATE {table} SET completed = ? WHERE id = ?;"
	insertTasksValueText = "INSERT INTO {table} (text) VALUES (?);"
	selectAllTasks = "SELECT id,text,completed FROM {table};"
)

type TaskSQL struct{
	DB *sql.DB
}

func NewDB(dbSourceName string) (*TaskSQL, error) {
	db, err := sql.Open("sqlite3", dbSourceName)
	if err != nil {
		return  nil, err
	}
	return &TaskSQL{DB: db}, nil
}

func (tsql TaskSQL) CloseTaskSQl() error {
	if tsql.DB != nil {
		return tsql.DB.Close()
	}
	return nil
}

func (tsql TaskSQL) CreateTableIfNotExist(table string) error {
	createTableIfNotExistWithTable := replaceTableName(createTableIfNotExist, table)
	log.Println(createTableIfNotExistWithTable)
	_, err := tsql.DB.Exec(createTableIfNotExistWithTable)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) PostTask(table, task string) error {
	insert := replaceTableName(insertTasksValueText, table)
	_, err := tsql.DB.Exec(insert, task)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) DeleteWhereDeletedTrue(table string) error {
	delete := replaceTableName(deleteWhereDeletedTrue, table)
	_, err := tsql.DB.Exec(delete, true)
	if err != nil {
		return err
	}
	return nil
}

func (tsql TaskSQL) UpdateTaskToDelete(table string, id int) error {
	
	_, err := tsql.DB.Exec(updateDeletedTrueWereID, true, id)
	if err != nil {
		return err
	}
	return nil
}

type Task struct {
	Id        int
	Text      string
	Completed bool
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

func replaceTableName(query ,tableName string) string {
	return 	strings.Replace(query, "{table}", tableName, -1)

}