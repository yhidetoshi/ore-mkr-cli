package oremkrcli

import (
	"encoding/json"
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
)

const (
	MONITOR = "monitor"
	BLANK   = ""
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

type MonitorConnectivityValues struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Memo                 string `json:"memo,omitempty"`
	Type                 string `json:"type,omitempty"`
	IsMute               bool   `json:"isMute,omitempty"`
	NotificationInterval uint64 `json:"notificationInterval,omitempty"`

	Scopes        []string `json:"scopes,omitempty"`
	ExcludeScopes []string `json:"excludeScopes,omitempty"`
}

type MonitorExternalHTTPValues struct {
	ID                   string `json:"id,omitempty"`
	Name                 string `json:"name,omitempty"`
	Memo                 string `json:"memo,omitempty"`
	Type                 string `json:"type,omitempty"`
	IsMute               bool   `json:"isMute,omitempty"`
	NotificationInterval uint64 `json:"notificationInterval,omitempty"`

	Method                          string   `json:"method,omitempty"`
	URL                             string   `json:"url,omitempty"`
	MaxCheckAttempts                uint64   `json:"maxCheckAttempts,omitempty"`
	Service                         string   `json:"service,omitempty"`
	ResponseTimeCritical            *float64 `json:"responseTimeCritical,omitempty"`
	ResponseTimeWarning             *float64 `json:"responseTimeWarning,omitempty"`
	ResponseTimeDuration            *uint64  `json:"responseTimeDuration,omitempty"`
	RequestBody                     string   `json:"requestBody,omitempty"`
	ContainsString                  string   `json:"containsString,omitempty"`
	CertificationExpirationCritical *uint64  `json:"certificationExpirationCritical,omitempty"`
	CertificationExpirationWarning  *uint64  `json:"certificationExpirationWarning,omitempty"`
	SkipCertificateVerification     bool     `json:"skipCertificateVerification,omitempty"`
}

func FetchMonitorIDs(client *mackerel.Client) {
	var listMonitorsHostIDs []string
	var listMonitorConnectivityIDs []string
	var listMonitorsExternalIDs []string

	monitors, err := client.FindMonitors()
	if err != nil {
		fmt.Println("fail get monitors")
	}

	for _, v := range monitors {
		switch v.MonitorType() {
		case "host":
			listMonitorsHostIDs = append(listMonitorsHostIDs, v.MonitorID())
		case "connectivity":
			listMonitorConnectivityIDs = append(listMonitorConnectivityIDs, v.MonitorID())
		case "external":
			listMonitorsExternalIDs = append(listMonitorsExternalIDs, v.MonitorID())
		}
	}
	mhm := &MonitorHostMetric{}
	meh := &MonitorExternalHTTPValues{}
	mc := &MonitorConnectivityValues{}

	MergeMonitorResult(
		mhm.DescribeMonitorHostByID(client, listMonitorsHostIDs),
		meh.DescribeMonitorExternalByID(client, listMonitorsExternalIDs),
		mc.MonitorConnectivityByID(client, listMonitorConnectivityIDs),
	)
}

func MergeMonitorResult(hostResult [][]string, externalResult [][]string, connectivityResult [][]string) {
	merged := append(hostResult, externalResult...)
	OutputFormat(append(merged, connectivityResult...), MONITOR)
}

func (mc *MonitorConnectivityValues) MonitorConnectivityByID(client *mackerel.Client, list []string) [][]string {
	var monitorConnectivityValues MonitorConnectivityValues
	monitorLists := [][]string{}

	for i := range list {
		res, _ := client.GetMonitor(list[i])
		valueBytesJSON, _ := json.Marshal(res)
		bytesValue := []byte(valueBytesJSON)

		if err := json.Unmarshal(bytesValue, &monitorConnectivityValues); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}
		scope := strings.Join(monitorConnectivityValues.Scopes, ":")
		monitorList := []string{
			monitorConnectivityValues.ID,
			monitorConnectivityValues.Name,
			scope,
			BLANK,
			BLANK,
			BLANK,
			BLANK,
			monitorConnectivityValues.Memo,
		}
		monitorLists = append(monitorLists, monitorList)
	}
	return monitorLists
}

func (mhm *MonitorHostMetric) DescribeMonitorHostByID(client *mackerel.Client, list []string) [][]string {
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
			//monitorHostValues.Type,
			scope,
			stringWarninng,
			stringCritical,
			fmt.Sprint(monitorHostValues.Duration),
			fmt.Sprint(monitorHostValues.MaxCheckAttempts),
			monitorHostValues.Memo,
		}

		monitorLists = append(monitorLists, monitorList)
	}
	//OutputFormat(monitorLists, MONITOR)
	return monitorLists
}

func (meh MonitorExternalHTTPValues) DescribeMonitorExternalByID(client *mackerel.Client, list []string) [][]string {
	var monitorExternalHTTPValues MonitorExternalHTTPValues

	monitorLists := [][]string{}

	for i := range list {
		res, _ := client.GetMonitor(list[i])
		valueBytesJSON, _ := json.Marshal(res)
		bytesValue := []byte(valueBytesJSON)

		if err := json.Unmarshal(bytesValue, &monitorExternalHTTPValues); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}
		monitorList := []string{
			monitorExternalHTTPValues.ID,
			monitorExternalHTTPValues.Name,
			monitorExternalHTTPValues.Service,
			fmt.Sprint(*monitorExternalHTTPValues.ResponseTimeWarning),
			fmt.Sprint(*monitorExternalHTTPValues.ResponseTimeCritical),
			fmt.Sprint(*monitorExternalHTTPValues.ResponseTimeDuration),
			fmt.Sprint(monitorExternalHTTPValues.MaxCheckAttempts),
			monitorExternalHTTPValues.Memo,
		}
		monitorLists = append(monitorLists, monitorList)
	}
	//OutputFormat(monitorLists, MONITOR)
	return monitorLists
}
