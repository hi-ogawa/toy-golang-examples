package ppm

import (
	"fmt"
	"io"

	m "github.com/go-gl/mathgl/mgl32"
)

const Header = `P3
%d %d
255
`

func Fprint(writer io.Writer, data []m.Vec4, width, height int) error {
	if _, err := fmt.Fprintf(writer, Header, width, height); err != nil {
		return err
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			v := data[y*width+x]
			rgb := [3]uint8{}
			for i := 0; i < 3; i++ {
				rgb[i] = uint8(m.Clamp(v[i]*256, 0, 255))
			}
			if _, err := fmt.Fprintf(writer, "%d %d %d\n", rgb[0], rgb[1], rgb[2]); err != nil {
				return err
			}
		}
	}
	return nil
}
