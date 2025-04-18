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
	testDB, err := tasksql.NewDB("test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer testDB.CloseTaskSQl()

	cfg = config{
		testDB:     testDB,
		tableNames: []string{"users", "data"},
	}

	code := m.Run()
	log.Println("[TEST START]tasksql")

	os.Exit(code)
	log.Println("[TEST FINISHED]tasksql")
}

func Test_CreateTestData(t *testing.T) {
	tableName := "data"
	err := cfg.testDB.CreateTableIfNotExist(tableName)
	if err != nil {
		t.Fatalf("[Test_CreateTableIfNotExist] %s", err)
	}

	for range 4 {
		text := fmt.Sprintf("text%v", uuid.NewString())
		err = cfg.testDB.Post(tableName, text)
		if err != nil {
			t.Fatal(err)
		}
	}

}

func Test_Get(t *testing.T) {
	data, err := cfg.testDB.Get("")
	if err != nil {
		t.Fatal(err)
	}

	for _, text := range data {
		log.Println(text)
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
		log.Println("expected: false\nactual: ", tc.expected)
		if actual != tc.expected {
			t.Fatalf("expected: false\nactual: %v", actual)
		}
	}
}
