package facade1

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name    string
	Product []Product
}

func (shop Shop) Sell(user User, Product string) error {
	fmt.Println("[Магазин] Запрос к пользователю для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Маганзин] Промерка - может ли [%s] купить товар\n", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Product {
		if prod.Name != Product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] Недостаточно средств для покупки товара!")
		}
		fmt.Printf("[Маганзин] Товар [%s] - куплен \n", prod.Name)

	}
	return nil
}
