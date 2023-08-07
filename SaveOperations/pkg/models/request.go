package models

type Operation struct {
	OperationCode      string            `json:"operation_code"`
	DebtorAccountCode  string            `json:"debtor_account_code"`
	DebtorEntityCode   string            `json:"debtor_entity_code"`
	CreditoAccountCode string            `json:"creditor_account_code"`
	CreditorEntityCode string            `json:"creditor_entity_code"`
	Amount             float32           `json:"amount"`
	InternalNumber     string            `json:"internal_number,omitempty"`
	PaymentInstruction map[string]string `json:"payment_instruction"`
}

type Request struct {
	Data []Operation `json:"data"`
}
