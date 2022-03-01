package quic

import (
	"github.com/lucas-clemente/quic-go/internal/congestion"
	"github.com/lucas-clemente/quic-go/internal/protocol"
)

const (
	Max_packet_size_def_InitialPacketSizeIPv4 uint64 = 1252
	Max_packet_size_def_InitialPacketSizeIPv6 uint64 = 1232
	Max_packet_size_def_MaxPacketBufferSize   uint64 = 1452
	Max_packet_size_def_MinInitialPacketSize  uint64 = 1200
	Max_packet_size_def_max_packet_size       uint64 = 1200
)

func Set_nonstandard_max_packet_size(max_packet_size uint64) {
	n := max_packet_size
	//Cfg_Set(1252, 1232, 1452, 1200, 1200)
	Set_nonstandard_Max_packet_size_cfg(
		52+n-52,
		32+n-52,
		52+52+n-52,
		0+n-52,
		0+n+52,
	)
}

func Set_nonstandard_Max_packet_size_cfg(
	InitialPacketSizeIPv4,
	InitialPacketSizeIPv6,
	MaxPacketBufferSize,
	MinInitialPacketSize,
	max_packet_size uint64,
) {
	protocol.InitialPacketSizeIPv4 = InitialPacketSizeIPv4
	protocol.InitialPacketSizeIPv6 = InitialPacketSizeIPv6
	protocol.MaxPacketBufferSize = protocol.ByteCount(MaxPacketBufferSize)
	protocol.MinInitialPacketSize = MinInitialPacketSize
	protocol.Max_packet_size = max_packet_size
	congestion.CubeFactor = 1 << congestion.CubeScale / congestion.CubeCongestionWindowScale / congestion.MaxDatagramSize
	congestion.MaxDatagramSize = protocol.ByteCount(protocol.InitialPacketSizeIPv4)
}
