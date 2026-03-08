package validator

import (
	"errors"

	"github.com/ktirthesh/basepayment/payment"
)

type PhoneValidator struct {
	BaseValidator
}

func (v *PhoneValidator) Validate(req payment.PaymentLinkRequest) error {
	if len(req.Phone) != 10 {
		return errors.New("Invalid Phone")
	}
	return v.Next(req)
}
