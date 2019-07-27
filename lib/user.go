package oremkrcli

import (
	"fmt"
	"github.com/mackerelio/mackerel-client-go"
	"strings"
	"time"
)

const (
	// USER user
	USER = "user"
)

// UserValues information
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

// FetchUsers fetch users
func FetchUsers(client *mackerel.Client) error {
	userLists := [][]string{}
	users, err := client.FindUsers()
	if err != nil {
		fmt.Println(err)
		return err
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
	err = OutputFormat(userLists, USER)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeleteUser(client *mackerel.Client, userID string) error {
	user, err := client.DeleteUser(userID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Success Deleted %s\n", user.Email)

	return nil
}
