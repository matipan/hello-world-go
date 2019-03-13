package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-03-13 19:40:11.815711165 +0000 UTC m=+287.755536231"),
		StatusCode: http.StatusOK,
	}, nil
}
