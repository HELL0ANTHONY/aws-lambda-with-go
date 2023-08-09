package models

type Operation struct {
	Amount             float32           `json:"amount"`
	CreditoAccountCode string            `json:"creditor_account_code"`
	CreditorEntityCode string            `json:"creditor_entity_code"`
	DebtorAccountCode  string            `json:"debtor_account_code"`
	DebtorEntityCode   string            `json:"debtor_entity_code"`
	InternalNumber     string            `json:"internal_number,omitempty"`
	OperationCode      string            `json:"operation_code"`
	PaymentInstruction map[string]string `json:"payment_instruction"`
}

type Request struct {
	Data []Operation `json:"data"`
}
