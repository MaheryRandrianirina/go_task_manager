package features

import (
	"encoding/json"
	"fmt"
	"go_task_manager/utils"
	"io"
	"os"
	"strings"
)

func ListTasks(cli_args *[]string, command_params_type *map[string]string) {
	if _, err := os.Stat(utils.FILENAME); os.IsNotExist(err) {
		fmt.Println("You don't have any tasks yet")
		return 
	}

	file, err := os.Open(utils.FILENAME)
	if err != nil {
		panic(fmt.Sprintf("Error opening file containing tasks: %v", err))
	}

	defer file.Close()

	fileByteValue, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("Error reading file containing tasks: %v", err))
	}

	var tasks []utils.Task
	err = json.Unmarshal(fileByteValue, &tasks)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling file containing tasks. File might not be json type: %v", err))
	}

	if len(tasks) == 0 {
		fmt.Println("You don't have any tasks yet")
		return
	}

	table := getArrayTable(tasks)

	fmt.Println(table)
}

// get a table with name, category and description as columns
// and the tasks as rows
func getArrayTable(tasks []utils.Task) string {
	var tableSb strings.Builder

	tableSb.WriteString(fmt.Sprintf("%-20s %-20s %-20s %-20s\n%s\n", "ID", "|Name", "| Category","| Description", strings.Repeat("-", 60)))
	
	for i, task := range tasks {
		id := i+1
		tableSb.WriteString(fmt.Sprintf("%-20d %-20s %-20s %-20s\n", id, task.Name, "| "+task.Category, "| "+task.Description))
		if i == len(tasks)-1 {
			tableSb.WriteString(strings.Repeat("-", 60) + "\n")
		}
	}

	tableSb.WriteString(fmt.Sprintf("Total: %d tasks.", len(tasks)) + "\n")

	return tableSb.String()
}