package fastimage

import (
	"testing"
)

func TestRemoteBMPImage(t *testing.T) {
	t.Parallel()

	testCases := map[string]ImageSize{
		// Windows BMP v2 or OS/2 BMP v1
		"https://samples.libav.org/image-samples/money-24bit-os2.bmp": ImageSize{455, 341},
		// OS/2 BMP v2
		"https://samples.libav.org/image-samples/bmp-files/test4os2v2.bmp": ImageSize{300, 22},
		// Windows BMP v3
		"http://www.ac-grenoble.fr/ien.vienne1-2/spip/IMG/bmp_Image004.bmp": ImageSize{477, 358},
		// Windows BMP v4
		"https://entropymine.com/jason/bmpsuite/bmpsuite/g/pal8v4.bmp": ImageSize{127, 64},
		// Windows BMP v5
		"https://entropymine.com/jason/bmpsuite/bmpsuite/g/pal8v5.bmp": ImageSize{127, 64},
	}

	for url, expectedSize := range testCases {
		imagetype, size, err := DetectImageType(url)
		if err != nil {
			t.Error("Failed to detect image type")
		}

		if imagetype != BMP {
			t.Error("Image is not BMP")
		}

		if size.Width != expectedSize.Width {
			t.Errorf("Image width is wrong. Expected %d, got %d", expectedSize.Width, size.Width)
		}

		if size.Height != expectedSize.Height {
			t.Errorf("Image height is wrong. Expected %d, got %d", expectedSize.Height, size.Height)
		}
	}
}
