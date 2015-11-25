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

func NewHttpRequestStream(wg *sync.WaitGroup) HttpRequestStream {
	//
	stream := HttpRequestStream{}
	stream.reader = tcpreader.NewReaderStream()
	stream.wg = wg;

	//
	logHRSCount++

	//
	go stream.start()
	return stream
}

func (stream HttpRequestStream) start() {
	buf := bufio.NewReader(&stream.reader)
	for {
		req, err := http.ReadRequest(buf)
		if err == io.EOF {
			logHRS.Debug("end request eof, %v", logHRSCount)
			stream.end()
			break
		} else if err != nil {
			logHRS.Debug("something error %v", err)
		} else {
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
