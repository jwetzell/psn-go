package main

import (
	"fmt"
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

	trackers := []psn.Tracker{}

	trackers = append(trackers, psn.Tracker{Id: 0, Name: "Sun"})
	trackers = append(trackers, psn.Tracker{Id: 1, Name: "Mercury"})
	trackers = append(trackers, psn.Tracker{Id: 2, Name: "Venus"})
	trackers = append(trackers, psn.Tracker{Id: 3, Name: "Earth"})
	trackers = append(trackers, psn.Tracker{Id: 4, Name: "Mars"})
	trackers = append(trackers, psn.Tracker{Id: 5, Name: "Jupiter"})
	trackers = append(trackers, psn.Tracker{Id: 6, Name: "Saturn"})
	trackers = append(trackers, psn.Tracker{Id: 7, Name: "Uranus"})
	trackers = append(trackers, psn.Tracker{Id: 8, Name: "Neptune"})
	trackers = append(trackers, psn.Tracker{Id: 9, Name: "Pluto"})

	orbits := []float32{1.0, 88.0, 224.7, 365.2, 687, 4332, 10760, 30700, 60200, 90600}
	distFromSun := []float32{0, 0.58, 1.08, 1.5, 2.28, 7.78, 14.29, 28.71, 45.04, 59.13}

	timestamp := 0

	lastInfoMillis := time.Now().UnixMilli()
	lastDataMillis := time.Now().UnixMilli()

	for {
		if (time.Now().UnixMilli() - lastInfoMillis) > 500 {
			fmt.Println("Sending Info Packets")
			infoPackets := encoder.GetInfoPackets(uint64(timestamp), trackers)
			for _, infoPacket := range infoPackets {
				client.Write(infoPacket)
			}
			lastInfoMillis = time.Now().UnixMilli()
		}

		if (time.Now().UnixMilli() - lastDataMillis) > 5 {
			for index, orbit := range orbits {
				a := 1.0 / orbit
				b := distFromSun[index]
				x := timestamp
				cb := math.Cos(float64(a*float32(x))) * float64(b)
				sb := math.Sin(float64(a*float32(x))) * float64(b)

				trackers[index].SetPos(float32(sb), 0, float32(cb))
				trackers[index].SetSpeed(a*float32(cb), 0, -a*float32(sb))
				trackers[index].SetOri(0, float32(x)/1000.0, 0)
				trackers[index].SetAccel(-a*a*float32(sb), 0, -a*a*float32(cb))
				trackers[index].SetTrgtPos(3, 14, 16)
				trackers[index].SetStatus(float32(index) / 10.0)
				trackers[index].SetTimestamp(uint64(timestamp))
			}

			fmt.Println("Sending Data Packets")
			dataPackets := encoder.GetDataPackets(uint64(timestamp), trackers)
			for _, DataPacket := range dataPackets {
				client.Write(DataPacket)
			}
			lastDataMillis = time.Now().UnixMilli()
			timestamp += 1
		}

	}

}
