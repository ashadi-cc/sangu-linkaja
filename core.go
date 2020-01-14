package linkaja

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
	"time"
)

const (
	PublicTokenRequestURL     = "linkaja-api/api/payment"
	CheckTransactionStatusURL = "linkaja-api/api/check/customer/transaction"
	RefundTransactionURL      = "tcash-api/api/rev/customer/transaction"
	GenerateApplinkPaymentURL = "applink/v1/create"

	MaxTimestampLength = 16
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
	data.Set("items", GenerateItems(req.Items))
	data.Set("msisdn", req.MSISDN)
	data.Set("default_language", req.DefaultLanguage)
	data.Set("default_template", req.DefaultTemplate)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", PublicTokenRequestURL, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) CheckTransactionStatus(req *TransactionRequest) (res TransactionResponses, err error) {
	data := url.Values{}
	data.Set("refNum", req.RefNum)
	data.Set("terminalId", gateway.Client.TerminalId)
	data.Set("userKey", gateway.Client.UserKey)
	data.Set("passKey", gateway.Client.Password)
	data.Set("signKey", gateway.Client.Signature)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", CheckTransactionStatusURL, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) RefundTransaction(req *TransactionRequest) (res TransactionResponses, err error) {
	data := url.Values{}
	data.Set("refNum", req.RefNum)
	data.Set("terminalId", gateway.Client.TerminalId)
	data.Set("userKey", gateway.Client.UserKey)
	data.Set("passKey", gateway.Client.Password)
	data.Set("signKey", gateway.Client.Signature)

	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	err = gateway.Call("POST", RefundTransactionURL, headers, strings.NewReader(data.Encode()), &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) GenerateApplinkPayment(req *GenerateApplinkPaymentRequest) (res GenerateApplinkPaymentResponse, err error) {
	timestamp := PadRight(fmt.Sprintf("%d", time.Now().Unix()), "0", MaxTimestampLength)
	auth := base64.StdEncoding.EncodeToString([]byte(gateway.Client.UserKey + ":" + gateway.Client.Password))

	headers := map[string]string{
		"Content-Type":  "text/plain",
		"Authorization": "Basic " + auth,
		"Timestamp":     timestamp,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return
	}

	plaintext := PKCS5Padding(body, aes.BlockSize)
	key := []byte(gateway.Client.Key)
	iv := []byte(timestamp)

	encrypted, err := gateway.encrypt(plaintext, key, iv)
	if err != nil {
		return
	}
	request := base64.StdEncoding.EncodeToString(encrypted)

	err = gateway.Call("POST", GenerateApplinkPaymentURL, headers, strings.NewReader(request), &res)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) DecryptInformPaymentRequest(encrypted string, auth string, timestamp string) (req InformPaymentRequest, err error) {
	userpass, err := base64.StdEncoding.DecodeString(auth)
	if err != nil {

		return
	}

	authtoken := strings.Split(string(userpass), ":")
	if len(authtoken) != 2 {
		err = fmt.Errorf("invalid auth")
		return
	}

	if gateway.Client.UserKey != authtoken[0] || gateway.Client.Password != authtoken[1] {
		err = fmt.Errorf("invalid auth")
		return
	}

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return
	}
	key := []byte(gateway.Client.Key)
	iv := []byte(PadRight(timestamp, "0", MaxTimestampLength))

	plaintext, err := gateway.decrypt(ciphertext, key, iv)
	if err != nil {
		return
	}

	err = json.Unmarshal(plaintext, &req)
	if err != nil {
		return
	}

	return
}

func (gateway *CoreGateway) encrypt(plaintext []byte, key []byte, iv []byte) (ciphertext []byte, err error) {
	if len(plaintext)%aes.BlockSize != 0 {
		err = fmt.Errorf("plaintext is not a multiple of the block size")
		return
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	ciphertext = make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)

	return
}

func (gateway *CoreGateway) decrypt(ciphertext []byte, key []byte, iv []byte) (plaintext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	if len(ciphertext) < aes.BlockSize {
		err = fmt.Errorf("ciphertext too short")
		return
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size")
		return
	}

	plaintext = make([]byte, len(ciphertext))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext)

	return PKCS5Trimming(plaintext), nil
}
