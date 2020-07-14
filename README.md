# GO-JUNO

Go-Juno is a third party SDK for [Juno Payment Solutions](https://juno.com.br)

## Install

```bash
go get -u github.com/eduardonunesp/go-juno
```

## Request OAuth Token

```go
package main

import (
 "fmt"
  "os"

  gojuno "github.com/eduardonunesp/go-juno"
)

func main() {
  ClientID = "exemplo-client-id"
  ClientSecret = "exemplo-client-secret"

  result, err := NewOauthToken(ClientID, ClientSecret)

  fmt.Println(result.AccessToken)
}
```

## IMPORTANT

Go-Juno is a incomplete and under development package use at your own risk

## LICENSE

MIT
