package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
)

const (
	// HOST host
	HOST = "host"
)

// FetchHost fetch hosts
func FetchHost(client *mackerel.Client) error {
	hostID, err := client.FindHosts(
		&mackerel.FindHostsParam{Statuses: []string{"working", "standby", "maintenance", "poweroff"}})
	if err != nil {
		fmt.Println(err)
	}

	hostLists := [][]string{}

	for _, v := range hostID {
		//fmt.Println(v.Name)

		// convert int32 to string
		createdAtString := fmt.Sprint(v.DateFromCreatedAt())
		roleName := strings.Join(v.GetRoleFullnames(), " ")

		hostList := []string{
			v.Status,
			v.Name,
			v.ID,
			v.Type,
			roleName,
			createdAtString,
		}
		hostLists = append(hostLists, hostList)
	}
	err = OutputFormat(hostLists, HOST)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

// MakeHostStatus change host status
func MakeHostStatus(client *mackerel.Client, hostIDs string, status string) error {
	// カンマ区切りを配列に変換
	targetHostIDs := strings.Split(hostIDs, ",")

	for i := range targetHostIDs {
		err := client.UpdateHostStatus(targetHostIDs[i], status)
		if err != nil {
			fmt.Println("Failed status change: %v\n", targetHostIDs[i])
		} else {
			fmt.Printf("Sucessed change status: %s\n", targetHostIDs[i])
		}
	}
	return nil
}
