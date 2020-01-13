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

type TransactionRequest struct {
	RefNum string
}

type GenerateApplinkPaymentRequest struct {
	TrxDate        string `json:"trxDate"`
	PartnerTrxID   string `json:"partnerTrxID"`
	MerchantID     string `json:"merchantID"`
	TerminalID     string `json:"terminalID"`
	TotalAmount    string `json:"totalAmount"`
	PartnerApplink string `json:"partnerApplink"`
	Items          []Item `json:"items"`
}

type Item struct {
	Name      string `json:"name"`
	UnitPrice string `json:"unitPrice"`
	Quantity  string `json:"qty"`
}
