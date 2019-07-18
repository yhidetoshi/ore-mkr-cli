package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
	"time"
)

const (
	USER = "user"
)

type UserValues struct {
	ID         string `json:"id,omitempty"`
	ScreenName string `json:"screenName,omitempty"`
	Email      string `json:"email,omitempty"`
	Authority  string `json:"email,omitempty"`

	IsInRegistrationProcess bool     `json:"isInRegistrationProcess,omitempty"`
	IsMFAEnabled            bool     `json:"isMFAEnabled,omitempty"`
	AuthenticationMethods   []string `json:"authenticationMethods,omitempty"`
	JoinedAt                int64    `json:"joinedAt,omitempty"`
}

func FetchUsers(client *mackerel.Client) {
	userLists := [][]string{}
	users, err := client.FindUsers()
	if err != nil {
		fmt.Println("fail get users")
	}

	for _, v := range users {

		authMethod := strings.Join(v.AuthenticationMethods, ",")
		userList := []string{
			v.ID,
			v.ScreenName,
			//v.Email,
			v.Authority,
			authMethod,
			fmt.Sprint(v.IsMFAEnabled),
			fmt.Sprint(v.IsInRegistrationProcess),
			fmt.Sprint(time.Unix(v.JoinedAt, 0)),
		}

		userLists = append(userLists, userList)
	}
	OutputFormat(userLists, USER)
}
