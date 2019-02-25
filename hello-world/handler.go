package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 20:34:32.879175321 +0000 UTC m=+586.922022707"),
		StatusCode: http.StatusOK,
	}, nil
}
