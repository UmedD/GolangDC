package main

import "fmt"

type BankAccount interface {
	Deposit(amount float64)
	Withdraw(amount float64)
	Balance() float64
}

type SavingsAccount struct {
	amount float64
}

func (s *SavingsAccount) Deposit(amount float64) {
	interest := amount * 0.02
	s.amount += amount + interest
}

func (s *SavingsAccount) Withdraw(amount float64) {
	if s.amount < amount {
		fmt.Println("Недостаточно средств на сберегательном счете")
		return
	}
	s.amount -= amount
}

func (s *SavingsAccount) Balance() float64 {
	fmt.Println("проверка баланса: Saving")
	return s.amount
}

type CheckingAccount struct {
	amount float64
}

func (c *CheckingAccount) Deposit(amount float64) {
	c.amount += amount
}

func (c *CheckingAccount) Withdraw(amount float64) {
	limit := -100.0
	if c.amount-amount < limit {
		fmt.Println("Превышен лимит на текущем счете")
		return
	}
	c.amount -= amount
}

func (c *CheckingAccount) Balance() float64 {
	fmt.Println("проверка баланса: Check")
	return c.amount
}

func main() {
	acc := []BankAccount{
		&SavingsAccount{
			amount: 40,
		},
		&CheckingAccount{
			amount: 50,
		},
	}

	for _, v := range acc {
		v.Deposit(50)
		v.Withdraw(10)
		fmt.Println(v.Balance())
	}
}
