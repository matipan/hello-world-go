package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 00:20:47.135316793 +0000 UTC m=+125.505262653"),
		StatusCode: http.StatusOK,
	}, nil
}
