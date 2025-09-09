package main
import "fmt"
func CallbackHelp() error {

	availableCommands := getCommand()
	fmt.Println("Welcome to Sharadex menu")
	fmt.Println("Here are the list of commands")

	for _,v := range availableCommands {
		fmt.Printf("- %s:  %s\n", v.name, v.description)
		
	}
	fmt.Println("")
	return nil
}