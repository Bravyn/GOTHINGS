package main

import (
	"fmt"
	"os/exec"
)

func main(){
	//get current directory
	cmd := exec.Command("powershell", "-Command", "Get-Location")

	output, err := cmd.CombinedOutput()

	if err!= nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(output))
}