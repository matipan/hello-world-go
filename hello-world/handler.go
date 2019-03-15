package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-15 13:39:59.453440255 +0000 UTC m=+31.095107699"),
		StatusCode: http.StatusOK,
	}, nil
}
