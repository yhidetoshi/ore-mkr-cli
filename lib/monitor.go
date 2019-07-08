package oremkrcli

import (
	"encoding/json"
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
)

const (
	MONITOR = "monitor"
)

type MonitorHostMetric struct {
	id          string
	name        string
	memo        string
	monitorType string
}

type MonitorHostValues struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Memo string `json:"memo,omitempty"`
	Type string `json:"type,omitempty"`

	IsMute               bool   `json:"isMute,omitempty"`
	NotificationInterval uint64 `json:"notificationInterval,omitempty"`

	Service          string   `json:"service,omitempty"`
	Metric           string   `json:"metric,omitempty"`
	Operator         string   `json:"operator,omitempty"`
	Warning          *float64 `json:"warning"`
	Critical         *float64 `json:"critical"`
	Duration         uint64   `json:"duration,omitempty"`
	Scopes           []string `json:"scopes,omitempty"`
	MaxCheckAttempts uint64   `json:"maxCheckAttempts,omitempty"`
}

func FetchMonitorIDs(client *mackerel.Client) {
	var listMonitorsIDs []string
	monitors, err := client.FindMonitors()
	if err != nil {
		fmt.Println("fail get monitors")
	}

	for _, v := range monitors {
		listMonitorsIDs = append(listMonitorsIDs, v.MonitorID())
	}
	mhm := &MonitorHostMetric{}
	mhm.DescribeMonitorByID(client, listMonitorsIDs)
}

func (m MonitorHostMetric) DescribeMonitorByID(client *mackerel.Client, list []string) {
	var monitorHostValues MonitorHostValues
	var stringCritical, stringWarninng string
	monitorLists := [][]string{}

	for i := range list {
		res, _ := client.GetMonitor(list[i])
		valueBytesJSON, _ := json.Marshal(res)
		bytesValue := []byte(valueBytesJSON)

		if err := json.Unmarshal(bytesValue, &monitorHostValues); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}
		scope := strings.Join(monitorHostValues.Scopes, ":")

		// warninngがセットされてない場合の処理
		if monitorHostValues.Warning == nil {
			stringWarninng = ""
		} else {
			stringWarninng = fmt.Sprint(*monitorHostValues.Warning)
		}

		// criticalがセットされてない場合の処理
		if monitorHostValues.Critical == nil {
			stringCritical = ""
		} else {
			stringCritical = fmt.Sprint(*monitorHostValues.Critical)
		}

		monitorList := []string{
			monitorHostValues.ID,
			monitorHostValues.Name,
			scope,
			stringWarninng,
			stringCritical,
			fmt.Sprint(monitorHostValues.Duration),
			fmt.Sprint(monitorHostValues.MaxCheckAttempts),
			monitorHostValues.Memo,
		}

		monitorLists = append(monitorLists, monitorList)
	}
	OutputFormat(monitorLists, MONITOR)
}
