package main
import (
	"flag"
	"github.com/binlaniua/sniffer/bootstrap"
)



var (
	device = flag.String("-e", "en0", "")
	bpfExp = flag.String("-b", "tcp", "")
)

func main() {
	//
	flag.Parse();

	//
	bootstrap.Start(*device, *bpfExp);
}
