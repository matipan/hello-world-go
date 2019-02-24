
package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-24 15:15:04.291917119 +0000 UTC m=+7.641290373"),
		StatusCode: http.StatusOK,
	}, nil
}