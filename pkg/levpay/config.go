package levpay

import (
	"fmt"
	"io"
	"net/http"

	"go.uber.org/zap"
)

// Config define the struct to configurated Levpay
type Config struct {
	ApiKey    string
	SecretKey string
	Trace     bool
	Logger    *zap.Logger
	Client    *http.Client
}

func (c *Config) Do(method, urlpart string, body io.Reader) (*http.Response, error) {
	fmt.Println("O que ta rolando")
	request, err := http.NewRequest(method, "https://homolog.levpay.com/publicapi"+urlpart, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(c.ApiKey, c.SecretKey)

	return c.Client.Do(request)
}
