package encoders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestInfoTrackerListChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		chunk       chunks.InfoTrackerListChunk
	}{
		{
			description: "InfoTrackerListChunk",
			expected: []byte{
				2, 0, 17, 128, 1, 0, 13, 128, 0, 0, 9, 0, 84, 114, 97, 99, 107, 101, 114, 32, 49,
			},
			chunk: chunks.InfoTrackerListChunk{
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

		trackerChunks := [][]byte{}
		for _, tracker := range testCase.chunk.Data.Trackers {
			trackerChunks = append(trackerChunks, EncodeInfoTrackerChunk(tracker.Chunk.Header.Id, EncodeInfoTrackerNameChunk(tracker.Data.TrackerName.Data.TrackerName)))
		}

		actual := EncodeInfoTrackerListChunk(trackerChunks)

		if !reflect.DeepEqual(actual, testCase.expected) {
			t.Errorf("Test '%s' failed to encode chunk properly", testCase.description)
			fmt.Printf("expected: %v\n", testCase.expected)
			fmt.Printf("actual: %v\n", actual)
		}
	}
}
