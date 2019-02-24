package function

import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-24 18:52:20.46096902 +0000 UTC m=+6140.292368805"),
		StatusCode: http.StatusOK,
	}, nil
}
