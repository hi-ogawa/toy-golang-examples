package main

import (
	"os"

	m "github.com/go-gl/mathgl/mgl32"
	"github.com/hi-ogawa/golang-examples/pkg/hash"
	"github.com/hi-ogawa/golang-examples/pkg/ppm"
)

type Config struct {
	Width        int
	Height       int
	SceneWidth   float32 // SceneHeight from aspect ratio
	SceneCenterX float32
	SceneCenterY float32
}

var DefaultConfig = Config{Width: 256, Height: 256, SceneWidth: 4, SceneCenterX: 0, SceneCenterY: 0}

func RenderPixel(coord m.Vec2) m.Vec4 {
	color := hash.Hash23(coord)
	return m.Vec4{color[0], color[1], color[2], 1}
}

func Render(config Config) []m.Vec4 {
	res := make([]m.Vec4, config.Width*config.Height)
	for y := 0; y < config.Height; y++ {
		for x := 0; x < config.Width; x++ {
			res[y*config.Width+x] = RenderPixel(m.Vec2{float32(x), float32(config.Height - y)})
		}
	}
	return res
}

func main() {
	config := DefaultConfig
	res := Render(config)
	ppm.Fprint(os.Stdout, res, config.Width, config.Height)
}
