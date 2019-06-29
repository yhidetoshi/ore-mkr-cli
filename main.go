package main

import (
	"flag"
	"github.com/mackerelio/mackerel-client-go"
	"github.com/yhidetoshi/ore-mkr-cli/lib"
	"os"
)

var (
	argList = flag.Bool("list", false, "host list")
	mkrKey   = os.Getenv("MKRKEY")
	client = mackerel.NewClient(mkrKey)
)

func main() {
	flag.Parse()
	oremkrcli.FetchHost(client)

}
