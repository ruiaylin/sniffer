package http
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"github.com/binlaniua/sniffer/logger"
)

var log = logger.Logger{"http.Http"}

type Http struct {
}

func (http Http) Init() {
}

func (http Http) IsSupportPort(port int) bool {
	return true
}

func (htp Http) SetFlow(netFlow, tcpFlow gopacket.Flow) {
	flow := fmt.Sprintf("%v:%v -> %v:%v",
		netFlow.Dst(),
		tcpFlow.Dst(),
		netFlow.Src(),
		tcpFlow.Src())
	log.Debug(flow)
}

func (http Http) Handler(rs []tcpassembly.Reassembly) {
}

func (http Http) Close() {
}