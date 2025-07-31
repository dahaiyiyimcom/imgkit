package go_image_RCS3

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

func ImgResize(width, height uint, quality int, imgBase64 string) (string, error) {
	index := strings.Index(imgBase64, ";base64,")
	if index < 0 {
		return "", errors.New("invalid image format")
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
		return "", errors.New("unsupported image format")
	}
	if err != nil {
		return "", err
	}

	resizedImage := resize.Resize(width, height, img, resize.Lanczos3)

	var buf bytes.Buffer
	if err = jpeg.Encode(&buf, resizedImage, &jpeg.Options{Quality: quality}); err != nil {
		return "", err
	}

	// SAF Base64 (prefix yok)
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}
