package decoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodDataPacketChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataPacketChunk
	}{
		{
			description: "DataPacketChunk",
			bytes: []byte{
				85, 103, 124, 128, 0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 104, 128, 1, 0, 100, 128, 0, 0,
				12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 2, 0, 12,
				0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64,
				0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
			},
			expected: chunks.DataPacketChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{
						0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123, 1, 0, 104, 128, 1, 0, 100, 128, 0, 0, 12, 0, 0, 0,
						128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 2, 0, 12, 0, 0, 0,
						128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0,
						64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
					},
					Header: chunks.ChunkHeader{
						DataLen: 124, Id: 26453, HasSubchunks: true,
					},
				},
				Data: chunks.DataPacketChunkData{
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
					TrackerList: &chunks.DataTrackerListChunk{
						Chunk: chunks.Chunk{
							ChunkData: []byte{
								1, 0, 100, 128, 0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0,
								0, 64, 0, 0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4,
								0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64,
								6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
							},
							Header: chunks.ChunkHeader{DataLen: 104, Id: 1, HasSubchunks: true},
						},
						Data: chunks.DataTrackerListChunkData{
							Trackers: []chunks.DataTrackerChunk{
								chunks.DataTrackerChunk{
									Chunk: chunks.Chunk{
										ChunkData: []byte{
											0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0,
											0, 64, 64, 2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 3, 0, 4, 0, 0, 0, 128, 63, 4, 0,
											12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64, 5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64,
											64, 6, 0, 8, 0, 210, 2, 150, 73, 0, 0, 0, 0,
										},
										Header: chunks.ChunkHeader{DataLen: 100, Id: 1, HasSubchunks: true},
									},
									Data: chunks.DataTrackerChunkData{
										Pos: &chunks.DataTrackerXYZChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
												Header:    chunks.ChunkHeader{DataLen: 12, Id: 0, HasSubchunks: false},
											},
											Data: chunks.DataTrackerXYZChunkData{
												X: 1.0,
												Y: 2.0,
												Z: 3.0,
											},
										},
										Speed: &chunks.DataTrackerXYZChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
												Header:    chunks.ChunkHeader{DataLen: 12, Id: 1, HasSubchunks: false},
											},
											Data: chunks.DataTrackerXYZChunkData{
												X: 1.0,
												Y: 2.0,
												Z: 3.0,
											},
										},
										Ori: &chunks.DataTrackerXYZChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
												Header:    chunks.ChunkHeader{DataLen: 12, Id: 2, HasSubchunks: false},
											},
											Data: chunks.DataTrackerXYZChunkData{
												X: 1.0,
												Y: 2.0,
												Z: 3.0,
											},
										},
										Status: &chunks.DataTrackerStatusChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63},
												Header:    chunks.ChunkHeader{DataLen: 4, Id: 3, HasSubchunks: false},
											},
											Data: chunks.DataTrackerStatusChunkData{
												Validity: 1.0,
											},
										},
										Accel: &chunks.DataTrackerXYZChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
												Header:    chunks.ChunkHeader{DataLen: 12, Id: 4, HasSubchunks: false},
											},
											Data: chunks.DataTrackerXYZChunkData{
												X: 1.0,
												Y: 2.0,
												Z: 3.0,
											},
										},
										TrgtPos: &chunks.DataTrackerXYZChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
												Header:    chunks.ChunkHeader{DataLen: 12, Id: 5, HasSubchunks: false},
											},
											Data: chunks.DataTrackerXYZChunkData{
												X: 1.0,
												Y: 2.0,
												Z: 3.0,
											},
										},
										Timestamp: &chunks.DataTrackerTimestampChunk{
											Chunk: chunks.Chunk{
												ChunkData: []byte{210, 2, 150, 73, 0, 0, 0, 0},
												Header:    chunks.ChunkHeader{DataLen: 8, Id: 6, HasSubchunks: false},
											},
											Data: chunks.DataTrackerTimestampChunkData{
												Timestamp: 1234567890,
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

		actual, err := DecodeDataPacketChunk(testCase.bytes)

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
