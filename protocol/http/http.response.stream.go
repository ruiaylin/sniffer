package http
import (
	"github.com/google/gopacket/tcpassembly"
	"bytes"
	"sync"
)

type HttpResponseStream struct {
	content *bytes.Buffer
	header  string
	body    string
	wg      *sync.WaitGroup
}

func NewHttpResponseStream(wg *sync.WaitGroup) HttpResponseStream {
	stream := HttpResponseStream{}
	stream.content = bytes.NewBuffer([]byte(""))
	stream.wg = wg;
	return stream
}

func (stream HttpResponseStream) Reassembled(rc []tcpassembly.Reassembly) {
	//
	for _, packet := range rc {
		stream.content.Write(packet.Bytes)
	}

	//	//
	//	var index int = -1
	//	var flag = []byte("\r\n\r\n")  //请求头和请求体的分隔符
	//	index = bytes.Index(stream.content.Bytes(), flag)
	//	if index != -1 {
	//		content := string(stream.content.Next(index + len(flag)))
	//		logHS.Debug(stream.typeStr + " content is " + content)
	//		stream.wg.Done()
	//	}
}

func (stream HttpResponseStream) ReassemblyComplete() {
}


