package oracle
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"strconv"
	"github.com/binlaniua/sniffer/logger"
)

var logO = logger.Logger{"oracle"}

type Oracle struct {
}

func NewOracle() Oracle {
	return Oracle{}
}

//
func (http Oracle) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	destPort, _ := strconv.Atoi(fmt.Sprintf("%v", tcpFlow.Dst()))

	//
	if isSupport(destPort) {
		return NewOracleStream()
	}

	//
	return nil
}

//
func isSupport(port int) bool {
	return port == 1521
}