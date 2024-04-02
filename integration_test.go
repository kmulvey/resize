package resize

import (
	"fmt"
	"image"
	"testing"
)

var images = map[string]image.Image{
	"imgAlpha":   image.NewAlpha(image.Rect(0, 0, 100, 100)),
	"imgAlpha16": image.NewAlpha16(image.Rect(0, 0, 100, 100)),
	"imgGray16":  image.NewGray16(image.Rect(0, 0, 100, 100)),
	"imgRGBA64":  image.NewRGBA64(image.Rect(0, 0, 100, 100)),
	"imgNRGBA":   image.NewNRGBA(image.Rect(0, 0, 100, 100)),
}
var interpolationFunctions = map[string]InterpolationFunction{
	"bilinear":          Bilinear,
	"bicubic":           Bicubic,
	"mitchellNetravali": MitchellNetravali,
	"lanczos2":          Lanczos2,
	"lanczos3":          Lanczos3,
}

func TestAllImageAndInterpolationCombos(t *testing.T) {
	t.Parallel()

	for imgType, img := range images {
		for funcName, function := range interpolationFunctions {
			m := Resize(50, 50, img, function)
			if m.Bounds().Dx() != 50 && m.Bounds().Dy() != 50 {
				fmt.Printf("test failed: image type: %s, InterpolationFunction: %s", imgType, funcName)
				t.Fail()
			}
		}
	}
}
