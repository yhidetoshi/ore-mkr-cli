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
	hostId, _ := client.FindHosts(
		&mackerel.FindHostsParam{Statuses: []string{"working", "standby", "maintenance", "poweroff"}})

	hostLists := [][]string{}

	for _, v := range hostId {
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
	OutputFormat(hostLists, HOST)
}

func MakeHostStatus(client *mackerel.Client, hostIDs string, status string) {
	// カンマ区切りを配列に変換
	targetHostIDs := strings.Split(hostIDs, ",")

	for i, _ := range targetHostIDs {
		err := client.UpdateHostStatus(targetHostIDs[i], status)
		if err != nil {
			fmt.Println("Failed status change: %v\n", targetHostIDs[i])
		} else {
			fmt.Printf("Sucessed change status: %s\n", targetHostIDs[i])
		}
	}

}
