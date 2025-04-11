package tasksql_test

import (
	"log"
	"testing"

	"github.com/hectorsvill/tasksql"

)

func Test_CreateTableIfNotExist(t *testing.T) {
	tasksql, err := tasksql.NewDB("test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer tasksql.CloseTaskSQl()

	table_name := "data"
	err = tasksql.CreateTableIfNotExist(table_name)
	if err != nil {
		t.Fatalf("[Test_CreateTableIfNotExist] %s", "err")	
	}
	
	err = tasksql.PostTask("hello")
	if err != nil {
		t.Fatal(err)
	}
	
}


