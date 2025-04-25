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

func WrapTextOver(text string, length int, column_nb int, recursive bool) string {
	
	if len(text) <= length {
		if recursive {
			return strings.Repeat(" ", length-len(text)-1) + text
		}
		
		return text
	}

	repeat_count := column_nb*15+1*column_nb+1
	

	return text[:length-2] + "-\n" + strings.Repeat(" ", repeat_count) + WrapTextOver(text[length-2:], length, column_nb, true)
}

func IsStatus(value string) bool {
	status := [3]string{"todo", "pending", "completed"}
	
	for _, v := range status {
		if value == v {
			return true
		}		
	}

	return false
}
