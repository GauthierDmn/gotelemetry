package models

import (
	"time"
)

type PacketType string

const (
	SpacePacket PacketType = "spacepacket"
	CSP PacketType = "csp"
)

type Packet struct {
	ReceivedTimestamp *time.Time
	PacketType *PacketType
	Data []byte
}
