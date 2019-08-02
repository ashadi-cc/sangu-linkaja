package linkaja

type PublicTokenResponse struct {
	PgpToken string `json:"pgpToken"`
	RefNum   string `json:"refNum"`
	FastTime uint64 `json:"fastTime"`
}
