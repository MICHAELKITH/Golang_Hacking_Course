package main

import (
 "fmt"
 "os"
 "os/signal"
 "syscall"
)

func main() {
 // Start the keylogger routine
 go keylogger()

 // Wait for termination signal to gracefully exit
 waitForExitSignal()
}

func keylogger() {
 // Keylogger implementation goes here
}

func waitForExitSignal() {
 c := make(chan os.Signal, 1)
 signal.Notify(c, os.Interrupt, syscall.SIGTERM)
 <-c

 // Clean up and exit
 fmt.Println("Exiting 1...")
 os.Exit(0)
}