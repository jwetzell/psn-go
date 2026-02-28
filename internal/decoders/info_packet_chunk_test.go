package decoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodInfoPacketChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.InfoPacketChunk
	}{
		{
			description: "InfoPacketChunk",
			bytes: []byte{
				86, 103, 51, 128, 0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 10, 0, 80, 83, 78, 32, 83, 101,
				114, 118, 101, 114, 2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: chunks.InfoPacketChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{
						0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101,
						114, 2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
					},
					Header: chunks.ChunkHeader{DataLen: 51, Id: 26454, HasSubchunks: true},
				},
				Data: chunks.InfoPacketChunkData{
					PacketHeader: &chunks.PacketHeaderChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
						},
						Data: chunks.PacketHeaderChunkData{
							PacketTimestamp:  1234567890,
							VersionHigh:      2,
							VersionLow:       3,
							FrameId:          1,
							FramePacketCount: 123,
						},
					},
					SystemName: &chunks.InfoSystemNameChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{80, 83, 78, 32, 83, 101, 114, 118, 101, 114},
							Header:    chunks.ChunkHeader{DataLen: 10, Id: 1, HasSubchunks: false},
						},
						Data: chunks.InfoSystemNameChunkData{
							SystemName: "PSN Server",
						},
					},
					TrackerList: &chunks.InfoTrackerListChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
							Header:    chunks.ChunkHeader{DataLen: 17, Id: 2, HasSubchunks: true},
						},
						Data: chunks.InfoTrackerListChunkData{
							Trackers: []chunks.InfoTrackerChunk{
								{
									Chunk: chunks.Chunk{
										ChunkData: []byte{0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
										Header:    chunks.ChunkHeader{DataLen: 13, Id: 1, HasSubchunks: true},
									},
									Data: chunks.InfoTrackerChunkData{
										TrackerName: &chunks.InfoTrackerNameChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
												Header:    chunks.ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
											},
											Data: chunks.InfoTrackerNameChunkData{
												TrackerName: "Tracker 1",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		{
			description: "empty InfoPacketChunk",
			bytes: []byte{
				0x56, 0x67, 0x22, 0x80, 0x00, 0x00, 0x0C, 0x00, 0x3D, 0x2C, 0x91, 0xDB, 0x9A, 0x01, 0x00, 0x00, 0x02, 0x00, 0x00,
				0x01, 0x01, 0x00, 0x0A, 0x00, 0x50, 0x53, 0x4E, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x02, 0x00, 0x00, 0x80,
			},
			expected: chunks.InfoPacketChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{
						0x00, 0x00, 0x0C, 0x00, 0x3D, 0x2C, 0x91, 0xDB, 0x9A, 0x01, 0x00, 0x00, 0x02, 0x00, 0x00,
						0x01, 0x01, 0x00, 0x0A, 0x00, 0x50, 0x53, 0x4E, 0x20, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x02, 0x00, 0x00, 0x80,
					},
					Header: chunks.ChunkHeader{DataLen: 34, Id: 26454, HasSubchunks: true},
				},
				Data: chunks.InfoPacketChunkData{
					PacketHeader: &chunks.PacketHeaderChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{0x3D, 0x2C, 0x91, 0xDB, 0x9A, 0x01, 0x00, 0x00, 0x02, 0x00, 0x00, 0x01},
							Header:    chunks.ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
						},
						Data: chunks.PacketHeaderChunkData{
							PacketTimestamp:  1764620315709,
							VersionHigh:      2,
							VersionLow:       0,
							FrameId:          0,
							FramePacketCount: 1,
						},
					},
					SystemName: &chunks.InfoSystemNameChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{80, 83, 78, 32, 83, 101, 114, 118, 101, 114},
							Header:    chunks.ChunkHeader{DataLen: 10, Id: 1, HasSubchunks: false},
						},
						Data: chunks.InfoSystemNameChunkData{
							SystemName: "PSN Server",
						},
					},
					TrackerList: &chunks.InfoTrackerListChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{},
							Header:    chunks.ChunkHeader{DataLen: 0, Id: 2, HasSubchunks: true},
						},
						Data: chunks.InfoTrackerListChunkData{
							Trackers: []chunks.InfoTrackerChunk{},
						},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeInfoPacketChunk(testCase.bytes)

			if err != nil {
				t.Errorf("failed to decode chunk properly, error: %v", err)
			}

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to decode chunk properly, expected: %+v, actual: %+v", testCase.expected, actual)
			}
		})
	}
}
