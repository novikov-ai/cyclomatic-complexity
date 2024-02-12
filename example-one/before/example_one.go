package main

import (
	"fmt"
)

type OrderStatus int

const (
	PENDING OrderStatus = iota
	PAID
	FAILED
)

type PaymentMethod int

const (
	CREDIT_CARD PaymentMethod = iota
	PAYPAL
)

type User struct {
	Name  string
	Email string
}

type CreditCard struct {
	Number         string
	ExpirationDate string
	CVV            int
}

type PayPalAccount struct {
	Email    string
	Password string
}

type Order struct {
	ID            int
	User          User
	Amount        float64
	Status        OrderStatus
	PaymentMethod PaymentMethod
	CreditCard    CreditCard
	PayPalAccount PayPalAccount
}

func processPayment(order *Order) {
	if order == nil {
		fmt.Println("Order is nil, cannot pwrocess payment.")
		return
	}
	if order.Status != PENDING {
		fmt.Println("Order is not pending, cannot process payment.")
		return
	}
	switch order.PaymentMethod {
	case CREDIT_CARD:
		if validateCreditCard(&order.CreditCard) {
			chargeCreditCard(&order.CreditCard, order.Amount)
			order.Status = PAID
		} else {
			order.Status = FAILED
			sendPaymentFailureNotification(&order.User)
		}
	case PAYPAL:
		if validatePayPalAccount(&order.PayPalAccount) {
			transferPayPalFunds(&order.PayPalAccount, order.Amount)
			order.Status = PAID
		} else {
			order.Status = FAILED
			sendPaymentFailureNotification(&order.User)
		}
	default:
		order.Status = FAILED
		sendPaymentFailureNotification(&order.User)
	}
}

func validateCreditCard(creditCard *CreditCard) bool {
	// Validation logic for credit card
	return true // Placeholder for simplicity
}

func chargeCreditCard(creditCard *CreditCard, amount float64) {
	// Logic to charge credit card
}

func validatePayPalAccount(payPalAccount *PayPalAccount) bool {
	// Validation logic for PayPal account
	return true // Placeholder for simplicity
}

func transferPayPalFunds(payPalAccount *PayPalAccount, amount float64) {
	// Logic to transfer funds via PayPal
}

func sendPaymentFailureNotification(user *User) {
	// Logic to send payment failure notification to user
	fmt.Printf("Payment failure notification sent to %s <%s>\n", user.Name, user.Email)
}

func main() {
	order := &Order{
		ID:            1,
		User:          User{Name: "John Doe", Email: "john@example.com"},
		Amount:        100.0,
		Status:        PENDING,
		PaymentMethod: PAYPAL,
		CreditCard:    CreditCard{},
		PayPalAccount: PayPalAccount{Email: "john@example.com", Password: "password"},
	}
	processPayment(order)
}
