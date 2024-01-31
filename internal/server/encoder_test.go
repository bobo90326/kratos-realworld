package server

import (
	"encoding/json"
	"fmt"
	"kratos-realworld/internal/errors"
	"testing"
)

func TestEncoder(t *testing.T) {
	a := &errors.HttpError{
		Errors: map[string][]string{"body": {"can't be empty"}},
	}
	res, err := json.Marshal(a)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
	fmt.Println(string(res))
}
