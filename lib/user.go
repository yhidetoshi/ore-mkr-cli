package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
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
	//fmt.Println(users)
	for _, v := range users {
		userList := []string{
			v.ID,
			v.ScreenName,
			v.Email,
			v.Authority,
			//v.AuthenticationMethods,
			fmt.Sprint(time.Unix(v.JoinedAt, 0)),
		}

		userLists = append(userLists, userList)
	}
	OutputFormat(userLists, USER)
}
