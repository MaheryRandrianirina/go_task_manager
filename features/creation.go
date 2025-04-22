package features

import (
	"encoding/json"
	"fmt"
	"go_task_manager/utils"
	"io"
	"os"
	"strconv"
	"time"
)

func CreateTasks(cli_args *[]string, command_params_type *map[string]string) {
	cli_args_pointer := *cli_args
	command_params_type_pointer := *command_params_type

	if len(cli_args_pointer) < 6 {
		panic(fmt.Errorf("6 args expected. %d given", len(cli_args_pointer)))
	}

	if cli_args_pointer[4] != command_params_type_pointer["CATEGORY"] {
		panic(fmt.Errorf("'-c' expected. %s received", cli_args_pointer[4]))
	}

	if len(cli_args_pointer) > 6 && cli_args_pointer[6] != command_params_type_pointer["DESCRIPTION"] {
		panic(fmt.Errorf("'-d' expected. %s received", cli_args_pointer[6]))
	}

	if len(cli_args_pointer) > 8 && cli_args_pointer[8] != command_params_type_pointer["DATE"] {
		panic(fmt.Errorf("'-dd' expected. %s received", cli_args_pointer[6]))
	}

	data, err := getFormatedData(cli_args_pointer)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	err = insertData(&data)
	if err != nil {
		panic(fmt.Sprintf("%v", err))
	}

	fmt.Println("Task(s) inserted successfully")
}

/*
 * get the data provided by the user in map format
 */
func getFormatedData(args []string) ([]utils.Task, error) {
	var err error

	_, parse_err3 := strconv.Atoi(args[3])
	_, parse_err5 := strconv.Atoi(args[5])

	errorMsgs := map[string]string{
		"EMPTY":           "excepted task %s. %d value length received",
		"STRING_EXPECTED": "excepted string value of %s. %s received",
	}

	task_name := args[3]
	task_category := args[5]
	var task_description string
	var task_due_date string

	if len(task_name) == 0 {
		err = fmt.Errorf(errorMsgs["EMPTY"], len(task_name))
	} else if parse_err3 == nil {
		err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], task_name)
	} else if len(task_category) == 0 {
		err = fmt.Errorf(errorMsgs["EMPTY"], len(task_category))
	} else if parse_err5 == nil {
		err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], task_category)
	}

	// check if the description is provided and valid
	if len(args) > 6 {

		if len(args) == 7 {
			err = fmt.Errorf(errorMsgs["EMPTY"], "description", 0)
		} else if _, parse_err8 := strconv.Atoi(args[7]); parse_err8 == nil {
			err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], "description", args[7])
		}else if len(args) == 9 || len(args[9]) == 0 {
			err = fmt.Errorf(errorMsgs["EMPTY"], "due date", 0)
		}else if len(args) == 10 {
			var due_date time.Time
			due_date, err = time.Parse("2006-01-02", args[9])
				
			duration := time.Until(due_date)
			days := int(duration.Hours()/24)
	
			if err != nil {
				err = fmt.Errorf("due date value should be a valid date. %s received", args[9])
			} else if days < 0 {
				err = fmt.Errorf("due date should be a future date")
			}
	
			task_due_date = args[9]
		}

		task_description = args[7]		
	}

	// return error if any
	if err != nil {
		return nil, err
	}

	data := make([]utils.Task, 0)
	data_pointer := &data
	if utils.IsSimpleString(task_name) {
		data = append(*data_pointer, utils.Task{
			Name:        task_name,
			Category:    task_category,
			Description: task_description,
			Date:        task_due_date,
		})

		return data, err
	}

	data, err = getArrayProvidedEntries(task_name, task_category, task_description, data_pointer)

	return data, err
}

/*
 * check if the data provided are array
 * display error if bracket is not closed, eg [1,2,3 or 1,2,3]
 * then returns data provided as array
 */
func getArrayProvidedEntries(task_name string, task_category string, task_description string, data_pointer *[]utils.Task) ([]utils.Task, error) {
	var err error

	if !utils.IsValidArray(task_name) {
		err = fmt.Errorf("task_name is not valid array. %s received", task_name)
	}

	if !utils.IsValidArray(task_category) {
		err = fmt.Errorf("task_category is not valid array. %s received", task_category)
	}

	if err == nil && len(task_description) > 0 && !utils.IsValidArray(task_description) {
		err = fmt.Errorf("task_description is not valid array. %s received", task_description)
	}

	task_names := utils.FormatTonatifArray(task_name)
	task_categories := utils.FormatTonatifArray(task_category)
	task_descriptions := utils.FormatTonatifArray(task_description)

	if len(task_names) != len(task_categories) {
		err = fmt.Errorf("names and categories should have the same length. %d and %d received", len(task_names), len(task_categories))

		return nil, err
	}

	data := utils.FormatToTaskSlice(&task_names, &task_categories, &task_descriptions, data_pointer)

	return data, err
}

func insertData(data *[]utils.Task) error {
	// check if the file exists
	// if not, create it and write the data to it
	if _, err := os.Stat(utils.FILENAME); os.IsNotExist(err) {
		return utils.CreateFile(utils.FILENAME, data)
	}

	file, err := os.Open(utils.FILENAME)
	if err != nil {
		return err
	}

	defer file.Close()

	fileByteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	var tasks []utils.Task

	err = json.Unmarshal(fileByteValue, &tasks)
	if err != nil {
		return err
	}
	// append the new data to the existing data
	tasks = append(tasks, *data...)

	file.Close()

	return utils.CreateFile(utils.FILENAME, &tasks)
}
