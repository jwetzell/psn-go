package main

import (
	"fmt"

	"github.com/jwetzell/psn-go"
)

func main() {
	encoder := psn.Encoder{
		SystemName:  "Server Name",
		VersionHigh: 2,
		VersionLow:  3,
	}

	tracker := &psn.Tracker{
		Id:   1,
		Name: "Tracker 1",
	}

	tracker.SetPos(1.0, 1.0, 1.0)

	trackers := []*psn.Tracker{
		tracker,
	}

	timestamp := 1

	dataPackets := encoder.GetDataPackets(uint64(timestamp), trackers)
	infoPackets := encoder.GetInfoPackets(uint64(timestamp), trackers)

	for _, dataPacket := range dataPackets {
		println("send packet somehow")
		fmt.Printf("%v\n", dataPacket)

	}

	for _, infoPacket := range infoPackets {
		println("send packet somehow")
		fmt.Printf("%v\n", infoPacket)

	}
}
