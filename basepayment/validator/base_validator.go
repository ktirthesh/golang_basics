package validator

import "github.com/ktirthesh/basepayment/payment"

type BaseValidator struct {
	next Validator
}

func (b *BaseValidator) SetNext(v Validator) {
	b.next = v
}

func (b *BaseValidator) Next(req payment.PaymentLinkRequest) error {
	if b.next != nil {
		return b.next.Validate(req)
	}
	return nil
}
