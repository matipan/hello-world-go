package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-14 15:48:48.803773589 +0000 UTC m=+257.077724069"),
		StatusCode: http.StatusOK,
	}, nil
}
