package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 22:10:20.66471465 +0000 UTC m=+27.460919558"),
		StatusCode: http.StatusOK,
	}, nil
}
