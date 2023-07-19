package deploy

import (
	"fmt"
	"os"
	"os/exec"
)

func DeployKindCluster() {
    cmd := exec.Command("kind", "create", "cluster")
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
	
	err := cmd.Run()
    if err != nil {
        fmt.Printf("Error creating Kind cluster: %v\n", err)
        os.Exit(1)
    }

	fmt.Println("Kind cluster created successfully!")
}