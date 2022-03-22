package main // import "hello-calc"

import (
	"fmt"
	"log"

	"github.com/mnogu/go-calculator"
	"github.com/shopspring/decimal"
)

func main() {
	val, err := calculator.Calculate("2.5 - 1.35 * 2.0")
	if err != nil {
		log.Fatal(err)
	}
	v := decimal.NewFromFloat(val).Round(5)
	fmt.Println(v) // 2.3

	val, err = calculator.Calculate("(2.5-1.35)*2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 2.3

	val, err = calculator.Calculate("-sin((-1+2.5)*pi)")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 1

	val, err = calculator.Calculate("180*atan2(log(e), log10(10))/pi")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 45
}
