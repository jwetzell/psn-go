package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestGoodDataTrackerXYZChunk(t *testing.T) {

	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.DataTrackerXYZChunk
	}{
		{
			description: "DataTrackerPosChunk",
			bytes:       []byte{0, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: chunks.DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerSpeedChunk",
			bytes:       []byte{1, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: chunks.DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerOriChunk",
			bytes:       []byte{2, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: chunks.DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerAccelChunk",
			bytes:       []byte{4, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: chunks.DataTrackerXYZChunk{
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
		},
		{
			description: "DataTrackerTrgtPosChunk",
			bytes:       []byte{5, 0, 12, 0, 0, 0, 128, 63, 0, 0, 0, 64, 0, 0, 64, 64},
			expected: chunks.DataTrackerXYZChunk{
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
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodeDataTrackerXYZChunk(testCase.bytes)

			require.NoError(t, err)

			require.Equal(t, testCase.expected, actual)
		})
	}

}
