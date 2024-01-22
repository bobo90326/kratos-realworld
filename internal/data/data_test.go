package data

import (
	"kratos-realworld/internal/conf"
	"testing"
)

func TestNewDB(t *testing.T) {

	db := NewDB(&conf.Data{})
	t.Log(db)
}
