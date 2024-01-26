package biz

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	str := hashPassword("abc")
	fmt.Println(str)
}

func TestVerifyPassword(t *testing.T) {
	a := assert.New(t)
	a.True(verifyPassword("$2a$10$OlUbHuL80zral8rziK4U0uJv9Dp.KlGXyFamtvcaFupbE.u4F7sny", "abc"))
}
