package tasksql_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/hectorsvill/tasksql"
)

type config struct {
	testDB     *tasksql.TaskSQL
	tableNames []string
}

var cfg config

func TestMain(m *testing.M) {
	var err error
	dbName := fmt.Sprintf("testdb_%s", uuid.NewString())
	testDB, err := tasksql.NewDB(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.CloseTaskSQl()

	cfg = config{
		testDB:     testDB,
		tableNames: []string{"users", "data", "notes"},
	}

	code := m.Run()
	// log.Println("[TEST START]tasksql")

	os.Remove(dbName)
	os.Exit(code)
	// log.Println("[TEST FINISHED]tasksql")
}

func Test_CreateTestData(t *testing.T) {
	for _, table := range cfg.tableNames {
		err := cfg.testDB.CreateTableIfNotExist(table)
		if err != nil {
			t.Fatalf("[Test_CreateTestData]: %s", err)
		}

		for range 12 {
			text := fmt.Sprintf("text%v", uuid.NewString())
			err = cfg.testDB.Post(table, text)
			if err != nil {
				t.Fatalf("[Test_CreateTestData]: %s", err)
			}
		}
	}
}

func Test_Get(t *testing.T) {
	// var data []string
	// var err error
	for _, table := range cfg.tableNames {
		_, err := cfg.testDB.Get(table)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func Test_DeleteWhereDeletedTrue(t *testing.T) {

}

func Test_IsValidTableID(t *testing.T) {
	type TestCase struct {
		input    string
		expected bool
	}

	testCases := []TestCase{
		{input: "data", expected: true},
		{input: "users", expected: true},
		{input: "213s_dqSelect-a1", expected: false},
		{input: "SELECT * FROM data;", expected: false},
		{input: "SELECT text FROM data WHERE id = 4;", expected: false},
	}

	for _, tc := range testCases {
		actual := tasksql.IsValidTableID(tc.input)
		if actual != tc.expected {
			t.Fatalf("expected: false actual: %v", tc.expected)
		}
	}
}
