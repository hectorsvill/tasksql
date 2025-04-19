# tasksql 🫠 
The tasksql module provides a lightweight wrapper around SQLite for managing tasks in a database. It includes functionality for creating tables, inserting, updating, deleting, and retrieving task data.

### Setup
```bash
go get github.com/hectorsvill/tasksql
```
import module

```go
import "github.com/hectorsvill/tasksql"
```
add import
 
### Example use of module 
Store prompt data from Google genai: [_example](https://github.com/hectorsvill/tasksql/tree/main/_example)
```go
func main() {

	taskSql, err := tasksql.NewDB("data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer taskSql.Close()

	taskSql.CreateTableIfNotExist("question")
	taskSql.CreateTableIfNotExist("answer")

	question := "Write an article about the golang net/http package."

	taskSql.Post("question", question)
	gem1 := Gemini{
		Model: Gemini_2_0_turbo,
	}
	log.Println(gem1.Model)
	answer := gem1.QueryText(question)
	taskSql.Post("answer", answer[0])
	log.Println(answer)

}
```
#

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
func (tsql TaskSQL) Close() error
```
Close TaskSQL object

```go 
func (tsql TaskSQL) CreateTableIfNotExist(table string) error 
```
Create table with createTableIfNotExist query

```go
func (tsql TaskSQL) Post(table, text string) error
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
func (tsql TaskSQL) Get(table string) ([]string, error)
```
Get text from table
