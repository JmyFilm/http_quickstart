package utils

import (
	"io"
	"net/http"
	"time"
)

type SMap map[string]string

func OkHTTP(method string, url string, header SMap, body io.Reader) (respBody []byte, statusCode int, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return respBody, statusCode, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := (&http.Client{Timeout: time.Second * 27}).Do(req)
	if resp == nil || err != nil {
		return respBody, statusCode, err
	}

	statusCode = resp.StatusCode
	if resp.Body != nil {
		defer func() {
			_ = resp.Body.Close()
		}()
		respBody, err = io.ReadAll(resp.Body)
	}

	return respBody, statusCode, err
}
