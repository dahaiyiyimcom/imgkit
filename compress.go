package go_image_RCS3

import (
	"bytes"
	"encoding/base64"
	"errors"
	webp2 "github.com/chai2010/webp"
	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func ImgCompress(width, height uint, quality int, imgBase64 string) (string, error) {

	var err error

	index := strings.Index(imgBase64, ";base64,")
	if index < 0 {
		return "", errors.New("Invalid image")
	}
	imgExt := imgBase64[11:index]

	unbasedImage, err := base64.StdEncoding.DecodeString(imgBase64[index+8:])
	if err != nil {
		return "", err
	}

	var newImage string

	switch imgExt {
	case "png":
		img, err := png.Decode(bytes.NewReader(unbasedImage))
		if err != nil {
			panic("bad png")
		}

		resizedImage := resize.Resize(width, height, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		if err = jpeg.Encode(buf, resizedImage, &jpeg.Options{Quality: quality}); err != nil {
			return "", err
		}

		//compressedImage, err := jpeg.Decode(buf)
		compressedImageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		//newImage = "data:image/jpeg;base64," + compressedImageBase64
		newImage = compressedImageBase64

	case "jpeg", "jpg":
		img, _, err := image.Decode(strings.NewReader(string(unbasedImage)))
		if err != nil {
			return "", err
		}
		resizedImage := resize.Resize(width, height, img, resize.Lanczos3)
		buf := new(bytes.Buffer)
		if err = jpeg.Encode(buf, resizedImage, &jpeg.Options{Quality: quality}); err != nil {
			return "", err
		}

		//compressedImage, err := jpeg.Decode(buf)
		compressedImageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		//newImage = "data:image/jpeg;base64," + compressedImageBase64
		newImage = compressedImageBase64
	case "webp":
		img, err := webp.Decode(bytes.NewReader(unbasedImage))
		if err != nil {
			return "", err
		}
		resizedImage := resize.Resize(width, height, img, resize.Lanczos3)

		var buf bytes.Buffer
		if err = jpeg.Encode(&buf, resizedImage, nil); err != nil {
			return "", err
		}
		compressedImageBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		newImage = compressedImageBase64
	default:
		err = errors.New("invalid extension")
	}

	return newImage, err
}

func ImgCompressToWebP(width, height uint, maxSizeKB int, imgBase64 string) (string, error) {
	var err error
	var newImage string

	index := strings.Index(imgBase64, ";base64,")
	if index < 0 {
		return "", errors.New("Invalid image")
	}
	imgExt := imgBase64[11:index]
	unbasedImage, err := base64.StdEncoding.DecodeString(imgBase64[index+8:])
	if err != nil {
		return "", err
	}

	var img image.Image
	switch imgExt {
	case "png":
		img, err = png.Decode(bytes.NewReader(unbasedImage))
	case "jpeg", "jpg":
		img, _, err = image.Decode(bytes.NewReader(unbasedImage))
	case "webp":
		img, err = webp.Decode(bytes.NewReader(unbasedImage))
	default:
		return "", errors.New("Unsupported image format")
	}
	if err != nil {
		return "", err
	}

	// Resize
	resizedImage := resize.Resize(width, height, img, resize.Lanczos3)

	// Dynamic quality adjustment
	var buf bytes.Buffer
	for quality := float32(85); quality >= 60; quality -= 5 {
		buf.Reset()
		if err = webp2.Encode(&buf, resizedImage, &webp2.Options{Lossless: false, Quality: quality}); err != nil {
			return "", err
		}
		if buf.Len() <= maxSizeKB*1024 {
			break
		}
	}

	newImage = base64.StdEncoding.EncodeToString(buf.Bytes())

	// Son kalite hâlini döndür
	return newImage, nil
}
