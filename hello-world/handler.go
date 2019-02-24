package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-24 23:54:32.126965593 +0000 UTC m=+537.176423852"),
		StatusCode: http.StatusOK,
	}, nil
}
