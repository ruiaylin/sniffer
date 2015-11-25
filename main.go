package main
import (
	"flag"
	"github.com/binlaniua/sniffer/bootstrap"
)



var (
	device = flag.String("-e", "en0", "")
	bpfExp = flag.String("-b", "tcp and host 172.19.3.77", "")
)

func main() {
	//
	flag.Parse();

	//
	bootstrap.Start(*device, *bpfExp);
}
