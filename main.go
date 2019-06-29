package main

import (
	"flag"
	"github.com/mackerelio/mackerel-client-go"
	oremkrcli "github.com/yhidetoshi/ore-mkr-cli/lib"
	"os"
)

var (
	argOrg = flag.String("org", "", "set org")
	argType = flag.String("type", "", "set type")
	argHostname = flag.String("", "", "input image name")

	argRetire = flag.Bool("retire", false, "retire host")

	// set mkr key each org
	mkrKeyOrgA   = os.Getenv("MKRKEY_OrgA")
	mkrKeyOrgB  = os.Getenv("MKRKEY_OrgB")

	client = mackerel.NewClient("")


	OrgA = "orgA"
	OrgB = "orgB"
)

func main() {
	flag.Parse()

	// switch mkr apikey
	switch *argOrg{
	case OrgA:
		client = mackerel.NewClient(mkrKeyOrgA)

	case OrgB:
		client = mackerel.NewClient(mkrKeyOrgB)
	}

	// Host Commands
	if *argType == "host"{
		if *argRetire {

		}else{
			oremkrcli.FetchHost(client)
		}
	}

}
