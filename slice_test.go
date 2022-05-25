package lox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapErrNoErr(t *testing.T) {
	is := assert.New(t)

	result, err := MapErr([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, error) {
		return strconv.FormatInt(x, 10), nil
	})

	is.NoError(err)
	is.Equal(len(result), 4)
	is.Equal(result, []string{"1", "2", "3", "4"})
}

func TestMapErrWithErr(t *testing.T) {
	is := assert.New(t)

	result, err := MapErr([]int64{1, 2, 3, 4}, func(x int64, _ int) (string, error) {
		return "", fmt.Errorf("error")
	})

	is.ErrorContains(err, "error")
	is.Nil(result)
}
