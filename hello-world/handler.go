package function
import (
	"net/http"

	handler "github.com/openfaas-incubator/go-function-sdk"
)

// Handle is a function that always returns:  Hello world!
func Handle(req handler.Request) (handler.Response, error) {
	return handler.Response{
		Body:       []byte("2019-02-24 23:40:26.13478395 +0000 UTC m=+8.331328566"),
		StatusCode: http.StatusOK,
	}, nil
}
