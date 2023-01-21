package stats

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/Aziz-Gafuroff/banknew/v2/pkg/types"
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

	// Output: 3000000
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


func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 2_000_00},
		{ID: 2, Category: "food", Amount: 2_000_00},
		{ID: 3, Category: "auto", Amount: 3_000_00},
		{ID: 4, Category: "auto", Amount: 4_000_00},
		{ID: 5, Category: "fun", Amount: 5_000_00},
	}
	expected := map[types.Category]types.Money {
		"auto": 3_000_00,
		"food": 2_000_00,
		"fun": 5_000_00,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}
}