package downloader

import (
	"crypto/tls"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/R4v3nl0/MDownloader/config"
)

func getRequest(targetUrl string, cfg *config.Config) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}

	// set timeout
	client.Timeout = time.Duration(cfg.Requests.Timeout) * time.Second

	// set proxy
	if cfg.Requests.Proxy != "" {
		proxyUrl, err := url.Parse(cfg.Requests.Proxy)
		if err != nil {
			return nil, err
		}
		client.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}

	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	for k, v := range cfg.Requests.Headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
