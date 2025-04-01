package main

import (
	"fmt"
	"os"
	"strconv"
)

func displayOccuredError(message string) {
	err := fmt.Errorf("an error occured : %v", message)
	if err != nil {
		fmt.Println(err)
	}
}

/*
 *reads command line
 */
func readCommand() {
	cli_args := os.Args

	if len(cli_args) < 2 {
		displayOccuredError("You cannot provide command line with length under 2 words.")

		return
	}

	const COMMAND_NAME string = "gtm"
	if cli_args[1] != COMMAND_NAME {
		displayOccuredError(cli_args[1] + " is not valid command. Use '" + COMMAND_NAME + "' instead.")

		return
	}

	command_params_type := map[string]string{
		"CREATE":   "-n",
		"CATEGORY": "-c",
		"DESCRIPTION": "-d",
	}

	first_param := cli_args[2]
	switch first_param {
	case command_params_type["CREATE"]:
		handleTaskCreation(&cli_args, &command_params_type)
	}
}

func handleTaskCreation(cli_args *[]string, command_params_type *map[string]string) {
	cli_args_pointer := *cli_args
	command_params_type_pointer := *command_params_type

	if len(cli_args_pointer) < 6 {
		displayOccuredError("6 args expected. " + strconv.Itoa(len(cli_args_pointer)) + " given.")

		return
	}

	if cli_args_pointer[4] != command_params_type_pointer["CATEGORY"] {
		displayOccuredError("'-c' expected. " + cli_args_pointer[4] + " received")

		return
	}

	if len(cli_args_pointer) > 6 && cli_args_pointer[6] != command_params_type_pointer["DESCRIPTION"] {
		displayOccuredError("'-d' expected. " + cli_args_pointer[6] + " received")

		return
	}

	err := isEntryValid(cli_args_pointer)
	if err != nil {
		fmt.Printf("an error occured : %v", err)
	}

}

func isEntryValid(args []string) error {
	var err error

	_, parse_err3 := strconv.Atoi(args[3])
	_, parse_err5 := strconv.Atoi(args[5])
	
	errorMsgs := map[string]string{
		"EMPTY": "excepted task %s. %d value length received",
        "STRING_EXPECTED": "excepted string value of %s. %s received",
	}

	if len(args[3]) == 0 {
		err = fmt.Errorf(errorMsgs["EMPTY"], len(args[3]))
	} else if parse_err3 == nil {
		err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], args[3])
	} else if len(args[5]) == 0 {
		err = fmt.Errorf(errorMsgs["EMPTY"], len(args[5]))
	} else if parse_err5 == nil {
		err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], args[5])
	}

	if len(args) > 6 {
		if len(args) == 7 || len(args[7]) == 0{
				err = fmt.Errorf(errorMsgs["EMPTY"], "description", 0)
		}else if len(args[7]) == 0 {
				err = fmt.Errorf(errorMsgs["EMPTY"], "description", len(args[7]))
		} else if _, parse_err8 := strconv.Atoi(args[7]); parse_err8 == nil {
				err = fmt.Errorf(errorMsgs["STRING_EXPECTED"], "description", args[7])
		}
	}


	return err
}

func main() {
	readCommand()

}
