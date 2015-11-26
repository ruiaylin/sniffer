package http
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"strconv"
	"github.com/binlaniua/sniffer/logger"
)

var dataMap map[string]*HttpData
var logH = logger.Logger{"http"}

type Http struct {
}

func NewHttp() Http {
	dataMap = make(map[string]*HttpData)
	return Http{}
}

//
func (http Http) New(netFlow, tcpFlow gopacket.Flow) (ret tcpassembly.Stream) {
	destPort, _ := strconv.Atoi(fmt.Sprintf("%v", tcpFlow.Dst()))
	srcPort, _ := strconv.Atoi(fmt.Sprintf("%v", tcpFlow.Src()))

	//

	//
	var key string = ""
	var isUp bool
	if isSupport(destPort) {
		isUp = true
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Src(), tcpFlow.Src(), netFlow.Dst(), tcpFlow.Dst())
	} else if isSupport(srcPort) {
		isUp = false
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Dst(), tcpFlow.Dst(), netFlow.Src(), tcpFlow.Src())
	}

	//
	if key != "" {
		logH.Debug("%v", netFlow)
		httpData, ok := dataMap[key]
		if !ok {
			httpData = NewHttpData()
			dataMap[key] = httpData
		}

		//請求
		if isUp {
			stream := NewHttpRequestStream(httpData.wg);
			httpData.requestStream = stream
			return &stream.reader

		//响应
		} else if httpData.requestStream != nil {
			stream := NewHttpResponseStream(httpData.requestStream, httpData.wg)
			httpData.responseStream = stream

			//
			delete(dataMap, key)
			return &stream.reader
		}
	}

	//
	return nil
}

//
func isSupport(port int) bool {
	return port == 8099
}