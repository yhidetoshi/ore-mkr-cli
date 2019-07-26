package oremkrcli

import (
	"encoding/json"
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"time"
)

const (
	//ALERT alert
	ALERT = "alert"
)

// AlertValues information
type AlertValues struct {
	ID        string  `json:"id,omitempty"`
	Status    string  `json:"status,omitempty"`
	MonitorID string  `json:"monitorId,omitempty"`
	Type      string  `json:"type,omitempty"`
	HostID    string  `json:"hostId,omitempty"`
	Value     float64 `json:"value,omitempty"`
	Message   string  `json:"message,omitempty"`
	Reason    string  `json:"reason,omitempty"`
	OpenedAt  int64   `json:"openedAt,omitempty"`
	ClosedAt  int64   `json:"closedAt,omitempty"`
}

// FetchOpenAlertIDs fetch alerts.
func FetchOpenAlertIDs(client *mackerel.Client) error {
	var listOpenAlerts [][]string
	var alertValues AlertValues

	res, err := client.FindAlerts()
	if err != nil {
		fmt.Println(err)
	}

	for _, alert := range res.Alerts {

		valueBytesJSON, _ := json.Marshal(res)
		bytesValue := []byte(valueBytesJSON)

		if err := json.Unmarshal(bytesValue, &alertValues); err != nil {
			fmt.Println("JSON Unmarshal error", err)
		}

		listOpenAlert := []string{
			alert.ID,
			alert.Status,
			fmt.Sprint(time.Unix(alert.OpenedAt, 0)),
			alert.Message,
		}
		listOpenAlerts = append(listOpenAlerts, listOpenAlert)
	}

	OutputFormat(listOpenAlerts, ALERT)

	return nil
}
