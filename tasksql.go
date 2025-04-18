package tasksql

import (
	"database/sql"
	"errors"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	tableName               = "tasks"
	createTableIfNotExist   = "CREATE TABLE IF NOT EXISTS {table} (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, deleted BOOLEAN DEFAULT FALSE);"
	deleteWhereDeletedTrue  = "DELETE FROM {table} WHERE deleted = ?;"
	updateDeletedTrueWereID = "UPDATE {table} SET deleted = ? WHERE id = ?;"
	insertTasksValueText    = "INSERT INTO {table} (text) VALUES (?);"
	selectAllText           = "SELECT text FROM {table} ORDER BY id;"
)

type TaskSQL struct {
	DB *sql.DB
}

func NewDB(dbSourceName string) (*TaskSQL, error) {
	db, err := sql.Open("sqlite3", dbSourceName)
	if err != nil {
		return nil, err
	}
	return &TaskSQL{DB: db}, nil
}

func (tsql *TaskSQL) CloseTaskSQl() error {
	if tsql.DB != nil {
		return tsql.DB.Close()
	}
	return nil
}

func (tsql *TaskSQL) CreateTableIfNotExist(table string) error {
	createTableIfNotExistWithTable, err := replaceTableName(createTableIfNotExist, table)
	if err != nil {
		return err
	}
	// log.Println(createTableIfNotExistWithTable)
	_, err = tsql.DB.Exec(createTableIfNotExistWithTable)
	if err != nil {
		return err
	}
	return nil
}

func (tsql *TaskSQL) Post(table, text string) error {
	insert, err := replaceTableName(insertTasksValueText, table)
	if err != nil {
		return err
	}
	_, err = tsql.DB.Exec(insert, text)
	if err != nil {
		return err
	}
	return nil
}

func (tsql *TaskSQL) UpdateToDelete(table string, id int) error {
	query, err := replaceTableName(updateDeletedTrueWereID, table)
	if err != nil {
		return err
	}
	_, err = tsql.DB.Exec(query, true, id)
	if err != nil {
		return err
	}
	return nil
}

func (tsql *TaskSQL) DeleteWhereDeletedTrue(table string) error {
	query, err := replaceTableName(deleteWhereDeletedTrue, table)
	if err != nil {
		return err
	}

	delete, err := replaceTableName(query, table)
	if err != nil {
		return err
	}
	_, err = tsql.DB.Exec(delete, true)
	if err != nil {
		return err
	}
	return nil
}

func (tsql *TaskSQL) Get(table string) ([]string, error) {
	data := []string{}
	selectAllTextWithTable, err := replaceTableName(selectAllText, table)
	if err != nil {
		return data, err
	}
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

func replaceTableName(query, tableName string) (string, error) {
	if !IsValidTableID(tableName) {
		return "", errors.New("[replaceTableName]: invalid table name")
	}
	return strings.Replace(query, "{table}", tableName, -1), nil
}

// TODO: move to tools folder
func IsValidTableID(tableName string) bool {
	for _, r := range tableName {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' || r == '_') {
			return false
		}
	}
	return true
}
