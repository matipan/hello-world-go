package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-26 11:32:49.367899628 +0000 UTC m=+233746.177332322"),
		StatusCode: http.StatusOK,
	}, nil
}
