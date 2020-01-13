package linkaja

type PublicTokenResponse struct {
	PgpToken string `json:"pgpToken"`
	RefNum   string `json:"refNum"`
	FastTime uint64 `json:"fastTime"`
}

type TransactionResponses struct {
	RefNum          string `json:"refNum"`
	Amount          string `json:"amount"`
	TransactionDate string `json:"transactionDate"`
	Status          string `json:"status"`
}

type GenerateApplinkPaymentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	URL string `json:"url"`
}
