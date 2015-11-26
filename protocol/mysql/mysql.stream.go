package mysql
import (
	"github.com/binlaniua/sniffer/logger"
	"github.com/google/gopacket/tcpassembly"
	"bytes"
)

var logOS = logger.Logger{"mysql.stream"}

type MySqlStream struct {
	content *bytes.Buffer
}

var logMSS = logger.Logger{"mysql.stream"}

func NewMySqlStream() *MySqlStream {
	//
	stream := MySqlStream{}
	stream.content = bytes.NewBuffer([]byte(""))

	//
	return &stream
}

func (stream *MySqlStream) Reassembled(reassembly []tcpassembly.Reassembly) {
	//
	logMSS.Debug("开始获取MYSQL数据， 等待数据来临")
	for _, r := range reassembly {

		//
		logMSS.Debug("数据来临:" + string(r.Bytes))
//		stream.content.Write(r.Bytes)
//		logOS.Debug(stream.content.String())
	}
}

// ReassemblyComplete implements tcpassembly.Stream's ReassemblyComplete function.
func (r *MySqlStream) ReassemblyComplete() {
	r.content = nil
}
