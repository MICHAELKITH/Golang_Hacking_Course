package main

import (
    "fmt"
    "github.com/google/gopacket/pcap"
)

func main() {
    devices, err := pcap.FindAllDevs()
    if err != nil {
        fmt.Println("Error finding devices:", err)
        return
    }

    for _, device := range devices {
        fmt.Println("Device one:", device.Name)
        fmt.Println("Device Description:", device.Description)
    }
}
