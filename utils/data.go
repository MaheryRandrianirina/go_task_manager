package utils

import (
	"strings"
)

/*
 * transform string to natif array value knowing that array provided via command line is read as string
 * @param value string
 * @return []string
 */
func FormatTonatifArray(value string) []string {
	if len(value) < 3 {
		return []string{}
	}

	bracket_removed_value := value[1 : len(value)-1]
	return strings.Split(bracket_removed_value, ",")
}

func IsValidArray(value string) bool {
	if len(value) < 2 {
		return false
	}

	if value[0] != '[' || value[len(value)-1] != ']' {
		return false
	}

	return true
}

/*
 * check if the string is simple string or not an array
*/
func IsSimpleString(value string) bool {
	if len(value) < 2 {
		return true
	}

	return value[0] != '[' && value[len(value)-1] != ']'
}

func FormatToTaskSlice(names *[]string, categories *[]string, descriptions *[]string, tasks_pointer *[]Task) []Task {
	
	for index, name := range *names {
		task := Task{
			Name:        name,
			Category:    (*categories)[index],
			Description: "",
		}

		if len(*descriptions) > 0 {
			task.Description = (*descriptions)[index]
		}
		
		*tasks_pointer = append(*tasks_pointer, task)
	}
	
	return *tasks_pointer
}

	