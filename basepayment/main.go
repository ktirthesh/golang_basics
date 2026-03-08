package main

import (
	"fmt"

	"github.com/ktirthesh/basepayment/payment"
	"github.com/ktirthesh/basepayment/validator"
)

func buildValidationChain() validator.Validator {
	name := &validator.NameValidator{}
	amount := &validator.AmountValidator{}
	email := &validator.EmailValidator{}
	phone := &validator.PhoneValidator{}
	merchant := &validator.MerchantValidator{}
	fraud := &validator.FraudValidator{}

	name.SetNext(amount)
	amount.SetNext(email)
	email.SetNext(phone)
	phone.SetNext(merchant)
	merchant.SetNext(fraud)
	return name
}
func main() {
	req := payment.PaymentLinkRequest{
		Name:   "sample",
		Email:  "sdsdf@gmail.com",
		Phone:  "9922180669",
		Amount: 10.00,
	}
	validatorChain := buildValidationChain()
	err := validatorChain.Validate(req)
	if err != nil {
		fmt.Println("validation failed:", err)
		return
	}
	fmt.Println("payment link validation success")
}
