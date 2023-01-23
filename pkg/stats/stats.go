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

func PeriodDynamic(first map[types.Category]types.Money, second map[types.Category]types.Money) map[types.Category]types.Money {
	period := map[types.Category]types.Money{}
	
	for key, value := range second {
		_, ok := first[key]
		if !ok {
			period[key] = 0
		}
		period[key] += value - first[key]
	}

	for key, value := range first {
		_, ok := second[key]
		if !ok {
			period[key] = 0 - value
		}
	
	}

	return period
}