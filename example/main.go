package main

import (
	"fmt"
	"github.com/otterize/lox"
	"strings"
)

func main() {
	names, err := lox.MapErr([]string{"Samuel", "Marc", "Samuel"}, func(s string, i int) (string, error) {
		if s == "" {
			return "", fmt.Errorf("empty name")
		}
		return s, nil
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Names: %s", strings.Join(names, ", "))
}
