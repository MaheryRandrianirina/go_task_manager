package features

import (
	"fmt"
	"go_task_manager/utils"
	"os"
	"strconv"
	"strings"
)

func ListTasks(cli_args *[]string) {
	if _, err := os.Stat(utils.FILENAME); os.IsNotExist(err) {
		fmt.Println("You don't have any tasks yet")
		return
	}

	tasks, err := utils.GetTasks()
	if err != nil {
		panic(fmt.Sprintf("Error fetching your tasks list: %v", err))
	}

	if len(tasks) == 0 {
		fmt.Println("You don't have any tasks yet")
		return
	}

	table, err := getArrayTable(tasks, cli_args)
	if err != nil {
		panic(fmt.Sprintf("Error fetching your tasks list: %v", err))
	}

	fmt.Println(table)
}

// get a table with name, category and description as columns
// and the tasks as rows
func getArrayTable(tasks []utils.Task, cli_args *[]string) (string, error) {
	var tableSb strings.Builder
	cli_args_pointer := *cli_args

	if len(cli_args_pointer) > 3 && cli_args_pointer[3] != "-n" && cli_args_pointer[3] != "-c" && cli_args_pointer[3] != "-s" {
		return "", fmt.Errorf("only -n, -c & -s args work with -l. %s provided", cli_args_pointer[3])
	} else if len(cli_args_pointer) == 4 && cli_args_pointer[3] == "-n" {
		return "", fmt.Errorf("arg '-n' requires a number to work. For exampe : -n 2")
	}

	tableSb.WriteString(fmt.Sprintf("%-15s %-15s %-15s %-15s %-15s %-15s\n%s\n", "ID", "|Name", "|Category", "|Description", "|Due date", "|Status", strings.Repeat("-", 100)))

	for i, task := range tasks {
		id := i + 1

		// show only tasks with the category provided in the command line
		// for example : gtm -l -c category_name
		if len(cli_args_pointer) > 3 && cli_args_pointer[3] == "-c" {
			if task.Category != cli_args_pointer[4] {
				continue
			}
		} else if len(cli_args_pointer) > 3 && cli_args_pointer[3] == "-s" {
			// show tasks with status mentioned with the -s arg
			if task.Status != cli_args_pointer[4] {
				continue
			}
		}

		tableSb.WriteString(fmt.Sprintf("%-15d %-15s %-15s %-15s %-15s %-15s\n", id, 
			"|"+utils.WrapTextOver(task.Name, 15, 1, false), 
			"|"+utils.WrapTextOver(task.Category, 15, 2, false), 
			"|"+utils.WrapTextOver(task.Description, 15, 3, false), 
			"|"+task.Date, "|"+task.Status))

		if len(cli_args_pointer) > 4 && cli_args_pointer[3] == "-n" {
			list_nb, err := strconv.Atoi(cli_args_pointer[4])
			if err != nil {
				return "", err
			}

			// truncate the table if the user provided a number
			if list_nb == id {
				tableSb.WriteString(strings.Repeat("-", 100) + "\n")

				break
			}
		}

		if i == len(tasks)-1 {
			tableSb.WriteString(strings.Repeat("-", 100) + "\n")
		}
	}

	tableSb.WriteString(fmt.Sprintf("Total: %d task(s).", len(tasks)) + "\n")

	return tableSb.String(), nil
}
