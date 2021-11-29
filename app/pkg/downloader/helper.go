package downloader

import (
	"net/http"
)

func HttpClient(a http.RoundTripper) *http.Client {
	return &http.Client{Transport: a}
}
func ClientDo(req *http.Request, client *http.Client) (*http.Response, error) {
	return client.Do(req)
}
func HttpGet(url string) (*http.Response, error) {
	return http.Get(url)
}
