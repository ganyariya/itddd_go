package value

import "fmt"

type CircleName struct {
	Name string
}

func NewCircleName(name string) (CircleName, error) {
	c := CircleName{Name: name}
	return c, c.Validate()
}

func (c CircleName) Validate() error {
	if len(c.Name) < 3 || len(c.Name) > 20 {
		return fmt.Errorf("circle name format is invalid")
	}
	return nil
}
