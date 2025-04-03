package features

import (
	"fmt"
	"go_task_manager/utils"
	"strconv"
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

	data, err := getFormatedData(cli_args_pointer)
	if err != nil {
		panic(fmt.Sprintf("an error occured : %v", err))
	}
	
	fmt.Printf("data received : %v", data)
}

/*
 * get the data provided by the user in map format
*/
func getFormatedData(args []string) (map[string]([]string), error) {
	var err error

	_, parse_err3 := strconv.Atoi(args[3])
	_, parse_err5 := strconv.Atoi(args[5])
	
	errorMsgs := map[string]string{
		"EMPTY": "excepted task %s. %d value length received",
        "STRING_EXPECTED": "excepted string value of %s. %s received",
	}

	task_name := args[3]
	task_category := args[5]
	var task_description string

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
		if len(args) == 7 || len(args[7]) == 0 {
			err = fmt.Errorf(errorMsgs["EMPTY"], "description", 0)
		} else if _, parse_err8 := strconv.Atoi(args[7]); parse_err8 == nil {
			err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], "description", args[7])
		}else {
			task_description = args[7]
		}
	}

	// return error if any
	if err != nil {
		return nil, err
	}

	var data map[string]([]string)

	if utils.IsSimpleString(task_name) {
		data = map[string]([]string){
			"task_names": {task_name},
			"task_categories": {task_category},
			"task_descriptions": {task_description},
		}

		return data, err
	}

	data, err = getArrayProvidedEntries(task_name, task_category, task_description)

	return data, err
}

/*
 * check if the data provided are array
 * display error if bracket is not closed, eg [1,2,3 or 1,2,3]
 * then returns data provided as array
 */
 func getArrayProvidedEntries(task_name string, task_category string, task_description string) (map[string]([]string), error) {
	var err error
	var data map[string]([]string)

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

	if len(task_names) != len(task_categories) {
		err = fmt.Errorf("names and categories should have the same length. %d and %d received", len(task_names), len(task_categories))

		return data, err
	}

	data = map[string]([]string){
		"task_names": task_names,
		"task_categories": task_categories,
		"task_descriptions": utils.FormatTonatifArray(task_description),
	}

	return data, err
}