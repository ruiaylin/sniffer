package http
import (
	"sync"
	"github.com/binlaniua/sniffer/logger"
)


var logHD = logger.Logger{"http.data"}

type HttpData struct {
	wg *sync.WaitGroup
	requestStream  *HttpRequestStream
	responseStream *HttpResponseStream
}

func NewHttpData() *HttpData {
	data := HttpData{}
	data.wg = new(sync.WaitGroup)
	return &data
}


func (data HttpData) StartResponse(s *HttpResponseStream) {
	data.responseStream = s
}

func (data HttpData) Wait() {
//	logHD.Debug("等待數據傳輸完成")
//	data.wg.Wait()
//	logHD.Debug("數據傳輸完成")
}