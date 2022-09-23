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
	if err := exec.Command("nmcli", "connection", "show", connName).Run(); err != nil {
		notify(err.Error(), true)
	}

	out, _ := exec.Command("nmcli", "-f", "GENERAL.STATE", "connection", "show", connName).Output()
	if string(out) == "" {
		if err := exec.Command("nmcli", "connection", "up", connName).Run(); err != nil {
			notify(err.Error(), true)
		}

		notify(fmt.Sprintf("Connected to %s", connName), false)

	} else {
		if err := exec.Command("nmcli", "connection", "down", connName).Run(); err != nil {
			notify(err.Error(), true)
		}

		notify(fmt.Sprintf("Disconnected from %s", connName), false)
	}
}

func notify(msg string, exit bool) {
	_ = exec.Command("notify-send", "Network Manager", msg).Run()
	_, _ = os.Stdout.WriteString(msg)

	if exit {
		os.Exit(1)
	}
}
