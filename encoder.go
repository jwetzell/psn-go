package psn

import "github.com/jwetzell/psn-go/internal/encoders"

type Encoder struct {
	SystemName  string
	VersionHigh uint8
	VersionLow  uint8
	dataFrameId uint8
	infoFrameId uint8
}

func (e *Encoder) ResetDataFrameId() {
	e.dataFrameId = 1
}

func (e *Encoder) ResetInfoFrameId() {
	e.infoFrameId = 1
}

func (e *Encoder) GetInfoPackets(timestamp uint64, trackers []Tracker) [][]byte {
	systemNameChunk := encoders.EncodeInfoSystemNameChunk(e.SystemName)

	trackerChunks := [][]byte{}

	for _, tracker := range trackers {
		trackerChunks = append(trackerChunks, tracker.GetInfoChunk())
	}

	infoPackets := [][]byte{}

	trackerChunksLists := [][][]byte{}
	currentTrackerList := [][]byte{}

	currentInfoPacketSize := PACKET_HEADER_SIZE + len(systemNameChunk) + CHUNK_HEADER_SIZE

	for _, trackerChunk := range trackerChunks {
		if (currentInfoPacketSize + len(trackerChunk)) > MAX_UDP_PACKET_SIZE {
			trackerChunksLists = append(trackerChunksLists, currentTrackerList)
			currentTrackerList = [][]byte{}
			currentInfoPacketSize = 0
		}
		currentTrackerList = append(currentTrackerList, trackerChunk)
		currentInfoPacketSize += len(trackerChunk)
	}
	trackerChunksLists = append(trackerChunksLists, currentTrackerList)
	header := encoders.EncodePacketHeaderChunk(timestamp, e.VersionHigh, e.VersionLow, e.infoFrameId, uint8(len(trackerChunksLists)))

	for _, trackerChunkList := range trackerChunksLists {
		infoPackets = append(infoPackets, encoders.EncodeInfoPacketChunk(header, systemNameChunk, encoders.EncodeInfoTrackerListChunk(trackerChunkList)))
	}

	if e.infoFrameId == 255 {
		e.infoFrameId = 0
	} else {
		e.infoFrameId += 1
	}
	return infoPackets
}

func (e *Encoder) GetDataPackets(timestamp uint64, trackers []Tracker) [][]byte {

	trackerChunks := [][]byte{}

	for _, tracker := range trackers {
		trackerChunks = append(trackerChunks, tracker.GetDataChunk())
	}

	dataPackets := [][]byte{}

	trackerChunksLists := [][][]byte{}
	currentTrackerList := [][]byte{}

	currentDataPacketSize := PACKET_HEADER_SIZE + CHUNK_HEADER_SIZE

	for _, trackerChunk := range trackerChunks {
		if (currentDataPacketSize + len(trackerChunk)) > MAX_UDP_PACKET_SIZE {
			trackerChunksLists = append(trackerChunksLists, currentTrackerList)
			currentTrackerList = [][]byte{}
			currentDataPacketSize = 0
		}
		currentTrackerList = append(currentTrackerList, trackerChunk)
		currentDataPacketSize += len(trackerChunk)
	}
	trackerChunksLists = append(trackerChunksLists, currentTrackerList)
	header := encoders.EncodePacketHeaderChunk(timestamp, e.VersionHigh, e.VersionLow, e.dataFrameId, uint8(len(trackerChunksLists)))

	for _, trackerChunkList := range trackerChunksLists {
		dataPackets = append(dataPackets, encoders.EncodeDataPacketChunk(header, encoders.EncodeDataTrackerListChunk(trackerChunkList)))
	}

	if e.dataFrameId == 255 {
		e.dataFrameId = 0
	} else {
		e.dataFrameId += 1
	}

	return dataPackets
}
