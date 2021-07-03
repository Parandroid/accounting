package classes

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var idCounterAccount int = 1000000000

type Account struct {
	id      int
	owner   int
	accType string
}

func NewAccount(owner int, accType string) *Account {
	a := new(Account)
	a.id = idCounterAccount
	idCounterAccount++
	a.owner = owner
	a.accType = accType

	return a
}

func (a *Account) Id() int {
	return a.id
}

func (a Account) String() string {
	return fmt.Sprintf("%d(%s)", a.id, a.accType)
}

func (a *Account) Balance() (balance decimal.Decimal) {
	for _, v := range Postings {
		if v.account == a {
			balance = balance.Add(v.amount)
		}
	}
	return
}
