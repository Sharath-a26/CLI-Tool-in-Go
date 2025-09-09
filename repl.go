package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
)
//struct type to define a command
type CliCommand struct {
	name string
	description string
	callback func() error 
}

//a function that returns the map of these commands  - the map stores the name to a struct pair
func getCommand() map[string]CliCommand {
	return map[string]CliCommand{
		"help": {
			name: "help",
			description: "Prints the help menu",
			callback: CallbackHelp,
		},
		"exit": {
			name: "exit",
			description: "Turns off the sharadex",
			 callback: CallbackExit,
		},
		"map": {
			name: "map",
			description: "List some areas of pokemon",
			 callback: callBackMap,
		}, 
	}
}
func startRepl() {
scanner := bufio.NewScanner(os.Stdin)
	
	for {
	fmt.Print("Enter text: ")
	scanner.Scan()
	text := scanner.Text()

	cleaned := cleanInput(text)
	

	if len(cleaned) == 0 {
		continue
	}
	

	//first word is a command 
	command:= cleaned[0]
	availableCommands := getCommand()

	com, ok := availableCommands[command]

	if !ok {
		fmt.Println("invalid command")
		continue
	}
	com.callback()
	
		
	
	}
}

func cleanInput(str string) []string {
	 lowered := strings.ToLower(str)
	 words:= strings.Fields(lowered)
	 return words

}