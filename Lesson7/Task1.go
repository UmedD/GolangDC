package main

import "fmt"

type Account struct {
	Owner   string
	Balance float64
}

// Deposit добавляет деньги на счёт
// Используем указатель (pointer receiver), чтобы метод изменял исходный объект Account.
// Иначе изменения происходили бы только с копией и терялись бы.
func (o *Account) Deposit(amount float64) {
	o.Balance += amount
}

// Withdraw снимает деньги со счёта, если хватает средств
// Также используем указатель, потому что метод должен изменить баланс,
// а не просто работать с копией.
func (o *Account) Withdraw(amount float64) bool {
	if o.Balance != 0 && o.Balance > amount {
		o.Balance -= amount
		return true
	}
	return false
}

// GetBalance возвращает текущий баланс
// Здесь достаточно передавать по значению (value receiver), потому что метод ничего не меняет,
// он только читает данные. Передача по значению чуть эффективнее для небольших структур.
func (b Account) GetBalance() float64 {
	return b.Balance
}

func main() {
	acc1 := Account{
		Owner:   "Umed",
		Balance: 3500,
	}

	acc2 := Account{
		Owner:   "Rustam",
		Balance: 2500,
	}

	acc1.Deposit(1000)
	success := acc1.Withdraw(2000)
	acc1.GetBalance()

	acc2.Deposit(500)
	success2 := acc2.Withdraw(4000)
	acc2.GetBalance()

	fmt.Println("Was withdrawal for Acc1 successful?", success)
	fmt.Println("Final balance:", acc1.GetBalance())

	fmt.Println("Was withdrawal for Acc2 successful?", success2)
	fmt.Println("Final balance:", acc1.GetBalance())
}
