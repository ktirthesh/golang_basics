package validator

import (
	"errors"

	"github.com/ktirthesh/basepayment/payment"
)

type AmountValidator struct {
	BaseValidator
}

func (v *AmountValidator) Validate(req payment.PaymentLinkRequest) error {
	if req.Amount <= 0 {
		return errors.New("invalid amount")
	}
	if req.Amount > 100000 {
		return errors.New("Amount Exceded limit ")
	}
	return v.Next(req)
}
