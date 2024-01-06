package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func PostForValues(_url, contentType string, body io.Reader) (url.Values, error) {
	resp, err := http.Post(_url, contentType, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// FIXME: on error, the body might not be URL-encoded
	values, err := url.ParseQuery(string(bytes))
	if err != nil {
		return nil, fmt.Errorf("failed to parse response body %s: %w", resp.Status, err)
	}
	return values, nil
}
