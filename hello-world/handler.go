package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 19:23:31.726759928 +0000 UTC m=+16.354836397"),
		StatusCode: http.StatusOK,
	}, nil
}
