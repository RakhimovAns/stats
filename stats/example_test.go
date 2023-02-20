package stats

import (
	"fmt"
	"github.com/RakhimovAns/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func ExampleAVG() {
	payments := []types.Payment{
		{
			Amount: 100,
		},
		{
			Amount: 100,
		},
	}
	fmt.Println(AVG(payments))
	//Output:
	//100
}
func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			Category: "Ansar",
			Amount:   100,
		},
		{
			Category: "Rakhimov",
			Amount:   100,
		},
	}
	fmt.Println(TotalInCategory(payments, "Ansar"))
	//Output:
	//100
}

func TestCategoriesAvg_notFound(t *testing.T) {
	arr := []types.Payment{
		{
			Category: "fun", Amount: 100,
		},
		{
			Category: "auto", Amount: 100,
		},
		{
			Category: "fun", Amount: 200,
		},
	}
	result := CategoriesAvg(arr)
	if len(result) != 0 {
		t.Error("result != 0")
	}
}
func TestCategoriesAvg_nill(t *testing.T) {
	var arr []types.Payment
	result := CategoriesAvg(arr)
	if len(result) != 0 {
		t.Error("result != 0")
	}
}
func TestCategoriesAvg_foundOne(t *testing.T) {
	arr := []types.Payment{
		{
			Category: "fun", Amount: 100,
		},
	}
	var expected = map[types.Category]types.Money{
		"fun": 100,
	}
	result := CategoriesAvg(arr)
	if reflect.DeepEqual(expected, result) == false {
		t.Error("not what expected")
	}
}

func TestCategoriesAvg_foundMultiple(t *testing.T) {
	arr := []types.Payment{
		{
			Category: "fun", Amount: 100,
		},
		{
			Category: "auto", Amount: 100,
		},
	}
	var expected = map[types.Category]types.Money{
		"fun":  100,
		"auto": 100,
	}
	result := CategoriesAvg(arr)
	if reflect.DeepEqual(expected, result) == false {
		t.Error("not what expected")
	}
}
