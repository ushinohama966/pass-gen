package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterJoin(t *testing.T) {
	t.Parallel()
	t.Run("shoud return empty when arg is empty", func(t *testing.T) {
		t.Parallel()
		a := assert.New(t)
		exp := []byte{}
		act := letterJoin([]byte{})

		a.Equal(exp, act)
	})
	t.Run("shoud return empty when arg is empty", func(t *testing.T) {
		t.Parallel()
		a := assert.New(t)
		exp := []byte{}
		act := letterJoin([]byte{})

		a.Equal(exp, act)
	})
	t.Run("shoud return same arg when arg is one byte array", func(t *testing.T) {
		t.Parallel()
		a := assert.New(t)
		exp := []byte{'a', 'b', 'c'}
		act := letterJoin([]byte{'a', 'b', 'c'})

		a.Equal(exp, act)
	})
	t.Run("shoud return joined arg when arg is two byte arrays", func(t *testing.T) {
		t.Parallel()
		a := assert.New(t)
		exp := []byte{'a', 'b', 'c', 'd', 'e', 'f'}
		arg1 := []byte{'a', 'b', 'c'}
		arg2 := []byte{'d', 'e', 'f'}
		act := letterJoin(arg1, arg2)

		a.Equal(exp, act)
	})
}
