package func_call

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type httpCallReq struct {
	Url     string            `json:"url" jsonschema_description:"url of http call"`
	Method  string            `json:"method" jsonschema_description:"http method"`
	Headers map[string]string `json:"headers" jsonschema_description:"http headers"`
	Data    string            `json:"data" jsonschema_description:"http req data"`
}

var client = &http.Client{}

func httpCall(ctx context.Context, httpReq httpCallReq) (interface{}, error) {
	httpRequest, err := http.NewRequest(httpReq.Method, httpReq.Url, bytes.NewBuffer([]byte(httpReq.Data)))
	if err != nil {
		return nil, fmt.Errorf("failed to create http request: %w", err)
	}

	httpRequest.WithContext(ctx)

	// 添加请求头
	for key, value := range httpReq.Headers {
		httpRequest.Header.Add(key, value)
	}

	// 发送HTTP请求
	resp, err := client.Do(httpRequest)
	if err != nil {
		return nil, fmt.Errorf("failed to send http request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应主体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return json.RawMessage(body), nil
}
