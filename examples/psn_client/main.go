package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jwetzell/psn-go"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "236.10.10.10:56565")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	client, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}
	defer client.Close()

	decoder := psn.NewDecoder()

	lastPrintMillis := time.Now().UnixMilli()

	for {
		buffer := make([]byte, 2048)

		length, _, err := client.ReadFromUDP(buffer)
		if err != nil {
			fmt.Printf("Error %v", err)
		}

		if length > 0 {
			decoder.Decode(buffer)
		}

		if time.Now().UnixMilli()-lastPrintMillis > 1000 {
			if decoder.SystemName != "" {
				fmt.Printf("System Name: %s\n", decoder.SystemName)
			}

			fmt.Printf("Tracker Count: %d\n", len(decoder.Trackers))

			for id, tracker := range decoder.Trackers {
				fmt.Printf("Tracker - id: %d | name: %s\n", id, tracker.Name)
				if tracker.Pos != nil {
					fmt.Printf("\tpos: %f,%f,%f\n", tracker.Pos.X, tracker.Pos.Y, tracker.Pos.Z)
				}

				if tracker.Speed != nil {
					fmt.Printf("\tspeed: %f,%f,%f\n", tracker.Speed.X, tracker.Speed.Y, tracker.Speed.Z)
				}

				if tracker.Ori != nil {
					fmt.Printf("\tori: %f,%f,%f\n", tracker.Ori.X, tracker.Ori.Y, tracker.Ori.Z)
				}

				if tracker.Validity != nil {
					fmt.Printf("\tstatus: %f\n", *tracker.Validity)
				}

				if tracker.Accel != nil {
					fmt.Printf("\taccel: %f,%f,%f\n", tracker.Accel.X, tracker.Accel.Y, tracker.Accel.Z)
				}

				if tracker.TrgtPos != nil {
					fmt.Printf("\ttrgtpos: %f,%f,%f\n", tracker.TrgtPos.X, tracker.TrgtPos.Y, tracker.TrgtPos.Z)
				}

				if tracker.Timestamp != nil {
					fmt.Printf("\ttimestamp: %d\n", *tracker.Timestamp)
				}
			}
			lastPrintMillis = time.Now().UnixMilli()
		}
	}
}
