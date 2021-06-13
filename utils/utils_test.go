package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestGenerateAndVerifyPassword(t *testing.T) {
	pwd := "123123"
	p, salt := GeneratePassword(pwd)
	fmt.Println(p)
	fmt.Println(salt)
	f := VerifyPassword(pwd, p, salt)
	assert.Equal(t, f, true)
}

