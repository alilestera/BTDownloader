package torrent_test

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"

	"btdownloader/pkg/torrent"
)

func TestPeer(t *testing.T) {
	tests := []struct {
		fileName string
		peer     torrent.PeerInfo
	}{
		{
			fileName: "debian-iso.torrent",
			peer: torrent.PeerInfo{
				Ip:   net.ParseIP("84.44.131.89"),
				Port: uint16(50000),
			},
		},
	}

	for _, tt := range tests {
		file, err := os.Open(filepath.Join("testdata", tt.fileName))
		if err != nil {
			t.Errorf("error: Open file [%s] failed, got %s", tt.fileName, err.Error())
			continue
		}

		var peerId [torrent.IDLEN]byte
		_, _ = rand.Read(peerId[:])
		tf, err := torrent.ParseFile(bufio.NewReader(file))
		if err != nil {
			t.Errorf("error: Parse file [%s] failed, got %s", tt.fileName, err.Error())
			continue
		}
		conn, err := torrent.NewConn(tt.peer, tf.InfoSHA, peerId)
		if err != nil {
			t.Errorf("error: New peer with file [%s] failed, got %s", tt.fileName, err.Error())
		}
		fmt.Println(conn)
	}
}
