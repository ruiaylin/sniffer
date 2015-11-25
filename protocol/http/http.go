package http
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"github.com/binlaniua/sniffer/logger"
	"strconv"
)

var log = logger.Logger{"http.Http"}
var dataMap map[string]*HttpData

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
	var key string = ""
	var isUp bool
	if isSupport(destPort) {
		log.Debug("分析请求数据")
		isUp = true
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Src(), tcpFlow.Src(), netFlow.Dst(), tcpFlow.Dst())
	} else if isSupport(srcPort) {
		log.Debug("分析响应数据")
		isUp = false
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Dst(), tcpFlow.Dst(), netFlow.Src(), tcpFlow.Src())
	}

	//
	if key != "" {
		httpData, ok := dataMap[key]
		if !ok {
			log.Debug("key: %v", key)
			httpData = NewHttpData()
			dataMap[key] = httpData
		}

		//
		if isUp {
			httpData.wg.Add(1)
			stream := NewHttpRequestStream(httpData.wg);
			return &stream.reader
		} else if httpData.requestStream != nil {
			httpData.wg.Add(1)
			stream := NewHttpResponseStream(httpData.requestStream, httpData.wg)
			delete(dataMap, key)
			go httpData.Wait()
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