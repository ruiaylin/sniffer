package datastream
import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
	"github.com/binlaniua/sniffer/protocol"
	"fmt"
	"strconv"
)


type TcpDataStreamFactory struct {
	Event chan interface{}
	Protocol []protocol.Protocol
}

func (dataStreamFactory TcpDataStreamFactory) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	//
	destPort, err := strconv.Atoi(fmt.Sprintf("%v", tcpFlow.Dst()))
	if err != nil {
		panic("非法数据")
	}

	//
	stream := TcpDataStream{}

	//
	for _, pcol := range dataStreamFactory.Protocol {
		if pcol.IsSupportPort(destPort) {
			pcol.SetFlow(netFlow, tcpFlow)
			stream.protocol = pcol
		}
	}
	return stream
}