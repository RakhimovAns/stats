package stats

import (
	"fmt"
	"github.com/RakhimovAns/bank/v2/pkg/types"
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
