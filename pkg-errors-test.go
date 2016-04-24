package main

import (
	"github.com/pkg/errors"
)

func main() {
	err := errors.New("where am I?")
	errors.Print(err)
}
