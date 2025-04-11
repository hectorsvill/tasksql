package tasksql_test

import (
	"fmt"
	"log"
	"math/rand"
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
	
	for range 3 {
		err = tasksql.PostTask(tableName,fmt.Sprintf("text%v", rand.Intn(10000)))
		if err != nil {
			t.Fatal(err)
		}
	}
	
	data, errGetTask := tasksql.GetTask(tableName)
	if errGetTask != nil {
		t.Fatal(err)
	}

	for _,text := range data {
		log.Println(text)
	}
}
