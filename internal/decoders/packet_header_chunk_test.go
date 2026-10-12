package decoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestGoodPacketHeaderChunkDecoding(t *testing.T) {
	testCases := []struct {
		description string
		bytes       []byte
		expected    chunks.PacketHeaderChunk
	}{
		{
			description: "PacketHeaderChunk",
			bytes: []byte{
				0, 0, 12, 0, 210, 2, 150, 73, 0, 0, 0, 0, 2, 3, 1, 123,
			},
			expected: chunks.PacketHeaderChunk{
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
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual, err := DecodePacketHeaderChunk(testCase.bytes)

			require.NoError(t, err)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
