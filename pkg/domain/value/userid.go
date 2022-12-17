package value

import (
	"fmt"
	"regexp"
)

const UserIdFormat = "^[0-9a-zA-Z]{40}$"

type UserId struct {
	Id string
}

func NewUserId(id string) (UserId, error) {
	u := UserId{id}
	return u, u.Validate()
}

func (u UserId) Validate() error {
	matched, err := regexp.MatchString(UserIdFormat, u.Id)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("bad format")
	}
	return nil
}
