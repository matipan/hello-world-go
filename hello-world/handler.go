package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 20:24:59.110818072 +0000 UTC m=+13.153665448"),
		StatusCode: http.StatusOK,
	}, nil
}
