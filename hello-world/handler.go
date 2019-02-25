package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 20:24:50.706016879 +0000 UTC m=+198.188093749"),
		StatusCode: http.StatusOK,
	}, nil
}
