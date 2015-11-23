package http
import (
	"github.com/google/gopacket/tcpassembly"
	"github.com/google/gopacket"
	"fmt"
	"github.com/binlaniua/sniffer/logger"
	"strconv"
)

var log = logger.Logger{"http.Http"}
var dataMap map[string]HttpData

type Http struct {
}

func NewHttp() Http {
	dataMap = make(map[string]HttpData)
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
		isUp = true
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Src(), tcpFlow.Src(), netFlow.Dst(), tcpFlow.Dst())
	} else if isSupport(srcPort) {
		isUp = false
		key = fmt.Sprintf("%v:%v-%v:%v", netFlow.Dst(), tcpFlow.Dst(), netFlow.Src(), tcpFlow.Src())
	}

	//
	if key != "" {
		log.Debug("key: %v", key)
		httpData, ok := dataMap[key]
		if !ok {
//			log.Debug("new http data ")
			httpData = NewHttpData()
			dataMap[key] = httpData
			go httpData.Wait()
		}

		//
		if isUp {
			httpData.StartRequest()
			httpData.upStream = NewHttpStream(isUp, httpData.wg)
			return httpData.upStream;
		} else {
			httpData.StartResponse()
			httpData.downStream = NewHttpStream(isUp, httpData.wg)
			delete(dataMap, key)
			return httpData.downStream;
		}
	}

	//
	return nil
}

//
func isSupport(port int) bool {
	return port == 8099
}