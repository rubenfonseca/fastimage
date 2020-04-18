package fastimage

import (
	"testing"
)

func TestPNGImage(t *testing.T) {
	t.Parallel()

	url := "http://fc08.deviantart.net/fs71/f/2012/214/7/c/futurama__bender_by_suzura-d59kq1p.png"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != PNG {
		t.Error("Image is not PNG")
	}

	if size.Width != 988 {
		t.Error("Image width is wrong")
	}

	if size.Height != 1240 {
		t.Error("Image height is wrong")
	}
}

func TestJPEGImage(t *testing.T) {
	t.Parallel()

	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != JPEG {
		t.Error("Image is not JPEG")
	}

	if size.Width != 5000 {
		t.Error("Image width is wrong")
	}

	if size.Height != 2813 {
		t.Error("Image height is wrong")
	}
}

func TestGIFImage(t *testing.T) {
	t.Parallel()

	url := "http://media.giphy.com/media/gXcIuJBbRi2Va/giphy.gif"

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if size == nil {
		t.Error("Failed to detect image size")
	}

	if imagetype != GIF {
		t.Error("Image is not GIF")
	}

	if size.Width != 500 {
		t.Error("Image width is wrong")
	}

	if size.Height != 286 {
		t.Error("Image height is wrong")
	}
}

func TestTIFFImage(t *testing.T) {
	t.Parallel()

	url := "http://www.fileformat.info/format/tiff/sample/c44cf1326c2240d38e9fca073bd7a805/download"
	expectedSize := ImageSize{1419, 1001}

	imagetype, size, err := DetectImageType(url)
	if err != nil {
		t.Error("Failed to detect image type")
	}

	if imagetype != TIFF {
		t.Error("Image is not TIFF")
	}

	if size.Width != expectedSize.Width {
		t.Errorf("Image width is wrong. Expected %d, got %d", expectedSize.Width, size.Width)
	}

	if size.Height != expectedSize.Height {
		t.Errorf("Image height is wrong. Expected %d, got %d", expectedSize.Height, size.Height)
	}
}

func TestCustomTimeout(t *testing.T) {
	t.Parallel()

	url := "http://slowwly.robertomurray.co.uk/delay/10000/url/http://loremflickr.com/500/500"

	imagetype, size, err := DetectImageTypeWithTimeout(url, 1000)
	t.Logf("imageType: %v", imagetype)
	t.Logf("size: %v", size)
	t.Logf("error: %v", err)
	if err == nil {
		t.Error("Timeout expected, but not occurred")
	}
}
