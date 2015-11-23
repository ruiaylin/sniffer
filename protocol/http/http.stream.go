package http
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/binlaniua/sniffer/logger"
	"bytes"
	"sync"
)

var logHS = logger.Logger{"http.stream"}

type HttpStream struct {
	isUp    bool
	typeStr string
	content *bytes.Buffer
	header  string
	body    string
	wg *sync.WaitGroup
}

func NewHttpStream(isUp bool, wg *sync.WaitGroup) HttpStream {
	var typeStr string
	if isUp {
		typeStr = "请求数据 "
	} else {
		typeStr = "响应数据"
	}
	stream := HttpStream{isUp: isUp, typeStr: typeStr}
	stream.content = bytes.NewBuffer([]byte(""))
	stream.wg = wg;
	return stream
}

func (stream HttpStream) Reassembled(rc []tcpassembly.Reassembly) {
	var index int = -1
	var flag = []byte("\r\n\r\n")
	for _, packet := range rc {
		stream.content.Write(packet.Bytes)
		index = bytes.Index(stream.content.Bytes(), flag)
		if index != -1 {
			break
		}
	}
	if index != -1 {
		content := string(stream.content.Next(index + len(flag)))
		logHS.Debug(stream.typeStr + " content is " + content)
		stream.wg.Done()
	}
}

func (stream HttpStream) ReassemblyComplete() {
	logHS.Debug(stream.typeStr + " content is %v", string(stream.content.Bytes()))
}