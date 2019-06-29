package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
)


const (
	HOST = "host"
)


func FetchHost(client *mackerel.Client) {
	//var roleInfo string

	hostId, _ := client.FindHosts(
		&mackerel.FindHostsParam{Statuses: []string{"working", "standby", "maintenance", "poweroff"}})

	hostLists := [][]string{}

	for _, v := range hostId {
		//fmt.Println(v.Name)

		// convert int32 to string
		createdAtString := fmt.Sprint(v.DateFromCreatedAt())

		roleName := strings.Join(v.GetRoleFullnames()," ")

		hostList := []string{
			v.Status,
			v.Name,
			v.Type,
			roleName,
			createdAtString,
		}
		hostLists = append(hostLists, hostList)
	}
	OutputFormat(hostLists, HOST)
}

