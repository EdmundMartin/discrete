package protocol

var activeEvents = map[string]interface{}{
	"completed": nil,
	"started":   nil,
	"":          nil, // Represents a new announce
}

func remove(p []*Peer, idx int) []*Peer {
	p[idx] = p[len(p)-1]
	return p[:len(p)-1]
}

func ActivePeers(peers []*Peer) []*Peer {
	var active []*Peer
	for _, p := range peers {
		_, ok := activeEvents[p.Event]
		if ok {
			active = append(active)
		}
	}
	return active
}

func FilterPeers(peers []*Peer, clientID string) []*Peer {
	peers = ActivePeers(peers)
	idx := -1
	for i, p := range peers {
		if p.ClientID == clientID {
			idx = i
		}
	}
	if idx >= 0 {
		remove(peers, idx)
	}
	return peers
}