package decoders

import (
	"reflect"
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
)

func TestGoodInfoSystemNameChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.InfoSystemNameChunk
	}{
		{
			description: "InfoSystemNameChunk",
			bytes: []byte{
				1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101, 114,
			},
			expected: chunks.InfoSystemNameChunk{
				Chunk: chunks.Chunk{
					ChunkData: []byte{80, 83, 78, 32, 83, 101, 114, 118, 101, 114},
					Header:    chunks.ChunkHeader{DataLen: 10, Id: 1, HasSubchunks: false},
				},
				Data: chunks.InfoSystemNameChunkData{
					SystemName: "PSN Server",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeInfoSystemNameChunk(testCase.bytes)

			if err != nil {
				t.Errorf("failed to decode chunk properly, error: %v", err)
			}

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("failed to decode chunk properly, expected: %+v, actual: %+v", testCase.expected, actual)
			}
		})
	}
}
