package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
)
func FetchMonitorIDs(client *mackerel.Client){
	monitors, err  := client.FindMonitors()
	if err != nil {
		fmt.Println("fail get monitors")
	}
	//fmt.Println(&monitors)

	for _, v := range monitors {
		//fmt.Printf("%v\t%v\n"  ,v.MonitorName(), v.MonitorID())
		fmt.Println(v.MonitorID())
	}
	//DescribeMonitorByID(client)
}

/*
func DescribeMonitorByID(client *mackerel.Client){
	res, _ := client.GetMonitor("3DWLj5WU92E")

	fmt.Println(res)

}
*/
