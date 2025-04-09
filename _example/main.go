package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hectorsvill/tasksql"
)

func getInput(str string) string {
	fmt.Print("\n" + str + ">: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("error")
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	return input
}

func runTask(cfg config) {
	
	for {
		fmt.Print("\n_______\n\nSet Task (s)\nView Tasks (v)\nUpdate Task To Completed(u)\nDelete all Completed Task(d)\nexit (x)\n-------")
		sv := getInput("")
		if sv == "s" {
			input := getInput("set")
			err := cfg.tsql.PostTask(input)
			if err != nil {
				log.Println(err)
			}
			fmt.Print("\n______________\nTask added\n______________\n")
			cfg.tasks = append(cfg.tasks, tasksql.Task{
				Id: len(cfg.tasks),
				Text: input,
				Completed: false,
				},
			)
		} else if sv == "v" {
			if cfg.tasks == nil {
				tasks, err := cfg.tsql.GetTask()
				if err != nil {
					log.Println(err)
				}
				cfg.tasks = tasks
			}
			printTask(cfg.tasks)
		} else if sv == "u" {
			printTask(cfg.tasks)
			id := getInput("Enter id of task to update")
			idInt, err := strconv.Atoi(id)
			if err != nil || (idInt < 1 && idInt > len(cfg.tasks)) {
				fmt.Println("Invalid id")
				continue
			}
			err = cfg.tsql.UpdateTaskToCompleted(cfg.tasks[idInt].Id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			cfg.tasks[idInt].Completed = true
			fmt.Print("\n______________\nTask updated\n______________\n")
		} else if sv == "d" {
			err := cfg.tsql.DeleteWhereTaskCompleted()
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("\n______________\nTask deleted\n______________\n")
			cfg.tasks = nil
			getTasks(cfg)

		} else {
			fmt.Println("Goodbye")
			return
		}
	}
}
func getTasks(cfg config) {
	if tasks, err := cfg.tsql.GetTask(); err != nil {
		log.Fatal(err)
	} else {
		cfg.tasks = tasks
	}
}


func printTask(tasks []tasksql.Task) {
	for i, t := range tasks {
		fmt.Printf("\n%v ] %v {%v}\n", i, t.Text, t.Completed)
	}
}

type config struct {
	tsql *tasksql.TaskSQL
	tasks []tasksql.Task
}

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
	// defer cfg.tsql.taskSql.CloseTaskSQl()
}