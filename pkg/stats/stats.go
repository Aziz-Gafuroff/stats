package stats

import (
	"github.com/Aziz-Gafuroff/banknew/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	
	for _, payment := range payments {
		avgPayment += payment.Amount
		
	}

	return avgPayment /types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	sumCategory := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sumCategory += payment.Amount
		}
		
	}

	return sumCategory
}