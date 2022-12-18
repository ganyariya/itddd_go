package value

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

const UserIdFormat = "^[0-9a-zA-Z-]{36}$"

type UserId struct {
	Id string
}

func NewUserId() (UserId, error) {
	uuid, _ := uuid.NewUUID()
	u := UserId{Id: uuid.String()}
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
