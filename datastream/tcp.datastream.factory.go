package datastream
import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/tcpassembly"
)


type TcpDataStreamFactory struct {
	Event chan interface{}
}


func (dataStreamFactory TcpDataStreamFactory) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	return
}