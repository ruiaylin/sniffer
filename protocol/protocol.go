package protocol
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
)


type Protocol interface {
	New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream)
}