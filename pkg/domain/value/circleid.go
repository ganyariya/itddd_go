package value

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

const CircleIdFormat = "^[0-9a-zA-Z-]{36}$"

type CircleId struct {
	Id string
}

func NewCircleId() (CircleId, error) {
	uuid, _ := uuid.NewUUID()
	c := CircleId{Id: uuid.String()}
	return c, c.Validate()
}

func NewCircleIdByString(id string) (CircleId, error) {
	c := CircleId{Id: id}
	return c, c.Validate()
}

func (c CircleId) Validate() error {
	matched, err := regexp.MatchString(CircleIdFormat, c.Id)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("bad format")
	}
	return nil
}
