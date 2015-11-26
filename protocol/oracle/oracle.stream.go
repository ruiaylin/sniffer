package oracle
import (
	"github.com/binlaniua/sniffer/logger"
	"github.com/google/gopacket/tcpassembly"
	"bytes"
)

var logOS = logger.Logger{"oracle.stream"}

type OracleStream struct {
	content *bytes.Buffer
}

func NewOracleStream() *OracleStream {
	//
	stream := OracleStream{}
	stream.content = bytes.NewBuffer([]byte(""))

	//
	return &stream
}

func (stream *OracleStream) Reassembled(reassembly []tcpassembly.Reassembly) {
	for _, r := range reassembly {
		stream.content.Write(r.Bytes)
		logOS.Debug(stream.content.String())
	}
}

// ReassemblyComplete implements tcpassembly.Stream's ReassemblyComplete function.
func (r *OracleStream) ReassemblyComplete() {
	r.content = nil
}
