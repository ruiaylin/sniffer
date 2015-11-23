package datastream
import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/binlaniua/sniffer/protocol"
	"github.com/binlaniua/sniffer/logger"
)

var log = logger.Logger{"tcp.datastream.factory"}

type TcpDataStreamFactory struct {
	Event    chan interface{}
	Protocol []protocol.Protocol
}

type NopStream struct {
}

func (nop NopStream) Reassembled([] tcpassembly.Reassembly) {
}

func (nop NopStream) ReassemblyComplete() {
}

func (dataStreamFactory TcpDataStreamFactory) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	for _, pcol := range dataStreamFactory.Protocol {
		stream := pcol.New(netFlow, tcpFlow)
		if stream != nil {
			return stream
		}
	}
	return NopStream{}
}