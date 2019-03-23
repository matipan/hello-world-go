package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-23 18:38:28.842566325 +0000 UTC m=+85.651999017"),
		StatusCode: http.StatusOK,
	}, nil
}
