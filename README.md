# LinkAja Sangu

## Usage blueprint

1. There is a type named `Client` (`linkaja.Client`) that should be instantiated through `NewClient` which hold any possible setting to the library.
2. There is a gateway classes which you will be using depending on whether you used. The gateway type need a Client instance.
3. Any activity (public token request) is done in the gateway level.

## Example

```go
    linkAjaClient := linkaja.NewClient()
    linkAjaClient.BaseUrl = "LINK_AJA_BASE_URL",
    linkAjaClientTerminalId = "YOUR_LINK_AJA_TERMINAL_ID",
    linkAjaClientUserKey = "YOUR_LINK_AJA_USER_KEY",
    linkAjaClientPassword = "YOUR_LINK_AJA_PASSWORD",
    linkAjaClientSignature = "YOUR_LINK_AJA_SIGNATURE",

    coreGateway := linkaja.CoreGateway{
        Client: linkAjaClient,
    }

    body := &linkaja.PublicTokenRequest{
        TrxId: "YOUR_TRANSACTION_ID",
        Total: "TOTAL",
        SuccessUrl: "YOUR_SUCCESS_URL_PAGE",
        FailedUrl: "YOUR_FAILED_URL_PAGE",
        Items: "Optional"
        MSISDN: "Optional"
        DefaultLanguage: "Default: 0"
        DefaultTemplate: "Default: 0"
    }

    res, _ := coreGateway.GetPublicToken(req)
```
