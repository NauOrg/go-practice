package main

import "fmt"

type PaymentGatway interface {
	Pay(amount float64) error
}

type PayTm struct{}

func (p PayTm) Pay(amount float64) error {
	fmt.Println("Paid with Paytm", amount)
	return nil
}

type Strip struct{}

func (p Strip) Pay(amount float64) error {
	fmt.Println("Paid with Stripe", amount)
	return nil
}

func Checkout(paymentGatway PaymentGatway) {
	paymentGatway.Pay(500)
}

func Identify(x interface{}) {
	switch v := x.(type) {
	case PayTm:
		fmt.Println("PayTm:Identify")
		v.Pay(700)
	case Strip:
		fmt.Println("Strip:Identify")
		v.Pay(700)
	}
}

func main() {

	Identify(PayTm{})
	Identify(Strip{})
}
