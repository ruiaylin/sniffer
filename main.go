package main
import (
	"flag"
	"github.com/binlaniua/sniffer/bootstrap"
)



var (
	device = flag.String("-e", "en0", "")
	bpfExp = flag.String("-b", "tcp port 1521", "")
)

func main() {
	//
	flag.Parse();

	//
	bootstrap.Start(device, bpfExp);
}
