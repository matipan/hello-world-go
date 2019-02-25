package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 00:30:41.288911695 +0000 UTC m=+304.783286093"),
		StatusCode: http.StatusOK,
	}, nil
}
