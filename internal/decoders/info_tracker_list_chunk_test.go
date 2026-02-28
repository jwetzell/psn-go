package decoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodInfoTrackerListChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.InfoTrackerListChunk
	}{
		{
			description: "InfoTrackerListChunk",
			bytes: []byte{
				2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			expected: chunks.InfoTrackerListChunk{
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
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeInfoTrackerListChunk(testCase.bytes)

			if err != nil {
				t.Errorf("failed to decode chunk properly, error: %v", err)
			}

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to decode chunk properly, expected: %+v, actual: %+v", testCase.expected, actual)
			}
		})
	}
}
