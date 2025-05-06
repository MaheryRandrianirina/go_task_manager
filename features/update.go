package features

import (
	"fmt"
	"go_task_manager/utils"
	"strconv"
	"time"
)

func Update(cli_args *[]string) {
	cli_args_pointer := *cli_args
	if len(cli_args_pointer) < 5 {
		panic("You should provide at least 2 arguments: -u <task_id> <new_task>")
	}

	task_id, err := strconv.Atoi(cli_args_pointer[3])
	if err != nil {
		panic("You should provide integer value as task id")
	}

	tasks, err := utils.GetTasks()
	if err != nil {
		panic(fmt.Sprintf("Error getting tasks: %v", err))
	}else if task_id > len(tasks) {
		panic(fmt.Sprintf("Task with id %d does not exist", task_id))
	}

	for i, task := range tasks {
		id := i+1
		if id != task_id {
			continue
		}

		task_name := task.Name
		if cli_args_pointer[4] != "" && cli_args_pointer[4] != "-s" && cli_args_pointer[4] != "-dd" {
			task_name = cli_args_pointer[4]
		}

		task_category := task.Category
		task_description := task.Description
		task_status := task.Status
		task_due_date := task.Date

		if len(cli_args_pointer) > 5 {
			if cli_args_pointer[4] != "-s" && cli_args_pointer[4] != "-dd" && cli_args_pointer[5] != "" {
				task_category = cli_args_pointer[5]
			}else if cli_args_pointer[4] == "-s" && utils.IsStatus(cli_args_pointer[5]) {
				task_status = cli_args_pointer[5]
			}else if _, err := time.Parse("2006-01-02", cli_args_pointer[5]); cli_args_pointer[4] == "-dd" && err == nil {
				task_due_date = cli_args_pointer[5]
			}else if cli_args_pointer[4] == "-s" && !utils.IsStatus(cli_args_pointer[5]) {
				panic(fmt.Sprintf("You should provide a status (todo, pending, completed) as a value of -s. %s provided", cli_args_pointer[5]))
			}else if _, err := time.Parse("2006-01-02", cli_args_pointer[5]); cli_args_pointer[4] == "-dd" && err != nil {
				panic(fmt.Sprintf("You should provide a date YYYY/mm/dd as a value of -dd. %s provided", cli_args_pointer[5]))
			}
			

			if cli_args_pointer[4] != "-s" && cli_args_pointer[4] != "-dd" && len(cli_args_pointer) > 6 && cli_args_pointer[6] != "" {
				task_description = cli_args_pointer[6]
			}

		}

		updated_task := utils.Task{
			Name:        task_name,
			Category:    task_category,
			Description: task_description,
			Status : task_status,
			Date: task_due_date,
		}
		
		tasks[i] = updated_task
	}

	err = utils.CreateFile(utils.FILENAME, &tasks)
	if err != nil {
		panic(fmt.Sprintf("Error creating file: %v", err))
	}

	fmt.Printf("Task %d updated successfully\n", task_id)
}