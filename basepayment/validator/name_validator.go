package validator

import (
	"errors"

	"github.com/ktirthesh/basepayment/payment"
)

type NameValidator struct {
	BaseValidator
}

func (v *NameValidator) Validate(req payment.PaymentLinkRequest) error {
	if req.Name == "" {
		return errors.New("Name required")
	}
	return v.Next(req)
}
