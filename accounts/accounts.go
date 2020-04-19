package accounts

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

//Constructor Account
func NewAccount(owner string) *Account {

	account := Account{owner: owner, balance: 0}
	return &account
}

//ADD Deposit amount in balance
func (a *Account) Deposit(amount int) {

	a.balance += amount
	fmt.Println(amount, "원 입금완료 !")
}

//Account in Balance
func (a Account) Balanace() {
	fmt.Println("계좌조회 : ", a.balance)
}

//Delete int Balance
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return errors.New("계좌에서 출금할 금액이 계좌잔액보다 많습니다.")
	} else {
		a.balance -= amount
		fmt.Println("계좌에서 ", amount, "원을 출금했습니다.")
		a.Balanace()
	}

	return nil
}
