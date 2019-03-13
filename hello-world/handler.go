package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-13 19:57:13.49182997 +0000 UTC m=+1309.431655045"),
		StatusCode: http.StatusOK,
	}, nil
}
