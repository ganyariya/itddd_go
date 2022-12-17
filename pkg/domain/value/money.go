package value

import "fmt"

type Money struct {
	amount int
	kind   string
}

func NewMoney(amount int, kind string) Money {
	return Money{amount, kind}
}

func (m Money) Add(o Money) (Money, error) {
	if m.kind != o.kind {
		return Money{}, fmt.Errorf("different kind")
	}
	return NewMoney(m.amount+o.amount, m.kind), nil
}
