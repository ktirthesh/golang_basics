package payment

type PaymentLinkRequest struct {
	Name   string
	Amount float64
	Phone  string
	Email  string

	udf1 string
	udf2 string
	udf3 string
	udf4 string
	udf5 string
}
