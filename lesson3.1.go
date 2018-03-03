package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
	if err := cmd0.Start(); err != nil {
		fmt.Printf("Error: The command No.0 can not be startup:%s\n", err)
		return
	}

	stdout0, err := cmd0.StdoutPipe()

	if err != nil {
		fmt.Printf("Error: Couldn't obtain the stdout pipe for command No.0:%s\n", err)
	}

	output0 := make([]byte, 30)
	m, err := stdout0.Read(output0)
	if err != nil {
		fmt.Printf("Error: Couldn't read data from the pipe:%s\n", err)
		return
	}

	fmt.Printf("%s\n", output0[:m])
}
