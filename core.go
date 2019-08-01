package linkaja

import (
	"io"
	"net/url"
	"strings"
)

const (
	PublicTokenRequestUrl = "linkaja-api/api/payment"
)

// CoreGateway struct
type CoreGateway struct {
	Client Client
}

// Call : base method to call Core API
func (gateway *CoreGateway) Call(method, path string, header map[string]string, body io.Reader, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = gateway.Client.BaseUrl + path

	return gateway.Client.Call(method, path, header, body, v)
}

func (gateway *CoreGateway) GetPublicToken(req *PublicTokenRequest) (res PublicTokenResponse, err error) {
	data := url.Values{}
	data.Set("trxId", req.TrxId)
	data.Set("terminalId", gateway.Client.TerminalId)
	data.Set("userKey", gateway.Client.UserKey)
	data.Set("password", gateway.Client.Password)
	data.Set("signature", gateway.Client.Signature)
	data.Set("total", req.Total)
	data.Set("successUrl", req.SuccessUrl)
	data.Set("failedUrl", req.FailedUrl)
	data.Set("items", req.Items)
	data.Set("msisdn", req.MSISDN)
	data.Set("default_language", req.DefaultLanguage)
	data.Set("default_template", req.DefaultTemplate)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", PublicTokenRequestUrl, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}
