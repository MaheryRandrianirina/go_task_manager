package features

import (
	"fmt"
	"go_task_manager/utils"
	"os"
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
	
}