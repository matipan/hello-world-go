package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 19:30:20.117681852 +0000 UTC m=+35.930044724"),
		StatusCode: http.StatusOK,
	}, nil
}
