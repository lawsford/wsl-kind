package main

import "fmt"

func displayHelp() {
	helpMessage := `Usage: wslkind <command> [options]
	
	Commands:
	  deploy    Deploy a kind cluster`
	fmt.Println(helpMessage)
}
