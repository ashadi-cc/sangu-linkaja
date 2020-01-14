package linkaja

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

type SanguLinkajaTestSuite struct {
	suite.Suite
}

func TestSanguLinkaja(t *testing.T) {
	suite.Run(t, new(SanguLinkajaTestSuite))
}

func (s *SanguLinkajaTestSuite) TestGenerateApplinkPaymentSuccess() {
	client := NewClient()
	client.BaseUrl = "some-string"
	client.MerchantId = "some-string"
	client.TerminalId = "some-string"
	client.UserKey = "some-string"
	client.Password = "some-string"
	client.Key = "some-string"
	client.LogLevel = 3

	gateway := CoreGateway{
		Client: client,
	}

	req := &GenerateApplinkPaymentRequest{
		TrxDate:        "some-string",
		PartnerTrxID:   "some-string",
		MerchantID:     client.MerchantId,
		TerminalID:     "some-string",
		TotalAmount:    "some-string",
		PartnerApplink: "some-string",
		Items: []Item{
			Item{
				Name:      "some-string",
				UnitPrice: "some-string",
				Quantity:  "some-string",
			},
		},
	}

	resp, err := gateway.GenerateApplinkPayment(req)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "00", resp.Status)
}

func (s *SanguLinkajaTestSuite) TestDecryptInformPaymentSuccess() {
	client := NewClient()
	client.BaseUrl = "some-string"
	client.MerchantId = "some-string"
	client.TerminalId = "some-string"
	client.UserKey = "some-string"
	client.Password = "some-string"
	client.Key = "some-string"
	client.LogLevel = 3

	gateway := CoreGateway{
		Client: client,
	}

	encrypted := "some-string"
	auth := "some-string"
	timestamp := "some-string"

	resp, err := gateway.DecryptInformPaymentRequest(encrypted, auth, timestamp)

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), resp)
	assert.Equal(s.T(), "00", resp.Status)
}
