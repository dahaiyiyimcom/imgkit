package go_image_RCS3

import (
	"image"
	"image/color"
	"image/draw"
)

// flattenToWhiteRGBA composites img onto a solid white background and returns
// the result as an opaque RGBA image. JPEG has no alpha channel: encoding a
// PNG/WebP image with transparency straight to JPEG makes the encoder read
// the raw RGB values sitting underneath the alpha channel (commonly black)
// and bake them in as the visible background instead of turning white.
func flattenToWhiteRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	flat := image.NewRGBA(bounds)
	draw.Draw(flat, bounds, image.NewUniform(color.White), image.Point{}, draw.Src)
	draw.Draw(flat, bounds, img, bounds.Min, draw.Over)
	return flat
}
