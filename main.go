package main

import (
	"fmt"

	"github.com/joon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("joon")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
