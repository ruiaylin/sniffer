package bootstrap
import (
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket"
	"github.com/binlaniua/sniffer/logger"
	"github.com/binlaniua/sniffer/datastream"
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket/layers"
)

var _MAX_PACKAGE_SIZE int32 = 1024 * 1024
var log = logger.Logger{"bootstrap"}
var event chan interface{}

func Start(device string, bpfExp string) {
	//
	event = make(chan interface{}, 1024)

	//
	handle, err := pcap.OpenLive(device, _MAX_PACKAGE_SIZE, true, pcap.BlockForever)
	if err != nil {
		log.Error("开启监听失败[%v]", err)
	}

	//
	err = handle.SetBPFFilter(bpfExp)
	if err != nil {
		log.Error("bpf表达式错误")
	}

	//
	packageSource := gopacket.NewPacketSource(handle, handle.LinkType())

	//
	go doCaptureLoop(packageSource)

	//
	doEventLoop()
}


func doCaptureLoop(packageSource *gopacket.PacketSource) {
	//
	dataStream := datastream.TcpDataStreamFactory{event}
	streamPool := tcpassembly.NewStreamPool(dataStream)
	assembler := tcpassembly.NewAssembler(streamPool)
	log.Debug("构建tcp分析者完成")

	//
	for packet := range packageSource.Packets() {
		log.Debug("数据包街获取")
		netLayer := packet.NetworkLayer()
		if netLayer == nil {
			continue
		}
		transportLayer := packet.TransportLayer()
		if transportLayer == nil {
			continue
		}
		tcp, _ := transportLayer.(*layers.TCP)
		udp, _ := transportLayer.(*layers.UDP)
		if tcp == nil || udp == nil {
			continue

			//udp packet
		} else if tcp == nil {
			log.Debug("udp coming, continue")

			//tcp packet
		} else if udp == nil {
			log.Debug("tcp coming")
			event <- packet.Data()
			assembler.AssembleWithTimestamp(
				netLayer.NetworkFlow(),
				tcp,
				packet.Metadata().CaptureInfo.Timestamp)
		}
	}

	//
	assembler.FlushAll()
}


func doEventLoop() {
	for evt := range event {
		log.Debug("data: %v", evt)
	}
}