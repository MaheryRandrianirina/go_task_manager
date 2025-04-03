package main

import (
	"fmt"
	"go_task_manager/features"
	"os"
)

func main() {
	readCommand()
}

/*
 *reads command line
 */
 func readCommand() {
	cli_args := os.Args

	if len(cli_args) < 2 {
		panic(fmt.Errorf("you cannot provide command line with length under 2 words"))
	}

	const COMMAND_NAME string = "gtm"
	if cli_args[1] != COMMAND_NAME {
		panic(fmt.Errorf("an error occured : %s is not valid command. Use %s instead", cli_args[1], COMMAND_NAME))
	}

	command_params_type := map[string]string{
		"CREATE":   "-n",
		"CATEGORY": "-c",
		"DESCRIPTION": "-d",
	}

	first_param := cli_args[2]

	switch first_param {
	case command_params_type["CREATE"]:
		features.CreateTasks(&cli_args, &command_params_type)
	}
}
