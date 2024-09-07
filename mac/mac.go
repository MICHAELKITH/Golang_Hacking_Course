package main

import (
    "flag"
    "fmt"
    "os"
    "os/exec"
)

func main() {
    interfaceName := flag.String("interface", "", "Name of the network interface")
    newMacAddress := flag.String("mac", "", "New MC address")
    flag.Parse()

    if *interfaceName == "" || *newMacAddress == "" {
        fmt.Println("Usage: change_mac -interface [interface_name] -mac [new_mac_address]")
        os.Exit(1)
    }

    // Disable the network interface
    exec.Command("sudo", "ifconfig", *interfaceName, "down").Run()

    // Change the MAC address
    exec.Command("sudo", "ifconfig", *interfaceName, "hw", "ether", *newMacAddress).Run()

    // Re-enable the network interface
    exec.Command("sudo", "ifconfig", *interfaceName, "up").Run()

    fmt.Println("MAC address changed successfully")
}
