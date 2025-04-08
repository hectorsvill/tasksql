package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
	"strconv"
	"strings"
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
			cfg.tasks = append(cfg.tasks, Task{id: len(cfg.tasks), text: input})
		} else if sv == "v" {
			if cfg.tasks == nil {
				tasks, err := cfg.tsql.GetTask()
				if err != nil {
					log.Println(err)
				}
				cfg.tasks = tasks
			}
			cfg.tsql.printTask(cfg.tasks)
		} else if sv == "u" {
			cfg.tsql.printTask(cfg.tasks)
			id := getInput("Enter id of task to update")
			idInt, err := strconv.Atoi(id)
			if err != nil || (idInt < 1 && idInt > len(cfg.tasks)) {
				fmt.Println("Invalid id")
				continue
			}
			err = cfg.tsql.UpdateTaskToCompleted(cfg.tasks[idInt].id)
			if err != nil {
				fmt.Println(err)
				continue
			}
			cfg.tasks[idInt].completed = true
			fmt.Print("\n______________\nTask updated\n______________\n")
		} else if sv == "d" {
			err := cfg.tsql.DeleteTask()
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

type config struct {
	tsql TaskSQL
	tasks []Task
}

func setCfg() config {
	cfg := config{
		tsql: TaskSQL {
			dbFile: "task.db",
		},
	}

	if err := cfg.tsql.CreateTable(); err != nil {
		log.Fatal(err)
	}
	
	getTasks(cfg)
	
	return cfg
}


func main() {
	runTask(setCfg())
}