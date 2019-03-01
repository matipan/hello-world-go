package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-01 16:51:57.431702808 +0000 UTC m=+459.437149635"),
		StatusCode: http.StatusOK,
	}, nil
}
