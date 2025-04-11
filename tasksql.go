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
	updateDeletedTrueWereID = "UPDATE {table} SET deleted = ? WHERE id = ?;"
	insertTasksValueText = "INSERT INTO {table} (text) VALUES (?);"
	selectAllText = "SELECT text FROM {table} ORDER BY id LIMIT 10;"
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

func (tsql TaskSQL) PostTask(table, text string) error {
	insert := replaceTableName(insertTasksValueText, table)
	_, err := tsql.DB.Exec(insert, text)
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

func (tsql TaskSQL) GetTask(table string) ([]string, error) {
	data := []string{}
	selectAllTextWithTable := replaceTableName(selectAllText, table)
	rows, err := tsql.DB.Query(selectAllTextWithTable)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var text string
		err = rows.Scan(&text)
		if err != nil {
			return nil, err
		}
		data = append(data, text)
	}
	return data, nil
}

func replaceTableName(query ,tableName string) string {
	return 	strings.Replace(query, "{table}", tableName, -1)

}