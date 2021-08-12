package gorse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"go.uber.org/zap"
)

type Request struct {
	URL         string
	Params      map[string]interface{}
	Body        []byte
	ContentType string
	Headers     map[string]string
	Method      string
}

func (r *Request) Do() ([]byte, error) {
	r.URL = JoinURL(r.URL, r.Params)
	logger.Debug(r.Method, zap.String("url", r.URL))
	req, err := http.NewRequest(r.Method, r.URL, bytes.NewReader(r.Body))
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	if r.Headers != nil {
		for k, v := range r.Headers {
			req.Header.Set(k, v)
		}
	}
	if r.ContentType == "" && r.Method != "GET" {
		r.ContentType = "application/json"
	}
	if r.ContentType != "" {
		req.Header.Set("content-type", r.ContentType)
	}
	client := http.Client{}
	result, err := client.Do(req)
	if result != nil {
		defer result.Body.Close()
	}
	if err != nil {
		logger.Error("err:", zap.String("err", err.Error()))
		return nil, err
	}
	data, err := ioutil.ReadAll(result.Body)
	logger.Debug(string(data))
	if result.StatusCode != http.StatusOK {
		logger.Error("response", zap.Int("status", result.StatusCode))
		return nil, fmt.Errorf("status code:%d", result.StatusCode)
	}
	return data, err
}

func JoinURL(uri string, params map[string]interface{}) string {
	values := url.Values{}
	if params != nil {
		for k, v := range params {
			values.Add(k, fmt.Sprintf("%v", v))
		}
		if !strings.HasSuffix(uri, "?") {
			uri = uri + "?"
		}
	}
	return uri + values.Encode()
}
