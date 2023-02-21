package properties

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Marshal(t *testing.T) {
	t.Run("Shoud return a struct with the given values", func(t *testing.T) {
		type T struct {
			ST string `properties:"ST"`
			St string `properties:"St"`
		}
		value := "ST=asd\nSt=dsa"

		r, err := Marshal(T{
			ST: "asd",
			St: "dsa",
		})
		assert.Equal(t, err, nil)
		assert.Equal(t, r, []byte(value))
	})
}
