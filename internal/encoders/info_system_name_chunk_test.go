package encoders

import (
	"testing"

	"github.com/jwetzell/psn-go/internal/chunks"
	"github.com/stretchr/testify/require"
)

func TestInfoSystemNameChunkEncoding(t *testing.T) {
	testCases := []struct {
		description string
		expected    []byte
		data        chunks.InfoSystemNameChunkData
	}{
		{
			description: "InfoSystemNameChunk",
			expected: []byte{
				1, 0, 10, 0, 80, 83, 78, 32, 83, 101, 114, 118, 101, 114,
			},
			data: chunks.InfoSystemNameChunkData{
				SystemName: "PSN Server",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := EncodeInfoSystemNameChunk(testCase.data.SystemName)

			require.Equal(t, testCase.expected, actual)
		})
	}
}
