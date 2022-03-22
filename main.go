package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		// print usage
		// get file name this binary
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Usage: %s <connection name>\n", os.Args[0]))
		os.Exit(1)
	}

	connName := os.Args[1]
	nmcliConnSwitch(connName)
}

func nmcliConnSwitch(connName string) {
	//run os command
	cmd := exec.Command("nmcli", "-f", "GENERAL.STATE", "con", "show", connName)
	out, _ := cmd.Output()
	if string(out) == "" {
		cmd = exec.Command("nmcli", "con", "up", connName)
		_ = cmd.Run()
	} else {
		cmd = exec.Command("nmcli", "con", "down", connName)
		_ = cmd.Run()
	}
}
