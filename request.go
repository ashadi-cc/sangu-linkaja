package linkaja

type PublicTokenRequest struct {
	TrxId           string
	Total           string
	SuccessUrl      string
	FailedUrl       string
	Items           string
	MSISDN          string
	DefaultLanguage string
	DefaultTemplate string
}

type CheckTransactionStatusRequest struct {
	RefNum string
}
