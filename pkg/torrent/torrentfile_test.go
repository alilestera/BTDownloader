package torrent_test

import (
	"bufio"
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
	"testing"

	"btdownloader/pkg/torrent"

	"github.com/stretchr/testify/assert"
)

var update = flag.Bool("update", false, "update .golden files")

func TestParseFile(t *testing.T) {
	tests := []struct {
		fileName string
		source   string
	}{
		{
			fileName: "debian-iso.golden",
			source:   "debian-iso.torrent",
		},
	}

	for _, tt := range tests {
		file, err := os.Open(filepath.Join("testdata", tt.source))
		assert.Equal(t, nil, err)
		tf, err := torrent.ParseFile(bufio.NewReader(file))
		assert.Equal(t, nil, err)
		got, err := json.Marshal(tf)
		assert.Equal(t, nil, err, "error in marshaling torrent file struct")

		golden := filepath.Join("testdata", tt.fileName)
		if *update {
			t.Log("update golden file")
			if err := os.WriteFile(golden, got, 0o644); err != nil {
				t.Fatalf("failed to update golden file: %s", err)
			}
		}

		want, err := os.ReadFile(golden)
		assert.Equal(t, nil, err)
		assert.Equal(t, want, got)
	}
}
