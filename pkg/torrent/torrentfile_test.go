package torrent_test

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"btdownloader/pkg/torrent"

	"github.com/stretchr/testify/assert"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		source  string
		torFile torrent.TorrentFile
	}{
		{
			source: "debian-iso.torrent",
			torFile: torrent.TorrentFile{
				Announce: "http://bttracker.debian.org:6969/announce",
				FileName: "debian-11.2.0-amd64-netinst.iso",
				FileLen:  396361728,
				PieceLen: 262144,
				InfoSHA: [20]byte{
					0x28, 0xc5, 0x51, 0x96, 0xf5, 0x77, 0x53, 0xc4, 0xa,
					0xce, 0xb6, 0xfb, 0x58, 0x61, 0x7e, 0x69, 0x95, 0xa7, 0xed, 0xdb,
				},
			},
		},
	}

	for _, tt := range tests {
		file, err := os.Open(filepath.Join("testdata", tt.source))
		assert.Equal(t, nil, err)
		tf, err := torrent.ParseFile(bufio.NewReader(file))
		assert.Equal(t, nil, err)
		tf.PieceSHA = nil
		assert.NotEqualValues(t, tt.torFile, tf, "")
	}
}
