package partial_update

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPartialUpdate(t *testing.T) {
	type S struct {
		A string
		B int
	}

	value := S{
		A: "a",
		B: 1,
	}
	updater := struct {
		A string
	}{
		A: "b",
	}

	assert.NoError(t, PartialUpdate(&value, updater))
	assert.Equal(t, S{
		A: "b",
		B: 1,
	}, value)
}
