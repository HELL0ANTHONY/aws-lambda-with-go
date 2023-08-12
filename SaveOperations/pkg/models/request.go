package models

type Operation struct {
	Amount             float64           `json:"amount"`
	CreditoAccountCode string            `json:"creditorAccountCode"`
	CreditorEntityCode string            `json:"creditorEntityCode"`
	DebtorAccountCode  string            `json:"debtorAccountCode"`
	DebtorEntityCode   string            `json:"debtorEntityCode"`
	InternalNumber     string            `json:"internalNumber,omitempty"`
	OperationCode      string            `json:"operationCode"`
	PaymentInstruction map[string]string `json:"paymentInstruction"`
}

type Request struct {
	Data []Operation `json:"data"`
}
