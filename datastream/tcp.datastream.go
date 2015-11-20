package datastream
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/binlaniua/sniffer/protocol"
)

type TcpDataStream struct {
	protocol protocol.Protocol
}

func (tds TcpDataStream) Reassembled(rs []tcpassembly.Reassembly) {
	if tds.protocol != nil {
		tds.protocol.Handler(rs)
	}
}

func (tds TcpDataStream) ReassemblyComplete() {
	if tds.protocol != nil {
		tds.protocol.Close()
	}
}