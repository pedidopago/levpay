package levpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
)

const baseURL string = "https://homolog.levpay.com/publicapi"

// Config define the struct to configurated Levpay
type Config struct {
	ApiKey    string
	SecretKey string
	Trace     bool
	Logger    *zap.Logger
	Client    *http.Client
}

// Do is a function to execute every api of levpay
func (c *Config) Do(method, urlpart string, body io.Reader) (*http.Response, error) {

	// generate token
	token, err := c.getLevpayAuthenticationToken()
	if err != nil {
		fmt.Println("[LEVPAY] GetLevpayAvailableAccounts e1", err.Error())
		return nil, err
	}

	// execute main api
	request, err := http.NewRequest(method, baseURL+urlpart, body)
	if err != nil {
		return nil, err
	}

	bearer := "Bearer " + token.Token
	request.Header.Set("Authorization", bearer)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(c.ApiKey, c.SecretKey)

	if c.Trace {
		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("%v %v %v\n", request.Method, request.URL, request.Proto))
		buf.WriteString(fmt.Sprintf("Host: %v\n", request.Host))
		// headers
		for name, headers := range request.Header {
			for _, v := range headers {
				buf.WriteString(fmt.Sprintf("%v: %v\n", name, v))
			}
		}

		if request.Body != nil {
			buf.WriteString("\n")
			buf2 := new(bytes.Buffer)
			io.Copy(buf2, request.Body)
			bbytes := buf2.Bytes()
			request.Body.Close()
			buf2.Reset()
			buf2.Write(bbytes)
			request.Body = ioutil.NopCloser(buf2)
			buf.Write(bbytes)
			buf.WriteRune('\n')
		}
		fmt.Println(buf.String())
	}
	//
	if c.Client == nil {
		return http.DefaultClient.Do(request)
	}

	return c.Client.Do(request)
}

// getLevpayAuthenticationToken generate authentication token to execute every levpay api
func (c *Config) getLevpayAuthenticationToken() (LevpayToken, error) {
	var token LevpayToken
	request, err := http.NewRequest("POST", baseURL+"/auth/", nil)
	if err != nil {
		return token, err
	}

	request.Header.Set("Content-type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.SetBasicAuth(c.ApiKey, c.SecretKey)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return token, err
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return token, err
	}

	err = json.Unmarshal(responseBody, &token)

	return token, nil
}
