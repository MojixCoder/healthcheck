package helpers

import (
	"net/http"
)

type HttpRequestResult struct {
	Response *http.Response
	Error    error
}

func HeadRequest(URL string, ch chan HttpRequestResult) {
	res, err := http.Head(URL)
	ch <- HttpRequestResult{
		Response: res,
		Error:    err,
	}
}
