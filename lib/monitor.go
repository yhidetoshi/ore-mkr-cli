package oremkrcli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mackerelio/mackerel-client-go"
)

const (
	// MONITOR monitor
	MONITOR = "monitor"
	// BLANK null
	BLANK = ""
)

// MonitorHostMetric information
type MonitorHostMetric struct {
	id          string
	name        string
	memo        string
	monitorType string
}

// MonitorHostValues information
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

// MonitorConnectivityValues information
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

// MonitorExternalHTTPValues information
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

// FetchMonitorIDs fetch monitor ids.
func FetchMonitorIDs(client *mackerel.Client) error {
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
	mhv := &MonitorHostValues{}
	meh := &MonitorExternalHTTPValues{}
	mc := &MonitorConnectivityValues{}

	MergeMonitorResult(
		mhv.DescribeMonitorHostByID(client, listMonitorsHostIDs),
		meh.DescribeMonitorExternalByID(client, listMonitorsExternalIDs),
		mc.MonitorConnectivityByID(client, listMonitorConnectivityIDs),
	)
	return nil
}

// MergeMonitorResult merge monitor results.
func MergeMonitorResult(hostResult [][]string, externalResult [][]string, connectivityResult [][]string) {
	merged := append(hostResult, externalResult...)
	OutputFormat(append(merged, connectivityResult...), MONITOR)
}

//MonitorConnectivityByID find monitor connectivity by id.
func (mc *MonitorConnectivityValues) MonitorConnectivityByID(client *mackerel.Client, list []string) [][]string {
	monitorLists := [][]string{}

	for i := range list {
		res, err := client.GetMonitor(list[i])
		if err != nil {
			fmt.Println(err)
		}
		valueBytesJSON, _ := json.Marshal(res)

		if err := json.Unmarshal(valueBytesJSON, mc); err != nil {
			fmt.Println(err)
		}
		scope := strings.Join(mc.Scopes, ":")
		monitorList := []string{
			//monitorConnectivityValues.ID,
			mc.ID,
			mc.Name,
			scope,
			BLANK,
			BLANK,
			BLANK,
			BLANK,
			mc.Memo,
		}
		monitorLists = append(monitorLists, monitorList)
	}
	return monitorLists
}

// DescribeMonitorHostByID describe monitor hosts by id.
func (mhv *MonitorHostValues) DescribeMonitorHostByID(client *mackerel.Client, list []string) [][]string {
	var stringCritical, stringWarning string
	monitorLists := [][]string{}

	for i := range list {
		res, _ := client.GetMonitor(list[i])
		valueBytesJSON, _ := json.Marshal(res)

		if err := json.Unmarshal(valueBytesJSON, mhv); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}
		scope := strings.Join(mhv.Scopes, ":")
		// warninngがセットされてない場合の処理
		if mhv.Warning == nil {
			stringWarning = ""
		} else {
			stringWarning = fmt.Sprint(*mhv.Warning)
		}

		// criticalがセットされてない場合の処理
		if mhv.Critical == nil {
			stringCritical = ""
		} else {
			stringCritical = fmt.Sprint(*mhv.Critical)
		}

		monitorList := []string{
			mhv.ID,
			mhv.Name,
			//mhv.Type,
			scope,
			stringWarning,
			stringCritical,
			fmt.Sprint(mhv.Duration),
			fmt.Sprint(mhv.MaxCheckAttempts),
			mhv.Memo,
		}

		monitorLists = append(monitorLists, monitorList)
	}
	return monitorLists
}

// DescribeMonitorExternalByID describe monitor external by id.
func (meh *MonitorExternalHTTPValues) DescribeMonitorExternalByID(client *mackerel.Client, list []string) [][]string {
	monitorLists := [][]string{}

	for i := range list {
		res, _ := client.GetMonitor(list[i])
		valueBytesJSON, _ := json.Marshal(res)

		if err := json.Unmarshal(valueBytesJSON, meh); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}
		monitorList := []string{
			meh.ID,
			meh.Name,
			meh.Service,
			fmt.Sprint(*meh.ResponseTimeWarning),
			fmt.Sprint(*meh.ResponseTimeCritical),
			fmt.Sprint(*meh.ResponseTimeDuration),
			fmt.Sprint(meh.MaxCheckAttempts),
			meh.Memo,
		}
		monitorLists = append(monitorLists, monitorList)
	}
	return monitorLists
}
