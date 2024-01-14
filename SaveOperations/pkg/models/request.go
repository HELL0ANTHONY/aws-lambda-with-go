package models

type Operation struct {
	PaymentInstruction map[string]string `json:"paymentInstruction"`
	CreditoAccountCode string            `json:"creditorAccountCode"`
	CreditorEntityCode string            `json:"creditorEntityCode"`
	DebtorAccountCode  string            `json:"debtorAccountCode"`
	DebtorEntityCode   string            `json:"debtorEntityCode"`
	InternalNumber     string            `json:"internalNumber,omitempty"`
	OperationCode      string            `json:"operationCode"`
	Amount             float64           `json:"amount"`
}

type Request struct {
	Data []Operation `json:"data"`
}
