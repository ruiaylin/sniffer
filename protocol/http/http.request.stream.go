package http
import (
	"github.com/google/gopacket/tcpassembly"
	"bytes"
	"sync"
	"strings"
	"github.com/binlaniua/sniffer/logger"
)

var logHRS = logger.Logger{"http.request.stream"}

type HttpRequestStream struct {
	content *bytes.Buffer
	header  map[string]string
	body    string
	wg      *sync.WaitGroup
}

func NewHttpRequestStream(wg *sync.WaitGroup) HttpRequestStream {
	stream := HttpRequestStream{}
	stream.content = bytes.NewBuffer([]byte(""))
	stream.wg = wg;
	return stream
}

func (stream HttpRequestStream) Reassembled(rc []tcpassembly.Reassembly) {
	//
	for _, packet := range rc {
		stream.content.Write(packet.Bytes)
	}

	//
	//	var index int = -1
	//	var flag = []byte("\r\n\r\n")  //请求头和请求体的分隔符
	//	index = bytes.Index(stream.content.Bytes(), flag)
	//	if index != -1 {
	//		content := string(stream.content.Next(index + len(flag)))
	//		logHS.Debug(stream.typeStr + " content is " + content)
	//		stream.wg.Done()
	//	}

	//
	if stream.header == nil {
		stream.getHeaders()
	}
}

func (stream HttpRequestStream) ReassemblyComplete() {
}


func (stream HttpRequestStream) getHeaders() {
	defer errorHandler()

	//请求头和请求体的分隔符
	var flag = []byte("\r\n\r\n")

	//
	index := bytes.Index(stream.content.Bytes(), flag)
	var content string
	if index != -1 {
		content = string(stream.content.Next(index - len(flag)))
	}

	//
	if content == "" {
		return
	}

	//按\r\n分割
	stream.header = make(map[string]string)
	for line, headLine := range strings.Split(content, "\r\n") {
		if line == 0 { //第一行是method, path, protocol
			stream.getRequestInfo(headLine)
		} else {  //这些都是请求头了
			splitIndex := strings.Index(headLine, ":")
			if splitIndex == -1 {
				logHRS.Debug("错误的请求头: %v", headLine)
			} else {
				key := headLine[:splitIndex]
				val := headLine[splitIndex:]
				stream.header[key] = val
			}
		}
	}
}

func (stream HttpRequestStream) getRequestInfo(head string) {
	hs := strings.Split(head, " ")
	stream.header["requestMethod"] = hs[0]
	stream.header["requestURI"] = hs[1]
	stream.header["requestProtocol"] = hs[2]
}

func errorHandler() {
	if err := recover(); err != nil {
		logHRS.Debug("%v", err)
	}
}
