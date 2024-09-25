package facade1

import (
	"errors"
	"fmt"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank Bank) RequestBalance(cardNumber string) error {
	fmt.Println(fmt.Sprintf("[Банк] Получение остатка по карте %s", cardNumber))
	time.Sleep(time.Millisecond * 800)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	fmt.Println("[Банк] Остаток положительный")
	return nil
}
