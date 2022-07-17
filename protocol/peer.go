package protocol

import (
	"fmt"
	"github.com/EdmundMartin/discrete/torrent_errors"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	infoHashField   = "info_hash"
	portField       = "port"
	peerIdField     = "peer_id"
	uploadedField   = "uploaded"
	downloadedField = "downloaded"
	leftField       = "left"
	event           = "event"
)

var optionalInfo = []string{
	uploadedField,
	downloadedField,
	leftField,
	event,
}

type Peer struct {
	InfoHash   string
	ClientID   string
	Ip         net.IP
	IpV6       bool
	Port       int
	Event      Event
	Seen       time.Time
	Uploaded   int
	Downloaded int
	Seed       bool
}

func (p Peer) String() string {
	return fmt.Sprintf("InfoHash: %s, ClientID: %s", p.InfoHash, p.ClientID)
}

func isSeed(left string) bool {
	seed := false
	num, err := strconv.Atoi(left)
	if err != nil && num == 0 {
		seed = true
	}
	return seed
}

func forceToInt(value string) int {
	val, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return val
}

func parseInfoHash(hash string) string {
	return fmt.Sprintf("%x", hash)
}

func parseRemoteAddr(r *http.Request) string {
	return strings.Split(r.RemoteAddr, ":")[0]
}

func extractRequiredFields(peer *Peer, values map[string][]string) error {

	val, ok := values[infoHashField]
	if !ok {
		return torrent_errors.MissingInfoHash
	}
	peer.InfoHash = parseInfoHash(val[0])

	val, ok = values[portField]
	if !ok {
		return torrent_errors.MissingPort
	}
	portVal := forceToInt(val[0])
	// Do not allow connections form protected ports
	if portVal <= 1024 {
		return torrent_errors.InvalidPort
	}
	peer.Port = portVal

	val, ok = values[peerIdField]
	if !ok {
		return torrent_errors.MissingPeerID
	}
	peerID := val[0]
	if len(peerID) != 20 {
		return torrent_errors.InvalidPeerID
	}

	peer.ClientID = peerID

	return nil
}

func extractOptionalFields(peer *Peer, values map[string][]string) {
	optionalFields := make(map[string]string)

	for _, opt := range optionalInfo {
		val, ok := values[opt]
		if !ok {
			optionalFields[opt] = ""
		}
		if len(val) > 0 {
			// Event can be an empty array within request header - identifying announce
			optionalFields[opt] = val[0]
		} else {
			optionalFields[opt] = ""
		}
	}
	peer.Uploaded = forceToInt(optionalFields[uploadedField])
	peer.Downloaded = forceToInt(optionalFields[downloadedField])
	peer.Event = EventFromString(optionalFields[event])
	peer.Seed = isSeed(optionalFields[leftField])
}

func NewPeerFromRequest(r *http.Request) (*Peer, error) {

	peer := &Peer{}
	query := r.URL.Query()

	clientIP := GetIP(r)
	peer.Ip = clientIP.IP
	peer.IpV6 = clientIP.IpV6

	err := extractRequiredFields(peer, query)
	if err != nil {
		return nil, err
	}
	extractOptionalFields(peer, query)
	peer.Seen = time.Now()
	return peer, err
}
