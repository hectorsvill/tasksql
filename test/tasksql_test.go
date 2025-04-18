package tasksql_test

import (
	"fmt"
	"log"
	"math/rand"
	"testing"

	"github.com/hectorsvill/tasksql"
)

func Test_Complete(t *testing.T) {
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

	for range 4 {
		text := fmt.Sprintf("text%v", rand.Intn(1000))
		err = tasksql.Post(tableName, text)
		if err != nil {
			t.Fatal(err)
		}
	}

	data, errGetTask := tasksql.Get(tableName)
	if errGetTask != nil {
		t.Fatal(err)
	}

	for _, text := range data {
		log.Println(text)
	}
}

func Test_IsValidTableID(t *testing.T) {
	type actualExpected struct {
		input    string
		expected bool
	}

	testCases := []actualExpected{
		{input: "data", expected: true},
		{input: "213sdqSelec-te", expected: false},
	}

	for _, tc := range testCases {
		actual := tasksql.IsValidTableID(tc.input)
		log.Println("expected: false\nactual: ", tc.expected)
		if actual != tc.expected {
			t.Fatalf("expected: false\nactual: %v", actual)
		}
	}
}
