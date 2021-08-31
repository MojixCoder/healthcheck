package helpers

import (
	"net/http"
)

func HeadRequest(URL string) (*http.Response, error) {
	res, err := http.Head(URL)
	if err != nil {
		return nil, err
	}
	res.Body.Close()
	return res, nil
}
