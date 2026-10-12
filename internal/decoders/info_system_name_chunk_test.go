package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
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

			require.NoError(t, err)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
