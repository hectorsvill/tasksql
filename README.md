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
 
### Use case 
```go
func setCfg() config {
	taskSql, err := tasksql.NewDB("data.db")
	if err != nil {
		log.Fatal(err)
	}

	cfg := config{tsql: taskSql}

	if err := cfg.tsql.CreateTableIfNotExist(); err != nil {
		log.Fatal(err)
	}

	getTasks(cfg)
	return cfg
}


func main() {
	cfg := setCfg()
	runTask(cfg)
	defer cfg.tsql.CloseTaskSQl()
}

```


```go
const (
	tableName = "tasks"
	CreateTable = "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY AUTOINCREMENT, text TEXT NOT NULL, completed BOOLEAN DEFAULT FALSE);"
	DeleteWhereCompleted = "DELETE FROM tasks WHERE completed = ?;"
	UpdateTaskCompletedWereID = "UPDATE tasks SET completed = ? WHERE id = ?;"
	InsertTask = "INSERT INTO tasks (text) VALUES (?);"
	SelectAllTasks = "SELECT * FROM Tasks;"
)

```
Sql queries constants

```go
type TaskSQL struct{
	db *sql.DB
}

type Task struct {
	id        int
	text      string
	completed bool
}
```
structure for TaskSQL and Task

```go 
func (tsql TaskSQL) CreateTable() error
```
Create table with CreateTable query

```go
func (tsql TaskSQL) PutTask(task string) error {
```
Put a task in sql db

```go
func (tsql TaskSQL) DeleteTask() error 
```
Delete table from db

```go
func (tsql TaskSQL) UpdateTaskToCompleted(id int) error 
```
update to true by finding id

```go
func (tsql TaskSQL) GetTask() ([]Task, error) 
```
Get items from tasks table
