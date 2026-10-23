package main

import (
	"errors"
	"fmt"
)

type Transaction struct{
	FromID string
	ToID string
	Amount float64
}


type PaymentSystem struct{
	Users map[string]*User
	TransactionQueue []Transaction
}

func (ps *PaymentSystem) AddTransaction(t Transaction){
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem)AddUser(user *User){
	ps.Users[user.ID]=user
}

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (ps *PaymentSystem)ProcessingTransactions(t Transaction)error{
	 	fromUser,ok:=ps.Users[t.FromID]
		if !ok{
			return errors.New("Отправителя нет")
		}else{
			if fromUser.Balance>=t.Amount{
			fromUser.Withdraw(t.Amount)
			}else{
				return errors.New("Недостаточно средств")
			}
		}

		toID,ok1:=ps.Users[t.ToID]
		if !ok1{
			return errors.New("Получателя нет")
		}else{
			toID.Deposit(t.Amount)
		}
		return nil
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
	ps := &PaymentSystem{
    Users:            make(map[string]*User),
    TransactionQueue: []Transaction{},
	}

	user1:=&User{ID:"1",Name: "Pavel",Balance: 12500}
	user2:=&User{ID:"2",Name: "Ivan",Balance: 15500}
	user3:=&User{ID:"3",Name: "Danil",Balance: 15000.50}
	
	ps.AddUser(user1)
	ps.AddUser(user2)
	ps.AddUser(user3)

	t1:=Transaction{"1","2",100.0}
	t2:=Transaction{"1","3",200.0}
	t3:=Transaction{"2","3",300.0}

	ps.AddTransaction(t1)
	ps.AddTransaction(t2)
	ps.AddTransaction(t3)

	for _,t := range ps.TransactionQueue{
		err:=ps.ProcessingTransactions(t)
		if err!=nil{
			fmt.Println("Ошибка:", err)
		}
	}
	for _,user := range ps.Users{
		fmt.Println(user.Balance)
	}
	
}