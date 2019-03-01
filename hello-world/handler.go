package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-01 16:40:06.093342537 +0000 UTC m=+325812.889547450"),
		StatusCode: http.StatusOK,
	}, nil
}
