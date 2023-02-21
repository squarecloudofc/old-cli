package properties

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnmarshall(t *testing.T) {
	type T struct {
		ST string `properties:"ST"`
		St string `properties:"St"`
	}
	text := `
ST=asd
St=dsa
`

	var r T
	err := Unmarshall([]byte(text), &r)
	assert.Equal(t, err, nil)
	assert.Equal(t, r, T{
		ST: "asd",
		St: "dsa",
	})
}
