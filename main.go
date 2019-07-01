package main

import (
	"flag"
	"github.com/mackerelio/mackerel-client-go"
	oremkrcli "github.com/yhidetoshi/ore-mkr-cli/lib"
	"os"
)

var (
	// flag.String
	argOrg    = flag.String("org", "", "set org")
	argType   = flag.String("type", "", "set type")
	argHostID = flag.String("target", "", "input target hostID")

	// flag.Bool
	argWorking     = flag.Bool("working", false, "working")
	argStandby     = flag.Bool("standby", false, "standby")
	argRetire      = flag.Bool("retire", false, "retire host")
	argMaintenance = flag.Bool("maintenance", false, "maintenance host")
	argPoweroff    = flag.Bool("poweroff", false, "poweroff host")

	// set mkr key each org
	mkrKeyOrgA = os.Getenv("MKRKEY_OrgA")
	mkrKeyOrgB = os.Getenv("MKRKEY_OrgB")

	client = mackerel.NewClient("")

	OrgA        = "orgA"
	OrgB        = "orgB"
	WORKING     = "working"
	STANDBY     = "standby"
	MAINTENANCE = "maintenance"
	POWEROFF    = "poweroff"
)

func main() {
	flag.Parse()

	// switch mkr apikey
	switch *argOrg {
	case OrgA:
		client = mackerel.NewClient(mkrKeyOrgA)

	case OrgB:
		client = mackerel.NewClient(mkrKeyOrgB)
	}

	// Host Commands
	if *argType == "host" {
		if *argWorking {
			status := WORKING
			oremkrcli.MakeHostStatus(client, *argHostID, status)
		} else if *argStandby {
			status := STANDBY
			oremkrcli.MakeHostStatus(client, *argHostID, status)
		} else if *argMaintenance {
			status := MAINTENANCE
			oremkrcli.MakeHostStatus(client, *argHostID, status)
		} else if *argPoweroff {
			status := POWEROFF
			oremkrcli.MakeHostStatus(client, *argHostID, status)
		} else {
			oremkrcli.FetchHost(client)
		}
	}

}
