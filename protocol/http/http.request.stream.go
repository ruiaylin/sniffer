package http
import (
	"sync"
	"github.com/binlaniua/sniffer/logger"
	"bufio"
	"net/http"
	"io"
	"github.com/google/gopacket/tcpassembly/tcpreader"
	"fmt"
)

var logHRS = logger.Logger{"http.request.stream"}
var logHRSCount = 1

type HttpRequestStream struct {
	wg      *sync.WaitGroup
	reader  tcpreader.ReaderStream
	request *http.Request
}

func NewHttpRequestStream(wg *sync.WaitGroup) *HttpRequestStream {
	//
	stream := HttpRequestStream{}
	stream.reader = tcpreader.NewReaderStream()
	stream.wg = wg;

	//
	logHRSCount++
	stream.wg.Add(1)

	//
	go stream.start()
	return &stream
}

func (stream HttpRequestStream) start() {
	defer errorHandler()

	//
	logHRS.Debug("开始获取请求数据")

	//
	buf := bufio.NewReader(&stream.reader)
	for {
		logHRS.Debug("1")
		req, err := http.ReadRequest(buf)
		logHRS.Debug("2")
		if err == io.EOF {
			logHRS.Debug("请求数据全部获取完, %v", logHRSCount)
			stream.end()
			break
		} else if err != nil {
			logHRS.Debug("something error %v", err)
		} else {
			logHRS.Debug("请求数据获取完成")
			stream.request = req
			req.Body.Close()
		}
	}
}

func (stream HttpRequestStream) end() {
	stream.wg.Done()

	//
	req := stream.request

	//
	fmt.Println("========================")
	fmt.Printf("%20s host[%s] uri[%s] method[%s]\n", "", req.Host, req.RequestURI, req.Method)
	fmt.Println("========================")
}

func errorHandler()  {
	if err := recover(); err != nil {
		fmt.Printf("request error => %v", err)
	}
}
