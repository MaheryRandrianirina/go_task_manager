package features

import (
	"fmt"
	"go_task_manager/utils"
	"os"
	"sort"
	"strconv"
)

func Delete(cli_args *[]string) {
	cli_args_pointer := *cli_args

	// remove all tasks if no task id is specified : -r only
	if len(cli_args_pointer) == 3 {
		removeFileContainingTasks()

		return
	}

	// remove a specific task if task id is specified : -r <task_id>|<task_id[]>
	if utils.IsSimpleString(cli_args_pointer[3]) {
		intvalue, err := strconv.Atoi(cli_args_pointer[3])
		if err != nil {
			panic(fmt.Sprintf("You should provide integer value as task id. %v is not an integer", err))
		}

		msg, err := removeTask(intvalue)
		if err != nil {
			panic(fmt.Sprintf("Error removing task: %v", err))
		}

		fmt.Println(msg)
	}else if utils.IsValidArray(cli_args_pointer[3]) {
		task_ids := utils.FormatTonatifArray(cli_args_pointer[3])
		int_task_ids := make([]int, len(task_ids))

		for i, id := range task_ids {
			intvalue, err := strconv.Atoi(id)
			if err != nil {
				panic(fmt.Sprintf("You should provide integer value as task id. %v is not an integer", err))
			}

			int_task_ids[i] = intvalue
		}

		msg, err := removeTask(int_task_ids)
		if err != nil {
			panic(fmt.Sprintf("Error removing task: %v", err))
		}

		fmt.Println(msg)

	}
}

func removeFileContainingTasks() error {
	err := os.Remove(utils.FILENAME)
	if err != nil {
		return err
	}

	fmt.Println("All tasks have been deleted.")

	return nil
}

func removeTask(task_id interface{}) (string, error) {
	var msg string
	tasks, err := utils.GetTasks()
	if err != nil {
		return "", err
	}

	if len (tasks) == 0 {
		return "", fmt.Errorf("you don't have any tasks to delete")	
	}


	switch t_id := task_id.(type) {
	case int:
		if t_id < 1 || t_id > len(tasks) {
			return msg, fmt.Errorf("task id should be between 1 and %d", len(tasks))
		}
		
		tasks = append(tasks[:t_id-1], tasks[t_id:]...)
		msg = fmt.Sprintf("Task with id %d has been removed", t_id)
		
	case []int:
		// sort the task ids in increasing order
		sort.Ints(t_id)
		
		for index, id := range t_id {
			fmt.Println(tasks)
			// In first loop turn, we remove the element with the the first id provided
			// With that, ids (that id index+1) changed. It means that the remaining ids provided by the user became id-index
			if index > 0 {
				id = id - index
			}

			tasks = append(tasks[:id-1], tasks[id:]...)
		}
		
		msg = fmt.Sprintf("Tasks with ids %v have been removed", task_id)
	}

	err = utils.CreateFile(utils.FILENAME, &tasks)

	return msg, err
}