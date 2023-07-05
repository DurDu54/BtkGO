package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for _, v := range credits {
		paymentTotal += float64(v.Calculate())
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{creditPaymentTotal: 100000, address: "ankara", rate: 10}
	credit2 := Mortgage{creditPaymentTotal: 100000, address: "ankara", rate: 10}
	credit3 := Car{creditPaymentTotal: 6000, rate: 12, carInfo: "polo"}
	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)
	fmt.Println("toplam odeme", total)
}
