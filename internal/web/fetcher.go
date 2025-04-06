package web

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func GetBody(client *resty.Client, url string) ([]byte, error) {
	res, err := client.R().
		Get(url)
	if err != nil {
		return nil, fmt.Errorf("error while POST-request: %s", err.Error())
	}
	if res.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode(), res.Status())
	}

	return res.Body(), nil
}
