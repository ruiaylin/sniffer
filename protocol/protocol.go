package protocol
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
)


type Protocol interface {

	//初始化
	Init()

	//是否支持对该端口的解析
	IsSupportPort(port int) bool

	//设置数据
	SetFlow(netFlow, tcpFlow gopacket.Flow)

	//处理数据
	Handler(rs []tcpassembly.Reassembly)

	//
	Close()
}