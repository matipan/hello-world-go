package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-26 20:49:38.101685546 +0000 UTC m=+267154.911118272"),
		StatusCode: http.StatusOK,
	}, nil
}
