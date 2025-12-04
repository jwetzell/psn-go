package main

import (
	"fmt"
	"log/slog"
	"math"
	"net"
	"time"

	"github.com/jwetzell/psn-go"
)

func main() {
	client, err := net.Dial("udp", "236.10.10.10:56565")
	if err != nil {
		fmt.Printf("Error %v\n", err)
		return
	}

	encoder := psn.Encoder{
		SystemName:  "Test PSN Server",
		VersionHigh: 2,
		VersionLow:  0,
	}

	trackers := []*psn.Tracker{}

	trackers = append(trackers, &psn.Tracker{Id: 0, Name: "Sun"})
	trackers = append(trackers, &psn.Tracker{Id: 1, Name: "Mercury"})
	trackers = append(trackers, &psn.Tracker{Id: 2, Name: "Venus"})
	trackers = append(trackers, &psn.Tracker{Id: 3, Name: "Earth"})
	trackers = append(trackers, &psn.Tracker{Id: 4, Name: "Mars"})
	trackers = append(trackers, &psn.Tracker{Id: 5, Name: "Jupiter"})
	trackers = append(trackers, &psn.Tracker{Id: 6, Name: "Saturn"})
	trackers = append(trackers, &psn.Tracker{Id: 7, Name: "Uranus"})
	trackers = append(trackers, &psn.Tracker{Id: 8, Name: "Neptune"})
	trackers = append(trackers, &psn.Tracker{Id: 9, Name: "Pluto"})

	orbits := []float32{1.0, 88.0, 224.7, 365.2, 687, 4332, 10760, 30700, 60200, 90600}
	distFromSun := []float32{0, 0.58, 1.08, 1.5, 2.28, 7.78, 14.29, 28.71, 45.04, 59.13}

	timestamp := 0

	dataTicker := time.NewTicker(time.Millisecond * 16)
	infoTicker := time.NewTicker(time.Millisecond * 1000)

	go func() {
		for {
			for index, tracker := range trackers {
				orbit := orbits[index]
				a := 1.0 / orbit
				b := distFromSun[index]
				x := timestamp
				cb := math.Cos(float64(a*float32(x))) * float64(b)
				sb := math.Sin(float64(a*float32(x))) * float64(b)

				tracker.SetPos(float32(sb), 0, float32(cb))
				tracker.SetSpeed(a*float32(cb), 0, -a*float32(sb))
				tracker.SetOri(0, float32(x)/1000.0, 0)
				tracker.SetAccel(-a*a*float32(sb), 0, -a*a*float32(cb))
				tracker.SetTrgtPos(3, 14, 16)
				tracker.SetStatus(float32(index) / 10.0)
				tracker.SetTimestamp(uint64(timestamp))
			}
			time.Sleep(time.Millisecond * 16)
		}
	}()

	for {
		select {
		case <-infoTicker.C:
			slog.Info("Sending Info Packets")
			infoPackets := encoder.GetInfoPackets(uint64(timestamp), trackers)
			for _, infoPacket := range infoPackets {
				client.Write(infoPacket)
			}
		case <-dataTicker.C:
			slog.Info("Sending Data Packets")
			dataPackets := encoder.GetDataPackets(uint64(timestamp), trackers)
			for _, DataPacket := range dataPackets {
				client.Write(DataPacket)
			}
			timestamp += 1
		}
	}
}
