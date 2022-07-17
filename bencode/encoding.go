package bencode

import (
	"bytes"
	"fmt"
	"github.com/EdmundMartin/discrete/protocol"
	"log"
	"net"
	"unicode/utf8"
)

func writeStringData(val1 string, val2 string) string {
	var buffer bytes.Buffer

	buffer.WriteString(val1)
	buffer.WriteString(val2)

	return buffer.String()
}

func ByteString(key string) string {
	return fmt.Sprintf("%d:%s", utf8.RuneCountInString(key), key)
}

func KeyValue(key string, value string) string {
	key = ByteString(key)
	if value[0] == 'i' || value[0] == 'l' || value[0] == 'd' {
		value = value
	} else {
		value = ByteString(value)
	}
	return writeStringData(key, value)
}

func Int(x int) string {
	return fmt.Sprintf("i%de", x)
}

func CompactIP(buf *bytes.Buffer, ip string, port int) error {
	cmpIP := net.ParseIP(ip).To4()
	if cmpIP == nil {
		fmt.Errorf("unable to coerce to IP, %s", ip)
	}
	buf.Write(cmpIP)
	portCmp := []byte{byte(port >> 8), byte(port)}
	buf.Write(portCmp)
	return nil
}

func CompactPeers(peers []*protocol.Peer) []byte {
	var buf bytes.Buffer
	for _, p := range peers {
		err := CompactIP(&buf, p.IP, p.Port)
		if err != nil {
			log.Printf("failed to compact IP/Port %s:%d", p.IP, p.Port)
		}
	}
	return buf.Bytes()
}

func countStatus(peers []*protocol.Peer, completeStatus bool) int {
	count := 0
	for _, p := range peers {
		if p.Seed == completeStatus {
			count++
		}
	}
	return count
}

func EncodePeerResponse(peers []*protocol.Peer) string {
	result := ""

	complete := countStatus(peers, true)
	result += KeyValue("complete", Int(complete))

	incomplete := countStatus(peers, false)
	result += KeyValue("incomplete", Int(incomplete))

	if len(peers) > 0 {
		ipPorts := string(CompactPeers(peers))
		result += KeyValue("peers", ipPorts)
	}

	resp := fmt.Sprintf("d%se", result)
	return resp
}
