package classes

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var idCounterPosting int
var Postings []*Posting

func init() {
	Postings = make([]*Posting, 0)
}

type Posting struct {
	id       int
	account  *Account
	amount   decimal.Decimal
	currency string
}

func NewPosting(account *Account, amount decimal.Decimal, currency string) *Posting {
	p := new(Posting)
	p.id = idCounterPosting
	idCounterPosting++

	p.account = account
	p.amount = amount
	p.currency = currency

	Postings = append(Postings, p)

	return p
}

func (p *Posting) String() string {
	s := fmt.Sprintf("id:%d account:%s amount:%s currency:%s", p.id, p.account.String(), p.amount.String(), p.currency)
	return s
}
