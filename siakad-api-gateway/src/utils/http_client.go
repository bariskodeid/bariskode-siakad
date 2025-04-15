package utils

import (
    "bytes"
    "io"
    "net/http"
)

func ForwardRequest(method, targetURL string, headers map[string]string, body []byte) (*http.Response, error) {
    req, err := http.NewRequest(method, targetURL, bytes.NewReader(body))
    if err != nil {
        return nil, err
    }

    for key, value := range headers {
        req.Header.Set(key, value)
    }

    client := &http.Client{}
    return client.Do(req)
}

func ReadBody(resp *http.Response) ([]byte, error) {
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}
