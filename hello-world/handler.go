package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-15 13:36:52.827242759 +0000 UTC m=+6.467537919"),
		StatusCode: http.StatusOK,
	}, nil
}
