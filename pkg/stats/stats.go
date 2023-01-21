package stats

import (
	"github.com/Aziz-Gafuroff/banknew/v2/pkg/types"
)


func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	var i int32
	
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		i++

		avgPayment += payment.Amount
		
	}

	return avgPayment / types.Money(i)
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumCategory := types.Money(0)

	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		if payment.Category == category {
			sumCategory += payment.Amount
		}
		
	}

	return sumCategory
}

func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	var counters = map[types.Category]int{}

	
	for _, payment := range payments {
		categories[payment.Category]+= payment.Amount
		counters[payment.Category] ++
	}

	for key, value := range counters {
		categories[key] = categories[key] / types.Money(value)
		
	}

	
	return categories
}