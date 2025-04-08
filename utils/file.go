package utils

import (
	"encoding/json"
	"io"
	"os"
)

func GetTasks() ([]Task, error) {
	file, err := os.Open(FILENAME)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileByteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(fileByteValue, &tasks)
	if err != nil {
		return nil, err	
	}

	return tasks, nil
}

/*
 * create the json file that will contain the data
 */
 func CreateFile(filename string, data *[]Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	// write the data to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(*data)
	if err != nil {
		return err
	}

	return nil
}