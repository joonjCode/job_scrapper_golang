package accounts

// Account struct
type account struct {
	owner   string
	balance int
}

// NewAccount create Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit money
func (a account) Deposit(amount int) {
	a.balance += amount
}

// Blance of your account
func (a account) Balance() int {
	return a.balance
}
