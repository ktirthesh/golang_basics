package validator

import (
	"errors"

	"github.com/ktirthesh/basepayment/payment"
)

type FraudValidator struct {
	BaseValidator
}

func (v *FraudValidator) Validate(req payment.PaymentLinkRequest) error {
	blackListedMail := "fraud@test.com"
	if req.Email == blackListedMail {
		return errors.New("Fraudlent email detected")
	}
	return v.Next(req)
}
