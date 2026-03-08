package validator

import (
	"errors"
	"strings"

	"github.com/ktirthesh/basepayment/payment"
)

type EmailValidator struct {
	BaseValidator
}

func (v *EmailValidator) Validate(req payment.PaymentLinkRequest) error {
	if !strings.Contains(req.Email, "@") {
		return errors.New("Invalid Email")
	}
	return v.Next(req)
}
