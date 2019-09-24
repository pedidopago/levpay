package levpay

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	req, err := http.NewRequest(method, "https://homolog.levpay.com/publicapi"+urlpart, body)
	if err != nil {
		return nil, err
	}
	fmt.Println("Passou - ", c)

	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.SetBasicAuth(c.ApiKey, c.SecretKey)

	if c.Trace {
		buf := new(bytes.Buffer)
		buf.WriteString(fmt.Sprintf("%v %v %v\n", req.Method, req.URL, req.Proto))
		buf.WriteString(fmt.Sprintf("Host: %v\n", req.Host))
		// headers
		for name, headers := range req.Header {
			for _, v := range headers {
				buf.WriteString(fmt.Sprintf("%v: %v\n", name, v))
			}
		}

		if req.Body != nil {
			buf.WriteString("\n")
			buf2 := new(bytes.Buffer)
			io.Copy(buf2, req.Body)
			bbytes := buf2.Bytes()
			req.Body.Close()
			buf2.Reset()
			buf2.Write(bbytes)
			req.Body = ioutil.NopCloser(buf2)
			buf.Write(bbytes)
			buf.WriteRune('\n')
		}
		fmt.Println(buf.String())
	}
	//
	if c.Client == nil {
		return http.DefaultClient.Do(req)
	}

	fmt.Println("passou 2")
	return c.Client.Do(req)
}
