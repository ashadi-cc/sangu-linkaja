package linkaja

type PublicTokenRequest struct {
	TrxId           string
	Total           string
	SuccessUrl      string
	FailedUrl       string
	Items           []PublicTokenItemRequest
	MSISDN          string
	DefaultLanguage string
	DefaultTemplate string
}

type PublicTokenItemRequest struct {
	Name     string
	Price    string
	Quantity string
}

type CheckTransactionStatusRequest struct {
	RefNum string
}
