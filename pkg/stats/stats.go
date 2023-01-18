package stats

import (
	"github.com/Aziz-Gafuroff/banknew/v2/pkg/types"
)


func Avg(payments []types.Payment) types.Money {
	avgPayment := types.Money(0)
	
	for _, payment := range payments {
		if payment.Status == types.StatusFail {
			continue
		}
		
		avgPayment += payment.Amount
		
	}

	return avgPayment / types.Money(len(payments))
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