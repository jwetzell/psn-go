package psn

import (
	"github.com/jwetzell/psn-go/internal/decoders"
)

type Decoder struct {
	lastInfoPacketHeader *decoders.PacketHeaderChunk
	lastDataPacketHeader *decoders.PacketHeaderChunk
	infoPacketFrames     map[uint8][]decoders.InfoPacketChunk
	dataPacketFrames     map[uint8][]decoders.DataPacketChunk
	Trackers             map[uint16]*Tracker
	SystemName           string
}

func NewDecoder() *Decoder {
	var decoder Decoder
	decoder.infoPacketFrames = map[uint8][]decoders.InfoPacketChunk{}
	decoder.dataPacketFrames = map[uint8][]decoders.DataPacketChunk{}
	decoder.Trackers = map[uint16]*Tracker{}
	return &decoder
}

func (d *Decoder) updateInfo(framePackets []decoders.InfoPacketChunk) {
	for _, framePacket := range framePackets {
		d.SystemName = framePacket.Data.SystemName.Data.SystemName
		for _, infoTrackerChunk := range framePacket.Data.TrackerList.Data.Trackers {
			tracker, ok := d.Trackers[infoTrackerChunk.Chunk.Header.Id]

			if ok {
				tracker.UpdateInfo(infoTrackerChunk)
			} else {
				d.Trackers[infoTrackerChunk.Chunk.Header.Id] = TrackerFromInfo(infoTrackerChunk)
			}
		}
	}
}

func (d *Decoder) updateData(framePackets []decoders.DataPacketChunk) {
	for _, framePacket := range framePackets {
		for _, dataTrackerChunk := range framePacket.Data.TrackerList.Data.Trackers {
			tracker, ok := d.Trackers[dataTrackerChunk.Chunk.Header.Id]

			if ok {
				tracker.UpdateData(dataTrackerChunk)
			} else {
				d.Trackers[dataTrackerChunk.Chunk.Header.Id] = TrackerFromData(dataTrackerChunk)
			}
		}
	}
}

func (d *Decoder) Decode(bytes []byte) error {
	chunk, err := decoders.DecodeChunk(bytes)

	if err != nil {
		return err
	}

	if chunk.Header.Id == 0x6756 {
		infoPacket, err := decoders.DecodeInfoPacketChunk(bytes)
		if err != nil {
			return err
		}
		currentInfoPacketHeader := infoPacket.Data.PacketHeader

		_, ok := d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId]

		if !ok {
			d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId] = []decoders.InfoPacketChunk{}
		}
		d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId] = append(d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId], infoPacket)

		if len(d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId]) == int(currentInfoPacketHeader.Data.FramePacketCount) {
			d.updateInfo(d.infoPacketFrames[currentInfoPacketHeader.Data.FrameId])
			delete(d.infoPacketFrames, currentInfoPacketHeader.Data.FrameId)
		}
	} else if chunk.Header.Id == 0x6755 {
		dataPacket, err := decoders.DecodeDataPacketChunk(bytes)
		if err != nil {
			return err
		}
		currentInfoPacketHeader := dataPacket.Data.PacketHeader

		_, ok := d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId]

		if !ok {
			d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId] = []decoders.DataPacketChunk{}
		}
		d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId] = append(d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId], dataPacket)

		if len(d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId]) == int(currentInfoPacketHeader.Data.FramePacketCount) {
			d.updateData(d.dataPacketFrames[currentInfoPacketHeader.Data.FrameId])
			delete(d.dataPacketFrames, currentInfoPacketHeader.Data.FrameId)
		}
	}
	return nil
}
