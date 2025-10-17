package decoders

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGoodInfoPacketChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    InfoPacketChunk
	}{
		{
			description: "InfoPacketChunk",
			bytes: []byte{
				86, 103, 51, 128, 0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 10, 0, 80, 83, 78, 32, 83, 101,
				114, 118, 101, 114, 2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: InfoPacketChunk{
				Chunk: Chunk{
					ChunkData: []byte{
						0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101,
						114, 2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
					},
					Header: ChunkHeader{DataLen: 51, Id: 26454, HasSubchunks: true},
				},
				Data: InfoPacketChunkData{
					PacketHeader: &PacketHeaderChunk{
						Chunk: Chunk{
							ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123},
							Header:    ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
						},
						Data: PacketHeaderChunkData{
							PacketTimestamp:  1234567890,
							VersionHigh:      2,
							VersionLow:       3,
							FrameId:          1,
							FramePacketCount: 123,
						},
					},
					SystemName: &InfoSystemNameChunk{
						Chunk: Chunk{
							ChunkData: []byte{80, 83, 78, 32, 83, 101, 114, 118, 101, 114},
							Header:    ChunkHeader{DataLen: 10, Id: 1, HasSubchunks: false},
						},
						Data: InfoSystemNameChunkData{
							SystemName: "PSN Server",
						},
					},
					TrackerList: &InfoTrackerListChunk{
						Chunk: Chunk{
							ChunkData: []byte{1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
							Header:    ChunkHeader{DataLen: 17, Id: 2, HasSubchunks: true},
						},
						Data: InfoTrackerListChunkData{
							Trackers: []InfoTrackerChunk{
								{
									Chunk: Chunk{
										ChunkData: []byte{0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49},
										Header:    ChunkHeader{DataLen: 13, Id: 1, HasSubchunks: true},
									},
									Data: InfoTrackerChunkData{
										TrackerName: &InfoTrackerNameChunk{
											Chunk: Chunk{
												ChunkData: []byte{84, 114, 97, 99, 107, 101, 114, 32, 49},
												Header:    ChunkHeader{DataLen: 9, Id: 0, HasSubchunks: false},
											},
											Data: InfoTrackerNameChunkData{
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
