package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mackerelio/mackerel-client-go"
	oremkrcli "github.com/yhidetoshi/ore-mkr-cli/lib"
)

const version = "0.0.1"

var (
	// flag.String
	argOrg    = flag.String("org", "", "set org")
	argType   = flag.String("type", "", "set type")
	argTarget = flag.String("target", "", "input target")

	// flag.Bool
	argWorking     = flag.Bool("working", false, "working")
	argStandby     = flag.Bool("standby", false, "standby")
	argRetire      = flag.Bool("retire", false, "retire host")
	argMaintenance = flag.Bool("maintenance", false, "maintenance host")
	argPoweroff    = flag.Bool("poweroff", false, "poweroff host")
	argDelete      = flag.Bool("delete", false, "delete user")

	// set mkr key each org
	mkrKeyOrgA = os.Getenv("MKRKEY_OrgA")
	mkrKeyOrgB = os.Getenv("MKRKEY_OrgB")

	client = mackerel.NewClient("")

	// OrgA first org
	OrgA = "orgA"

	// OrgB second org
	OrgB = "orgB"

	// WORKING status
	WORKING = "working"

	// STANDBY status
	STANDBY = "standby"

	// MAINTENANCE status
	MAINTENANCE = "maintenance"

	// POWEROFF status
	POWEROFF = "poweroff"
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
			err := oremkrcli.MakeHostStatus(client, *argTarget, status)
			if err != nil {
				fmt.Println(err)
			}
		} else if *argStandby {
			status := STANDBY
			err := oremkrcli.MakeHostStatus(client, *argTarget, status)
			if err != nil {
				fmt.Println(err)
			}
		} else if *argMaintenance {
			status := MAINTENANCE
			err := oremkrcli.MakeHostStatus(client, *argTarget, status)
			if err != nil {
				fmt.Println(err)
			}
		} else if *argPoweroff {
			status := POWEROFF
			err := oremkrcli.MakeHostStatus(client, *argTarget, status)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := oremkrcli.FetchHost(client)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// Monitor Commands
	if *argType == "monitor" {
		err := oremkrcli.FetchMonitorIDs(client)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Alert Commands
	if *argType == "alert" {
		err := oremkrcli.FetchOpenAlertIDs(client)
		if err != nil {
			fmt.Println(err)
		}
	}

	// User Commands
	if *argType == "user" {
		if *argDelete {
			err := oremkrcli.DeleteUser(client, *argTarget)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err := oremkrcli.FetchUsers(client)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
