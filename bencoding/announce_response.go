package bencoding

import (
	"bytes"
	"github.com/EdmundMartin/discrete/protocol"
)

type torrentState struct {
	numSeeders  int
	numLeechers int
}

func statusCounts(peers []*protocol.Peer) torrentState {
	seeders := 0
	for _, p := range peers {
		if p.Seed {
			seeders += 1
		}
	}
	leechers := len(peers) - seeders
	return torrentState{
		numSeeders:  seeders,
		numLeechers: leechers,
	}
}

func compactV6Peers(peers []*protocol.Peer) []byte {
	var buffer bytes.Buffer
	for _, p := range peers {
		if p.IpV6 {
			buffer.Write(p.Ip.To16())
			buffer.Write([]byte{byte(p.Port >> 8), byte(p.Port & 0xff)})
		}
	}
	return buffer.Bytes()
}

func compactPeers(peers []*protocol.Peer) []byte {
	var buffer bytes.Buffer
	for _, p := range peers {
		if !p.IpV6 {
			buffer.Write(p.Ip.To4())
			buffer.Write([]byte{byte(p.Port >> 8), byte(p.Port & 0xff)})
		}
	}
	return buffer.Bytes()
}

func AnnounceResponse(peers []*protocol.Peer, interval protocol.TrackerInterval, ipv6 bool) ([]byte, error) {
	state := statusCounts(peers)
	dict := Dict{
		"complete":     state.numSeeders,
		"incomplete":   state.numLeechers,
		"min interval": interval.MinIntervalSeconds,
		"interval":     interval.DefaultIntervalSeconds,
	}

	dict["peers"] = compactPeers(peers)
	if ipv6 {
		dict["peers6"] = compactV6Peers(peers)
	}

	var result bytes.Buffer

	if err := NewEncoder(&result).Encode(dict); err != nil {
		return nil, err
	}

	return result.Bytes(), nil
}
