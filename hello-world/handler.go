package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-24 23:45:43.282988691 +0000 UTC m=+8.332446951"),
		StatusCode: http.StatusOK,
	}, nil
}
