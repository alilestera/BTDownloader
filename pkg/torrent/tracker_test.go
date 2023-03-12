package torrent_test

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"btdownloader/pkg/torrent"

	"github.com/stretchr/testify/assert"
)

func TestFindPeers(t *testing.T) {
	file, err := os.Open(filepath.Join("testdata", "debian-iso.torrent"))
	assert.Equal(t, nil, err)
	tf, err := torrent.ParseFile(bufio.NewReader(file))
	assert.Equal(t, nil, err)

	var peerId [torrent.IDLEN]byte
	_, _ = rand.Read(peerId[:])

	peers := torrent.FindPeers(tf, peerId)
	for i, p := range peers {
		fmt.Printf("Peer %d, Ip: %s, Port: %d\n", i, p.Ip, p.Port)
	}
}
