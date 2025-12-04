package main

import (
	"fmt"
	"time"

	"github.com/jwetzell/psn-go"
)

func getNTrackers(n int) []*psn.Tracker {
	trackers := []*psn.Tracker{}

	for index := 0; index < n; index++ {
		tracker := &psn.Tracker{Id: uint16(index), Name: "Tracker"}
		tracker.SetPos(0, 0, 0)
		tracker.SetSpeed(0, 0, 0)
		tracker.SetOri(0, 0, 0)
		tracker.SetAccel(0, 0, 0)
		tracker.SetTrgtPos(0, 0, 0)
		tracker.SetStatus(1.0)
		tracker.SetTimestamp(uint64(time.Now().UnixMilli()))
		trackers = append(trackers, tracker)
	}
	return trackers
}

func main() {
	testSizes := []int{1, 10, 100, 1000}
	encoder := psn.Encoder{
		SystemName:  "test encoder",
		VersionHigh: 2,
		VersionLow:  3,
	}
	decoder := psn.NewDecoder()

	for _, iterations := range testSizes {
		for _, trackerCount := range testSizes {
			benchmark(trackerCount, iterations, encoder, *decoder)
		}
	}
}

type BenchmarkResults struct {
	data DataBenchmarkResults
	info DataBenchmarkResults
}

type DataBenchmarkResults struct {
	encode float64
	decode float64
}

func benchmark(trackerCount int, iterations int, encoder psn.Encoder, decoder psn.Decoder) {
	fmt.Printf("processing %d trackers %d times\n", trackerCount, iterations)
	benchmarkResults := BenchmarkResults{}
	trackers := getNTrackers(trackerCount)

	timestamp := time.Now().UnixMilli()
	// DATA
	dataEncoderStart := time.Now().UnixMicro()

	latestEncodedPackets := [][]byte{}

	for index := 0; index < iterations; index++ {
		latestEncodedPackets = encoder.GetDataPackets(uint64(timestamp), trackers)
	}
	benchmarkResults.data.encode = float64(time.Now().UnixMicro()-dataEncoderStart) / 1000.0

	_ = latestEncodedPackets
	dataDecodedStart := time.Now().UnixMicro()
	for index := 0; index < iterations; index++ {
		for _, packet := range latestEncodedPackets {
			decoder.Decode(packet)
		}
	}
	benchmarkResults.data.decode = float64(time.Now().UnixMicro()-dataDecodedStart) / 1000.0

	// INFO
	infoEncoderStart := time.Now().UnixMicro()
	for index := 0; index < iterations; index++ {
		latestEncodedPackets = encoder.GetInfoPackets(uint64(timestamp), trackers)
	}
	benchmarkResults.info.encode = float64(time.Now().UnixMicro()-infoEncoderStart) / 1000.0

	infoDecodeStart := time.Now().UnixMicro()
	for index := 0; index < iterations; index++ {
		for _, packet := range latestEncodedPackets {
			decoder.Decode(packet)
		}
	}
	benchmarkResults.info.decode = float64(time.Now().UnixMicro()-infoDecodeStart) / 1000.0

	fmt.Printf("%+v\n", benchmarkResults)
}
