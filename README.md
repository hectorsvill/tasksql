# tasksql 🫠 

### Setup
```bash
go get github.com/hectorsvill/tasksql
```
import module

```go
import "github.com/hectorsvill/tasksql"
```
add import
 
### Example 
Store prompt data from Google genai: [_example](https://github.com/hectorsvill/tasksql/tree/main/_example)
```go
func main() {

	taskSql, err := tasksql.NewDB("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer taskSql.CloseTaskSQl()

	taskSql.CreateTableIfNotExist("question")
	taskSql.CreateTableIfNotExist("answer")

	question := "Write an article about the golang net/http package."

	taskSql.PostTask("question", question)
	gem1 := Gemini{
		Model: Gemini_2_0_turbo,
	}
	log.Println(gem1.Model)
	answer := gem1.QueryText(question)
	taskSql.PostTask("answer", answer[0])
	log.Println(answer)

}
```
Create a tasksql object to pass arround and access methods 

```go
const (
	tableName = "tasks"
	createTableIfNotExist = "CREATE TABLE IF NOT EXISTS {table} (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, deleted BOOLEAN DEFAULT FALSE);"
	deleteWhereDeletedTrue = "DELETE FROM {table} WHERE completed = ?;"
	updateDeletedTrueWereID = "UPDATE {table} SET deleted = ? WHERE id = ?;"
	insertTasksValueText = "INSERT INTO {table} (text) VALUES (?);"
	selectAllText = "SELECT text FROM {table} ORDER BY id"
)

```
Sql queries constants

```go
type TaskSQL struct{
	db *sql.DB
}

```
TaskSQL object to pass around

```go
func NewDB(dbSourceName string) (*TaskSQL, error)
```
Create TaskSQL object

```go
func (tsql TaskSQL) CloseTaskSQl() error
```
Close TaskSQL object

```go 
func (tsql TaskSQL) CreateTableIfNotExist(table string) error 
```
Create table with createTableIfNotExist query

```go
func (tsql TaskSQL) PostTask(table, text string) error
```
Post to db with text

```go
func (tsql TaskSQL) DeleteWhereDeletedTrue(table string) error
```
Delete table from db

```go
func (tsql TaskSQL) UpdateToDelete(table string, id int) error 
```
update to true by finding id

```go
func (tsql TaskSQL) GetTask(table string) ([]string, error)
```
Get text from table
