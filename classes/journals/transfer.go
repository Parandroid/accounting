package journals

import (
	"parandroid/accounting/classes"

	"fmt"

	"github.com/shopspring/decimal"
)

type Transfer struct {
	JournalEntry

	accountFrom *classes.Account
	accountTo   *classes.Account
	amount      decimal.Decimal
	currency    string
}

func NewTransfer(accountFrom *classes.Account, accountTo *classes.Account, amount decimal.Decimal, currency string) *Transfer {
	p := new(Transfer)
	p.JournalEntry.SetId()
	p.accountFrom = accountFrom
	p.accountTo = accountTo
	p.amount = amount
	p.currency = currency

	return p
}

func (transfer *Transfer) Post() (debit, credit *classes.Posting) {
	debit = classes.NewPosting(transfer.accountFrom, transfer.amount.Neg(), transfer.currency)
	credit = classes.NewPosting(transfer.accountTo, transfer.amount, transfer.currency)

	return
}

func (transfer *Transfer) String() string {
	s := fmt.Sprintf("id:%d accountFrom:%d accountTo:%d amount:%s currency:%s",
		transfer.id,
		transfer.accountFrom.Id(),
		transfer.accountTo.Id(),
		transfer.amount.String(),
		transfer.currency)
	return s
}
