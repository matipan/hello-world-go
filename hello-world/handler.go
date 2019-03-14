package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-14 02:37:35.485121228 +0000 UTC m=+25331.424946297"),
		StatusCode: http.StatusOK,
	}, nil
}
