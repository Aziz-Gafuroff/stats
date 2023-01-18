package stats

import (
	"github.com/Aziz-Gafuroff/banknew/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 98,
			Amount: 20_000_00,
			Category: "food",
			Status: "OK",
		},
		{
			ID: 99,
			Amount: 30_000_00,
			Category: "auto",
			Status: "INPROGRESS",
		},
		{
			ID: 100,
			Amount: 10_000_00,
			Category: "mobile",
			Status: "FAIL",
		},
		{
			ID: 101,
			Amount: 40_000_00,
			Category: "auto",
			Status: "OK",
		},
	}
	avgPayment := Avg(payments)
	fmt.Println(avgPayment)

	// Output: 2250000
}


func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 98,
			Amount: 20_000_00,
			Category: "food",
			Status: "OK",
		},
		{
			ID: 99,
			Amount: 30_000_00,
			Category: "auto",
			Status: "INPROGRESS",
		},
		{
			ID: 100,
			Amount: 10_000_00,
			Category: "mobile",
			Status: "FAIL",
		},
		{
			ID: 101,
			Amount: 40_000_00,
			Category: "auto",
			Status: "OK",
		},
		{
			ID: 102,
			Amount: 5_000_00,
			Category: "mobile",
			Status: "OK",
		},
	}
	sumCategory := TotalInCategory(payments, "mobile")
	fmt.Println(sumCategory)

	// Output: 500000
}