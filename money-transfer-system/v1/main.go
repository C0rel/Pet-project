package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return errors.New("Недостаточно средств")
	}
	u.Balance -= amount
	return nil
}

func main() {
	amount:=15200.0
	user1:=User{ID:"1",Name: "Pavel",Balance: 12500}
	user2:=User{ID:"2",Name: "Ivan",Balance: 15500}
	user3:=User{ID:"3",Name: "Danil",Balance: 15000.50}

	user1.Deposit(amount)
	fmt.Println("Баланс:",user1.Balance)

	err:= user2.Withdraw(amount)
	if err!=nil{
		fmt.Println("Ошибка:" , err)
	}else{
		fmt.Println("Баланс:",user2.Balance)
	}

	err1:= user3.Withdraw(amount)
		if err1!=nil{
		fmt.Println("Ошибка:" , err1)
	}else{
		fmt.Println("Баланс:",user3.Balance)
	}
}