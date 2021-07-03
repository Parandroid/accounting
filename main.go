package main

import (
	"fmt"
	"math/rand"
	"parandroid/accounting/classes"
	"parandroid/accounting/classes/journals"
	"time"

	"github.com/shopspring/decimal"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	journals.CashAccount = classes.NewAccount(1, "CASH")

	const ACC_COUNT = 10
	var accounts [ACC_COUNT]*classes.Account

	for i := 0; i < ACC_COUNT; i++ {
		accounts[i] = classes.NewAccount(i+1, "PERSONAL")
	}

	fmt.Println("DEPOSITS")
	for i := 0; i < 20; i++ {
		var dep = journals.NewDeposit(accounts[rand.Intn(ACC_COUNT)], decimal.NewFromInt(int64(rand.Intn(1000))), "EUR")
		dep.Post()
		fmt.Println(dep)
	}

	fmt.Println()
	fmt.Println("TRANSFERS")
	for i := 0; i < 30; i++ {
		var t = journals.NewTransfer(accounts[rand.Intn(ACC_COUNT)], accounts[rand.Intn(ACC_COUNT)], decimal.NewFromInt(int64(rand.Intn(500))), "EUR")
		t.Post()
		fmt.Println(t)
	}

	fmt.Println()
	fmt.Println("WITHDRAWALS")
	for i := 0; i < 40; i++ {
		var w = journals.NewWithdrawal(accounts[rand.Intn(ACC_COUNT)], decimal.NewFromInt(int64(rand.Intn(400))), "EUR")
		w.Post()
		fmt.Println(w)
	}

	fmt.Println()
	fmt.Println("POSTINGS")

	for _, v := range classes.Postings {
		fmt.Println(v)
	}

	fmt.Println()
	fmt.Println("ACCOUNT BALANCE")
	fmt.Println(journals.CashAccount, journals.CashAccount.Balance())
	for i := 0; i < ACC_COUNT; i++ {
		fmt.Println(accounts[i], accounts[i].Balance())
	}

}
