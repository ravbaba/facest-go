package facest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// detectRes .
type detectRes struct {
	Data DetectRes `json:"data"`
}

// DetectRes .
type DetectRes struct {
	Count int            `json:"count"`
	Faces []DetectedFace `json:"faces"`
}

// DetectedFace .
type DetectedFace struct {
	Rectangle Rectangle `json:"rectangle"`
}

// Detect faces within a given image.
func (c *Client) Detect(image io.Reader) (*DetectRes, error) {
	var buf bytes.Buffer
	w := multipart.NewWriter(&buf)
	fw, err := w.CreateFormFile("image", "image.jpg")
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(fw, image); err != nil {
		return nil, err
	}
	if err = w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.baseURL+"/detect", &buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", w.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	var fullResponse detectRes
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	// Try to unmarshall into errorResponse
	if res.StatusCode != http.StatusOK {
		var errRes *errorResponse
		if err = json.NewDecoder(res.Body).Decode(errRes); err == nil {
			return nil, errors.New(errRes.Message)
		}

		return nil, fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return nil, err
	}

	return &fullResponse.Data, nil
}
