# Go-ReverseShell

A basic reverse shell. It will create a raw socket to create a tcp connection with server, spawn a cmd.exe process and redirect stdin, stdout and stderr to the socket. If the connection is lost the implant will try to connect back every 3 seconds (you can modify the time).
