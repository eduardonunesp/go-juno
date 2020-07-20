# GO-JUNO

Go-Juno is a third party SDK for [Juno Solução de Pagamento Online](https://juno.com.br)

## Install

```bash
go get -u github.com/eduardonunesp/go-juno
```

## Request OAuth Token

```go
package main

import (
  "fmt"
  "log"
  "os"

  gojuno "github.com/eduardonunesp/go-juno"
)

func CreateToken(ClientID, ClientSecret string) {
  tokenResult, err := gojuno.NewOauthToken(ClientID, ClientSecret)

  if err != nil {
    log.Fatalf("Error on get oauth token %+v\n", err)
  }

  fmt.Println(tokenResult.AccessToken)
}
```

## Create New Charge

```go
func Charge(AccessToken, ResourceToken string) {
  createChargeResult, err := gojuno.CreateCharge(gojuno.ChargeParams{
    Charge: gojuno.Charge{
      Description: "OK",
      Amount:      20.0,
      PaymentType: []string{gojuno.PaymentTypeCreditCard},
    },
    ChargeBilling: gojuno.ChargeBilling{
      Name:     "Foo Bar",
      Document: "96616796060",
    },
  }, AccessToken, ResourceToken)

  if err != nil {
    log.Fatalf("Failed to create charge cause %+v\n", err)
  }
}
```

## Create New Payment

```go
func Payment(ChargeID, CreditCardHash, AccessToken, ResourceToken string) {
  response, err := CreatePayment(PaymentParams{
    ChargeID: ChargeID,
    PaymentBilling: PaymentBilling{
      Email: "foobar@example.com",
      Address: Address{
        Street:   "Acacia Avenue",
        Number:   "22",
        City:     "FooCity",
        State:    "SC",
        PostCode: "08226021",
      },
    },
    CreditCardDetails: CreditCardDetails{
      CreditCardHash: CreditCardHash,
    },
  }, AccessToken, ResourceToken)

  if err != nil {
    log.Fatalf("Failed to create payment cause %+v\n", err)
  }
}
```

## IMPORTANT

Go-Juno is a incomplete and under development package use at your own risk

## LICENSE

MIT
