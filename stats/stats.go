package stats

import (
	"github.com/RakhimovAns/bank/v2/pkg/types"
)

func AVG(payments []types.Payment) types.Money {
	count := 0
	var sum types.Money
	sum = types.Money(0)
	for _, card := range payments {
		if card.Status != types.StatusFail && card.Amount > types.Money(0) {
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
		if card.Category == category && card.Status != types.StatusFail {
			sum += card.Amount
		}
	}
	return types.Money(sum)
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	count := map[types.Category]types.Money{}
	sum := map[types.Category]types.Money{}
	for _, payment := range payments {
		count[payment.Category] += types.Money(1)
		sum[payment.Category] += payment.Amount
	}
	answer := map[types.Category]types.Money{}
	for key := range sum {
		answer[key] = sum[key] / count[key]
	}
	return answer
}
