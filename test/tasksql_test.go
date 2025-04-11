package tasksql_test

import (
	"testing"

	"github.com/hectorsvill/tasksql"

)

func Test_CreateTableIfNotExist(t *testing.T) {
	tasksql, err := tasksql.NewDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer tasksql.CloseTaskSQl()

	tableName := "data"
	err = tasksql.CreateTableIfNotExist(tableName)
	if err != nil {
		t.Fatalf("[Test_CreateTableIfNotExist] %s", "err")	
	}
	
	err = tasksql.PostTask(tableName,"hello")
	if err != nil {
		t.Fatal(err)
	}
	

	
}
