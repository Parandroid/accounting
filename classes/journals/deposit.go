package journals

import (
	"parandroid/accounting/classes"

	"fmt"

	"github.com/shopspring/decimal"
)

var CashAccount *classes.Account

type Deposit struct {
	JournalEntry

	account  *classes.Account
	amount   decimal.Decimal
	currency string
}

func NewDeposit(account *classes.Account, amount decimal.Decimal, currency string) *Deposit {
	p := new(Deposit)
	p.JournalEntry.SetId()
	p.account = account
	p.amount = amount
	p.currency = currency

	return p
}

func (dep *Deposit) Post() (debit, credit *classes.Posting) {
	debit = classes.NewPosting(CashAccount, dep.amount.Neg(), dep.currency)
	credit = classes.NewPosting(dep.account, dep.amount, dep.currency)

	return
}

func (dep *Deposit) String() string {
	s := fmt.Sprintf("id:%d account!:%d amount:%s currency:%s", dep.id, dep.account.Id(), dep.amount.String(), dep.currency)
	return s
}
