package journals

import (
	"parandroid/accounting/classes"

	"fmt"

	"github.com/shopspring/decimal"
)

type Withdrawal struct {
	JournalEntry

	account  *classes.Account
	amount   decimal.Decimal
	currency string
}

func NewWithdrawal(account *classes.Account, amount decimal.Decimal, currency string) *Withdrawal {
	p := new(Withdrawal)
	p.JournalEntry.SetId()
	p.account = account
	p.amount = amount
	p.currency = currency

	return p
}

func (w *Withdrawal) Post() (debit, credit *classes.Posting) {
	debit = classes.NewPosting(w.account, w.amount.Neg(), w.currency)
	credit = classes.NewPosting(CashAccount, w.amount, w.currency)

	return
}

func (w *Withdrawal) String() string {
	s := fmt.Sprintf("id:%d account:%d amount:%s currency:%s", w.id, w.account.Id(), w.amount.String(), w.currency)
	return s
}
