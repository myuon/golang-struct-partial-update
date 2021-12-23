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

	type testCase struct {
		run  func(value *S) error
		want S
	}

	testCases := []testCase{
		{
			run: func(value *S) error {
				return PartialUpdate(value, struct{ A string }{
					A: "b",
				})
			},
			want: S{
				A: "b",
				B: 1,
			},
		},
		{
			// set nil to do nothing
			run: func(value *S) error {
				return PartialUpdate(value, struct{ A *string }{
					A: nil,
				})
			},
			want: S{
				A: "a",
				B: 1,
			},
		},
		{
			// pass a pointer to set the value
			run: func(value *S) error {
				return PartialUpdate(value, struct{ A *string }{
					A: nil,
				})
			},
			want: S{
				A: "a",
				B: 1,
			},
		},
	}

	for _, tt := range testCases {
		value := S{
			A: "a",
			B: 1,
		}

		assert.NoError(t, tt.run(&value))
		assert.Equal(t, tt.want, value)
	}
}
