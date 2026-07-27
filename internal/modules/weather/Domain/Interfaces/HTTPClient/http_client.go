package httpclient

import "io"

type HTTPClient interface {
	Get(url string, params map[string]string) io.ReadCloser
}
