package bootstrap
import (
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"github.com/binlaniua/sniffer/logger"
	"github.com/binlaniua/sniffer/datastream"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/layers"
	"github.com/binlaniua/sniffer/protocol"
	"github.com/binlaniua/sniffer/protocol/http"
	"time"
)

var eachPacketSize int32 = 1024 * 1024
var log = logger.Logger{"bootstrap"}
var ps = [...]protocol.Protocol{http.NewHttp()}

func Start(device *string, bpfExp *string) {
	//
	handle, err := pcap.OpenLive(*device, eachPacketSize, true, pcap.BlockForever)
	if err != nil {
		log.Error("开启监听失败[%v]", err)
	}

	//
	err = handle.SetBPFFilter(*bpfExp)
	if err != nil {
		log.Error("bpf表达式错误")
	}

	//
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())

	//
	doCaptureLoop(packageSource)
}


func doCaptureLoop(packageSource *gopacket.PacketSource) {
	//
	dataStream := datastream.TcpDataStreamFactory{ps[:len(ps)]}
	streamPool := tcpassembly.NewStreamPool(dataStream)
	assembler := tcpassembly.NewAssembler(streamPool)
	log.Debug("构建tcp分析者完成")

	//
	packets := packageSource.Packets()
	ticker := time.Tick(time.Minute)

	//
	for {
		select {
		case packet := <-packets:
			{
				if packet == nil {
					return
				}
				if packet.NetworkLayer() == nil || packet.TransportLayer() == nil {
					continue
				}

				//tcp
				if (packet.TransportLayer().LayerType() == layers.LayerTypeTCP) {
					tcp, _ := packet.TransportLayer().(*layers.TCP)
					assembler.AssembleWithTimestamp(
						packet.NetworkLayer().NetworkFlow(),
						tcp,
						packet.Metadata().Timestamp)
				}
			}
		case <-ticker:
			assembler.FlushOlderThan(time.Now().Add(time.Minute * -2))
		}
	}
}
