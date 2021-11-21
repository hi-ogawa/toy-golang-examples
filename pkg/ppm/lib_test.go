package ppm

import (
	"bytes"
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/stretchr/testify/assert"
)

func TestFprint(t *testing.T) {
	data := []mgl32.Vec4{{0.5, 1, 1}, {1, 0.5, 1}, {1, 1, 0.5}, {0.25, 1, 1}, {1, 0.25, 1}, {1, 1, 0.25}}
	buffer := bytes.Buffer{}
	Fprint(&buffer, data, 2, 3)
	actual := buffer.String()
	expected := `P3
2 3
255
128 255 255
255 128 255
255 255 128
64 255 255
255 64 255
255 255 64
`
	assert.Equal(t, actual, expected)
}
