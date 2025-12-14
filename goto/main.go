package main

import "fmt"

const (
	OrderStatusDraft = "draft"
	OrderStatusCompleted = "completed"
	OrderStatusPaymentPending = "payment_pending"
)

func main(){
	fmt.Println("Goto statement")
	var status string = OrderStatusDraft
	if status == OrderStatusDraft {
		goto Draft;
	} 
	if status == OrderStatusPaymentPending {
		goto Payment;
	}
	if status == OrderStatusCompleted {
		goto Completed;
	}

	Draft:
		fmt.Println("Order is in draft state. please jump into payment process");
	Payment:
		fmt.Println("Order is in payment pending state");
	Completed:
		fmt.Println("Order is in completed state.")
}