package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-15 12:27:31.195664696 +0000 UTC m=+3.187588379"),
		StatusCode: http.StatusOK,
	}, nil
}
