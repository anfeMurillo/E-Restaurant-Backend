package payment

import (
	paymentstatus "e-restaurant/models/enums/paymentStatus"
	"e-restaurant/models/payment"
)

type Repository interface {
	Create(payment.Payment)

	GetById(paymentId int)

	UpdateStatus(paymentId int, status paymentstatus.PaymentStatus)

	UpdateMethod(paymentId int, method paymentstatus.PaymentMethod)
}
