package stats

import (
	"github.com/Aziz-Gafuroff/banknew/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 98,
			Amount: 20_000_00,
			Category: "food",
		},
		{
			ID: 99,
			Amount: 30_000_00,
			Category: "auto",
		},
		{
			ID: 100,
			Amount: 10_000_00,
			Category: "mobile",
		},
		{
			ID: 101,
			Amount: 40_000_00,
			Category: "auto",
		},
	}
	avgPayment := Avg(payments)
	fmt.Println(avgPayment)

	// Output: 2500000
}


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 98,
			Amount: 20_000_00,
			Category: "food",
		},
		{
			ID: 99,
			Amount: 30_000_00,
			Category: "auto",
		},
		{
			ID: 100,
			Amount: 10_000_00,
			Category: "mobile",
		},
		{
			ID: 101,
			Amount: 40_000_00,
			Category: "auto",
		},
	}
	sumCategory := TotalInCategory(payments, "auto")
	fmt.Println(sumCategory)

	// Output: 7000000
}