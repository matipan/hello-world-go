package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-15 13:27:50.780810062 +0000 UTC m=+8.318105455"),
		StatusCode: http.StatusOK,
	}, nil
}
