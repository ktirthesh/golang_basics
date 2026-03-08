package validator

import (
	"errors"

	"github.com/ktirthesh/basepayment/payment"
)

type MerchantValidator struct {
	BaseValidator
}

func (v *MerchantValidator) Validate(req payment.PaymentLinkRequest) error {
	merchantActive := true
	maxLimit := 50000.0
	if !merchantActive {
		return errors.New("merchant disable")
	}
	if req.Amount > maxLimit {
		return errors.New("merchant limit exceeded")
	}
	return v.Next(req)
}
