package torrent

import (
	"fmt"
	"io"
)

const (
	Reserved = 8
	HsMsgLen = SHALEN + IDLEN + Reserved
)

type HandshakeMsg struct {
	PreStr  string
	InfoSHA [SHALEN]byte
	PeerId  [IDLEN]byte
}

func NewHandshakeMsg(infoSHA [SHALEN]byte, peerId [IDLEN]byte) *HandshakeMsg {
	return &HandshakeMsg{
		PreStr:  "BitTorrent protocol",
		InfoSHA: infoSHA,
		PeerId:  peerId,
	}
}

func WriteHandshake(w io.Writer, msg *HandshakeMsg) (int, error) {
	buf := make([]byte, len(msg.PreStr)+HsMsgLen+1) // 1 byte for preLen
	buf[0] = byte(len(msg.PreStr))
	curr := 1
	curr += copy(buf[curr:], msg.PreStr)
	curr += copy(buf[curr:], make([]byte, Reserved))
	curr += copy(buf[curr:], msg.InfoSHA[:])
	curr += copy(buf[curr:], msg.PeerId[:])
	return w.Write(buf)
}

func ReadHandshake(r io.Reader) (*HandshakeMsg, error) {
	lenBuf := make([]byte, 1)
	_, err := io.ReadFull(r, lenBuf)
	if err != nil {
		return nil, err
	}
	preLen := int(lenBuf[0])

	if preLen == 0 {
		err := fmt.Errorf("preLen cannot be zero")
		return nil, err
	}

	msgBuf := make([]byte, preLen+HsMsgLen)
	_, err = io.ReadFull(r, msgBuf)
	if err != nil {
		return nil, err
	}

	var peerId [IDLEN]byte
	var infoSHA [SHALEN]byte
	copy(infoSHA[:], msgBuf[preLen+Reserved:preLen+Reserved+SHALEN])
	copy(peerId[:], msgBuf[preLen+Reserved+SHALEN:])

	return &HandshakeMsg{
		PreStr:  string(msgBuf[:preLen]),
		InfoSHA: infoSHA,
		PeerId:  peerId,
	}, nil
}
