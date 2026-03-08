package validator

import (
	"github.com/ktirthesh/basepayment/payment"
)

type Validator interface {
	Validate(req payment.PaymentLinkRequest) error
	SetNext(Validator)
}
