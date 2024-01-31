package errors

import (
	"fmt"
	nethttp "net/http"

	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPError(code int, field, detail string) *HttpError {
	return &HttpError{
		Errors: map[string][]string{field: {detail}},
		Code:   code,
	}
}

type HttpError struct {
	Errors map[string][]string `json:"errors"`
	Code   int                 `json:"-"`
}

func FromError(err error) *HttpError {
	if err == nil {
		return nil
	}
	return &HttpError{
		Errors: map[string][]string{"body": {err.Error()}},
		Code:   200,
	}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTPError: %d: %v", e.Code, e.Errors)
}

func errorEncoder(w nethttp.ResponseWriter, r *nethttp.Request, err error) {
	se := FromError(err)
	codec, _ := http.CodecForRequest(r, "Accept")
	body, err := codec.Marshal(se)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/"+codec.Name())
	w.WriteHeader(se.Code)
	w.Write(body)
}
