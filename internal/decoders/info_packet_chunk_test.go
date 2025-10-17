package decoders

import (
	"fmt"
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
	}

	for _, testCase := range testCases {

		actual, err := DecodeInfoPacketChunk(testCase.bytes)

		if err != nil {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Println(err)
		}

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to decode chunk properly", testCase.description)
			fmt.Printf("expected: %+v\n", testCase.expected)
			fmt.Printf("actual: %+v\n", actual)
		}
	}
}
