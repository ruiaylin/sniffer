package http
import (
	"sync"
	"github.com/binlaniua/sniffer/logger"
)


var logHD = logger.Logger{"http.data"}

type HttpData struct {
	wg *sync.WaitGroup
	upStream   HttpStream
	downStream HttpStream
}

func NewHttpData() HttpData {
	data := HttpData{}
	data.wg = new(sync.WaitGroup)
	return data
}

func (data HttpData) StartRequest() {
	data.wg.Add(1)
}

func (data HttpData) StartResponse() {
	data.wg.Add(1)
}

func (data HttpData) Wait() {
	logHD.Debug("等待數據傳輸完成")
	data.wg.Wait()
	logHD.Debug("數據傳輸完成")
}