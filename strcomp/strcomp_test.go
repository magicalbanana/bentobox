package strcomp

import (
	"testing"

	"github.com/c2fo/testify/require"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		uncompressedStr string
		expected        string
	}{
		{
			uncompressedStr: "abcaaabbb",
			expected:        "abca3b3",
		},
		{
			uncompressedStr: "abcd",
			expected:        "abcd",
		},
		{
			uncompressedStr: "aaabaaaaccaaaaba",
			expected:        "a3ba4c2a4ba",
		},
		{
			uncompressedStr: "jkwqmttnwprlulkgmekyzusvsdmwsech",
			expected:        "jkwqmt2nwprlulkgmekyzusvsdmwsech",
		},
		{
			uncompressedStr: "xrdxuauowrrrrrrrrrrrrrrrrvhbpphaetaaaaxqmhebpvdfulzoojvuhohscnxn",
			expected:        "xrdxuauowr16vhbp2haeta4xqmhebpvdfulzo2jvuhohscnxn",
		},
	}
	for _, test := range tests {
		require.Equal(t, test.expected, Compress(test.uncompressedStr))
	}
}
