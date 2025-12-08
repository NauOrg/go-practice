package main

import "fmt"

type Logger struct{}

func (Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

type Payment interface {
	Pay(amount int)
}
type Paytm struct {
	Logger
}

type Stripe struct {
	Logger
}

func (Paytm) Pay(amount int) {
	fmt.Println("🔹 Paid via Paytm:", amount)
}

func (Stripe) Pay(amount int) {
	fmt.Println("🔹 Paid via Stripe:", amount)
}
func Process(p Payment, amount int) {
	p.Pay(amount)
}

func main() {
	p1 := Paytm{}
	p2 := Stripe{}

	Process(p1, 300)
	Process(p2, 800)

	// Logger automatically available due to embedding
	p1.Log("Payment completed")
	p2.Log("Payment failed due to insufficient balance")
}
