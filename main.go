package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 2 {
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Usage: %s <connection name>\n", os.Args[0]))
		os.Exit(1)
	}

	connName := os.Args[1]
	nmcliConnSwitch(connName)
}

func nmcliConnSwitch(connName string) {
	cmd := exec.Command("nmcli", "connection", "show", connName)
	if cmd.Run() != nil {
		// connection not found
		_, _ = os.Stderr.WriteString(fmt.Sprintf("Connection %s not found\n", connName))
		os.Exit(1)
	}

	cmd = exec.Command("nmcli", "-f", "GENERAL.STATE", "connection", "show", connName)
	out, _ := cmd.Output()
	if string(out) == "" {
		cmd = exec.Command("nmcli", "connection", "up", connName)
		_ = cmd.Run()
	} else {
		cmd = exec.Command("nmcli", "connection", "down", connName)
		_ = cmd.Run()
	}
}
