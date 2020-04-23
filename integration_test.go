// +build integration

package facest

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	faceID = "integration_face_id"
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
			assert.Equal(t, 390, res.Faces[0].Rectangle.Top, "expecting top 390")
			assert.Equal(t, 689, res.Faces[0].Rectangle.Left, "expecting left 689")
			assert.Equal(t, 107, res.Faces[0].Rectangle.Width, "expecting width 107")
			assert.Equal(t, 108, res.Faces[0].Rectangle.Height, "expecting height 108")
		}
	}
}

func TestTrain(t *testing.T) {
	c := NewClient(os.Getenv("FACEST_INTEGRATION_API_KEY"))

	nofaces, err := os.Open("./testdata/nofaces.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Train(nofaces, faceID)
		assert.NotNil(t, err, "expecting non-nil error")
		assert.Nil(t, res, "expecting nil result")

		if err != nil {
			assert.Equal(t, "no faces detected", err.Error(), "expecting no faces detected error")
		}
	}

	face1, err := os.Open("./testdata/face2.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Train(face1, faceID)
		assert.Nil(t, err, "expecting nil error")
		assert.NotNil(t, res, "expecting non-nil result")

		if res != nil {
			assert.NotEmpty(t, res.FaceToken, "expecting non-empty face_token")
			assert.NotEmpty(t, res.ImageToken, "expecting non-empty image_token")
			assert.NotEmpty(t, res.ImageURL, "expecting non-empty image_url")
		}
	}
}

func TestRecognize(t *testing.T) {
	c := NewClient(os.Getenv("FACEST_INTEGRATION_API_KEY"))

	nofaces, err := os.Open("./testdata/nofaces.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Recognize(nofaces)
		assert.Nil(t, err, "expecting nil error")
		assert.NotNil(t, res, "expecting non-nil result")
		if res != nil {
			assert.Equal(t, 0, res.Count, "expecting 0 faces found")
		}
	}

	faces, err := os.Open("./testdata/faces.jpg")
	assert.Nil(t, err, "unable to open test file")

	if err == nil {
		res, err := c.Recognize(faces)
		assert.Nil(t, err, "expecting nil error")
		assert.NotNil(t, res, "expecting non-nil result")
		if res != nil {
			assert.Equal(t, 1, res.Count, "expecting 1 face found")
			assert.Less(t, 0.6, res.Faces[0].Confidence, "expecting confidence > 0.6")
			assert.Greater(t, 0.75, res.Faces[0].Confidence, "expecting confidence < 0.75")
			assert.Equal(t, faceID, res.Faces[0].FaceID, "expecting correct face_id")
		}
	}
}
