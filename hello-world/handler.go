package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-25 18:18:54.424902098 +0000 UTC m=+63846.520788064"),
		StatusCode: http.StatusOK,
	}, nil
}
