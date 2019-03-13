package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-13 19:50:26.071845635 +0000 UTC m=+902.011670722"),
		StatusCode: http.StatusOK,
	}, nil
}
