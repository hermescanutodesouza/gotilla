package imageutil

import (
	"fmt"
	"image"
	"image/color"
	"strings"

	"golang.org/x/image/draw"
)

// AddBackgroundWhite adds a white background which is usable
// when the image has a transparent background.
func AddBackgroundWhite(imgSrc image.Image) image.Image {
	imgNew := image.NewRGBA(imgSrc.Bounds())
	draw.Draw(imgNew, imgNew.Bounds(),
		&image.Uniform{color.White}, image.Point{}, draw.Src)
	draw.Draw(imgNew, imgNew.Bounds(),
		imgSrc, imgSrc.Bounds().Min, draw.Over)
	return imgNew
}

// Resize scales an image to the provided size units. Use a 0
// to scale the aspect ratio. See gitub.com/nfnt/resize for Lanczos3, etc.
// https://github.com/nfnt/resize .
func Resize(width, height uint, src image.Image, scale draw.Scaler) image.Image {
	if width == 0 && height != 0 {
		width = uint(ImageAspect(src) * float64(height))
	} else if height == 0 && width != 0 {
		height = uint(float64(width) / ImageAspect(src))
	}
	rect := image.Rect(0, 0, int(width), int(height))
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

func ResizeMaxDimension(maxSide uint, src image.Image, scale draw.Scaler) image.Image {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	if width > height {
		if width == int(maxSide) {
			return src
		}
		return Resize(maxSide, 0, src, scale)
	}
	if height == int(maxSide) {
		return src
	}
	return Resize(0, maxSide, src, scale)
}

func Square(src image.Image) image.Image {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	if width == height {
		return src
	}
	if height > width {
		// https://www.golangprograms.com/how-to-add-watermark-or-merge-two-image.html
		newSq := image.NewRGBA(image.Rect(0, 0, height, height))
		offset := image.Point{
			X: int((float64(height) - float64(width)) / 2),
			Y: 0}
		draw.Draw(newSq, src.Bounds().Add(offset), src, image.Point{}, draw.Over)
		return newSq
	}
	newSq := image.NewRGBA(image.Rect(0, 0, width, width))
	offset := image.Point{
		X: 0,
		Y: int((float64(width) - float64(height)) / 2)}
	draw.Draw(newSq, src.Bounds().Add(offset), src, image.Point{}, draw.Over)
	return newSq
}

// Scale will resize the image to the provided rectangle using the
// provided interpolation function.
func Scale(src image.Image, rect image.Rectangle, scale draw.Scaler) image.Image {
	dst := image.NewRGBA(rect)
	scale.Scale(dst, rect, src, src.Bounds(), draw.Over, nil)
	return dst
}

// ScalerDefault returns a general best results interpolation
// algorithm. See more here https://blog.codinghorror.com/better-image-resizing/ ,
// https://support.esri.com/en/technical-article/000005606 ,
// https://stackoverflow.com/questions/384991/what-is-the-best-image-downscaling-algorithm-quality-wise/6171860 .
func ScalerDefault() draw.Scaler { return draw.BiLinear }

func ScalerBest() draw.Scaler { return draw.CatmullRom }

func ParseScaler(rawInterpolation string) (draw.Scaler, error) {
	rawInterpolation = strings.ToLower(strings.TrimSpace(rawInterpolation))
	switch rawInterpolation {
	case "nearestneighbor":
		return draw.NearestNeighbor, nil
	case "approxbilinear":
		return draw.ApproxBiLinear, nil
	case "bilinear":
		return draw.BiLinear, nil
	case "catmullrom":
		return draw.CatmullRom, nil
	}
	return draw.NearestNeighbor, fmt.Errorf("Cannot parse Scalar [%s]", rawInterpolation)
}

func PaintColorRGBA(img *image.RGBA, clr color.RGBA) {
	if img == nil {
		return
	}
	PaintColorRGBARectangle(img, clr, img.Bounds())
}

func PaintColorRGBARectangle(img *image.RGBA, clr color.RGBA, rectNew image.Rectangle) {
	if img == nil {
		return
	}
	rectImg := img.Bounds()

	for x := rectNew.Min.X; x < rectNew.Max.X; x++ {
		if x < rectImg.Min.X || x > rectImg.Max.X {
			continue
		}
		for y := rectNew.Min.Y; y < rectNew.Max.Y; y++ {
			if y < rectImg.Min.Y || y > rectImg.Max.Y {
				continue
			}
			img.Set(x, y, clr)
		}
	}
}
