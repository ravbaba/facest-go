// +build integration

package facest

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	c := NewClient(os.Getenv("FACEST_INTEGRATION_API_KEY"))

	nofaces, err := os.Open("./testdata/nofaces.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Detect(nofaces)
		assert.Nil(t, err, "expecting nil error")
		assert.NotNil(t, res, "expecting non-nil result")
		if res != nil {
			assert.Equal(t, 0, res.Count, "expecting 0 faces found")
		}
	}

	faces, err := os.Open("./testdata/faces.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Detect(faces)
		assert.Nil(t, err, "expecting nil error")
		assert.NotNil(t, res, "expecting non-nil result")
		if res != nil {
			assert.Equal(t, 3, res.Count, "expecting 3 faces found")
		}
	}
}
