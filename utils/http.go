package utils

import (
	"context"
	"crypto/tls"
	"io"
	"net/http"
	"shared-modules/common"
	"strings"
	"time"
)

func HttpDo(ctx context.Context, method string, url string, body string, header http.Header) ([]byte, http.Header, int, error) {
	switch method {
	case "GET":
		return HttpGet(ctx, url, header)
	case "POST":
		return HttpPost(ctx, url, body, header)
	case "PUT":
		return HttpPut(ctx, url, body, header)
	case "DELETE":
		return HttpDelete(ctx, url, body, header)
	}
	return nil, nil, 0, NewBusinessError(ctx, common.CODE_ILLEGAL_OPERATION)
}

func HttpGet(ctx context.Context, url string, header http.Header) ([]byte, http.Header, int, error) {

	cli := http.Client{
		Timeout: time.Second * Config.System.HTTPTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, 0, err
	}

	req = req.WithContext(ctx)
	req.Header = header

	resp, err := cli.Do(req)

	if err != nil {
		return nil, nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, 0, err
	}

	return data, resp.Header, resp.StatusCode, nil
}

func HttpPost(ctx context.Context, url string, body string, header http.Header) ([]byte, http.Header, int, error) {
	cli := http.Client{
		Timeout: time.Second * Config.System.HTTPTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
	if err != nil {
		return nil, nil, 0, err
	}

	req = req.WithContext(ctx)
	req.Header = header

	resp, err := cli.Do(req)

	if err != nil {
		return nil, nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, 0, err
	}

	return data, resp.Header, resp.StatusCode, nil
}

func HttpPut(ctx context.Context, url string, body string, header http.Header) ([]byte, http.Header, int, error) {
	cli := http.Client{
		Timeout: time.Second * Config.System.HTTPTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	req, err := http.NewRequest(http.MethodPut, url, strings.NewReader(body))

	if err != nil {
		return nil, nil, 0, err
	}

	req = req.WithContext(ctx)
	req.Header = header

	resp, err := cli.Do(req)

	if err != nil {
		return nil, nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, 0, err
	}

	return data, resp.Header, resp.StatusCode, nil
}

func HttpDelete(ctx context.Context, url string, body string, header http.Header) ([]byte, http.Header, int, error) {
	cli := http.Client{
		Timeout: time.Second * Config.System.HTTPTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // ignore expired SSL certificates
		},
	}

	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(body))

	if err != nil {
		return nil, nil, 0, err
	}

	req = req.WithContext(ctx)
	req.Header = header

	resp, err := cli.Do(req)

	if err != nil {
		return nil, nil, 0, err
	}

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, 0, err
	}

	return data, resp.Header, resp.StatusCode, nil
}
