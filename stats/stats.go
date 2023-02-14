package stats

import (
	"github.com/RakhimovAns/bank/pkg/types"
)

func AVG(payments []types.Payment) types.Money {
	count := 0
	var sum types.Money
	sum = types.Money(0)
	for _, card := range payments {
		if card.Amount > types.Money(0) {
			count++
			sum += card.Amount
		}
	}
	return sum / types.Money(count)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	sum = 0
	for _, card := range payments {
		if card.Category == category {
			sum += card.Amount
		}
	}
	return types.Money(sum)
}
