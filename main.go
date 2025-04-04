package main

import (
	"fmt"
	"go_task_manager/features"
	"os"
)

func main() {
	cli_args, err := readTerminal()
	if err != nil {
		panic(err)
	}

	cli_args_pointer := *cli_args
	command_params_type := map[string]string{
		"CREATE":   "-n",
		"CATEGORY": "-c",
		"DESCRIPTION": "-d",
	}

	first_param := cli_args_pointer[2]
	switch first_param {
	case command_params_type["CREATE"]:
		features.CreateTasks(cli_args, &command_params_type)
	}
}


/*
 *reads command line
 */
 func readTerminal() (*[]string, error) {
	cli_args := os.Args
	var err error

	if len(cli_args) < 2 {
		err = fmt.Errorf("you cannot provide command line with length under 2 words")
	}

	const COMMAND_NAME string = "gtm"
	if cli_args[1] != COMMAND_NAME {
		err = fmt.Errorf("%s is not valid command. Use %s instead", cli_args[1], COMMAND_NAME)
	}

	return &cli_args, err
}
