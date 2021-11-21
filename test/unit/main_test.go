package unit

import (
	"hello-world/controller"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchApi(t *testing.T) {
	data := controller.SearchApi("Avengers", "faf7e5bb&s", 2)
	assert.Greater(t, data, 0)
}
