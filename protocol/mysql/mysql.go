package mysql
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"strconv"
	"github.com/binlaniua/sniffer/logger"
)

var logO = logger.Logger{"mysql"}

type MySql struct {
}

func NewMySql() MySql {
	return MySql{}
}

//
func (http MySql) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	destPort, _ := strconv.Atoi(fmt.Sprintf("%v", tcpFlow.Dst()))

	//
	if isSupport(destPort) {
		return NewMySqlStream()
	}

	//
	return nil
}

//
func isSupport(port int) bool {
	return port == 3306
}