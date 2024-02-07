# Before (Cyclomatic Complexity of processPayment = 7)

~~~
func processPayment(order *Order) {
	if order == nil {
		fmt.Println("Order is nil, cannot process payment.")
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
~~~

# After (Cyclomatic Complexity of processPayment = 2 )

~~~
func processPayment(order Order) Order {
	if order.Status != PENDING {
		fmt.Println("Order is not pending, cannot process payment.")
		return order
	}

	return processMoney(order)
}

func processMoney(order Order) Order {
	if order.PaymentMethod == CREDIT_CARD {
		return processCreditCard(order)
	}

	if order.PaymentMethod == PAYPAL {
		return processPayPall(order)
	}

	order.Status = FAILED
	sendPaymentFailureNotification(&order.User)

	return order
}

func processCreditCard(order Order) Order {
	if validateCreditCard(&order.CreditCard) {
		chargeCreditCard(&order.CreditCard, order.Amount)
		order.Status = PAID
		return order
	}

	order.Status = FAILED
	sendPaymentFailureNotification(&order.User)
	return order
}

func processPayPall(order Order) Order {
	if validatePayPalAccount(&order.PayPalAccount) {
		transferPayPalFunds(&order.PayPalAccount, order.Amount)
		order.Status = PAID
		return order
	}

	order.Status = FAILED
	sendPaymentFailureNotification(&order.User)
	return order
}
~~~

## Methods used

- removed pointer and nil check
- removed else cases
- removed switch/cases 
- encapsulated processing logic into separated functions