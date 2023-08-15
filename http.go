package dingdingrobot

import (
	"bytes"
	"context"
	"net/http"
)

var client *http.Client

func init() {
	client = http.DefaultClient
}

func SendRequest(ctx context.Context, method string, uri string, header map[string]string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// set Header
	if len(header) > 0 {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	// with context
	req = req.WithContext(ctx)

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
