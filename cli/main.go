package main

import (
	"fmt"
	"os"

	"github.com/lawsford/wsl-kind/cli/cmd/deploy"
	"github.com/lawsford/wsl-kind/cli/helpers"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--help" {
		displayHelp()
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: wslkind <command> [options]")
		fmt.Println("Use 'wslkind --help' to see available commands.")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "deploy":
		fmt.Println("You are about to deploy a bare minimum Kind cluster with 1 node. Proceed?")
		if helpers.GetUserConfirmation() {
			deploy.DeployKindCluster()
		} else {
			fmt.Println("Deployment aborted...")
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Use 'cli --help' to see available commands.")
		os.Exit(1)
	}
}
