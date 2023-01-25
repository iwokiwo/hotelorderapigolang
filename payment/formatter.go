package payment

type PaymentTypeFormatter struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Slug   string `json:"slug"`
	Branch Branch `json:"branch"`
}

func FormatPaymentType(paymentType PaymentType) PaymentTypeFormatter {
	formatter := PaymentTypeFormatter{
		ID:     paymentType.ID,
		Name:   paymentType.Name,
		Branch: paymentType.Branch,
	}

	return formatter
}

func FormatCategories(category []PaymentType) []PaymentTypeFormatter {
	PaymentTypeFormatter := []PaymentTypeFormatter{}
	for _, category := range category {
		formatter := FormatPaymentType(category)
		PaymentTypeFormatter = append(PaymentTypeFormatter, formatter)
	}
	return PaymentTypeFormatter
}
