package bdd_in_go

import (
	"errors"
	"github.com/DATA-DOG/godog"
)

var testAccount *account

func iHaveABankAccountWith(balance int) error {
	testAccount = &account{balance: balance}

	return nil
}

func iDeposit(amount int) error {
	testAccount.deposit(amount)

	return nil
}

func itShouldHaveABalanceOf(balance int) error {
	if testAccount.balance == balance {
		return nil
	}
	return errors.New("неправильный баланс счета")
}

func iWithdraw(amount int) error {
	testAccount.withdraw(amount)

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^У меня есть банковский счет на (\d+)\$$`, iHaveABankAccountWith)
	s.Step(`^Я вкладываю (\d+)\$$`, iDeposit)
	s.Step(`^Я снимаю (\d+)\$$`, iWithdraw)
	s.Step(`^бланс должен быть (\d+)\$$`, itShouldHaveABalanceOf)
}
