package utils

import (
	"context"
	"io"
	"net/http"
	"time"
)

var clientOne = &http.Client{Timeout: time.Second * 27}

type SMap map[string]string

func OkHTTP(method string, url string, header SMap, body io.Reader) (respBody []byte, statusCode int, err error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*7))
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return respBody, statusCode, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := clientOne.Do(req)
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
