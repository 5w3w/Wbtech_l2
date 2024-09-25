package main

import (
	"fmt"

	"wbtechl2/pattern/facade/facade1"
)

var (
	bank = facade1.Bank{
		Name:  "Bank",
		Cards: []facade1.Card{},
	}
	card1 = facade1.Card{
		Name:    "CRD - 1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = facade1.Card{
		Name:    "CRD - 2",
		Balance: 5,
		Bank:    &bank,
	}
	user = facade1.User{
		Name: "Покупатель - 1",
		Card: &card1,
	}
	user2 = facade1.User{
		Name: "Покупатель - 2",
		Card: &card2,
	}
	prod = facade1.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = facade1.Shop{
		Name: "SHOP",
		Product: []facade1.Product{
			prod,
		},
	}
)

func main() {
	fmt.Println("[Банк] Выпуск кард")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
