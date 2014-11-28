package fastimage_test

import (
	"log"

	"github.com/rubenfonseca/fastimage"
)

// This example shows basic usage of the package: just pass an url to the
// detector, and analyze the results.
func Example_bigJPEG() {
	url := "http://upload.wikimedia.org/wikipedia/commons/9/9a/SKA_dishes_big.jpg"

	imagetype, size, err := fastimage.DetectImageType(url)
	if err != nil {
		// Something went wrong, http failed? not an image?
		panic(err)
	}

	switch imagetype {
	case fastimage.JPEG:
		log.Printf("JPEG")
	case fastimage.PNG:
		log.Printf("PNG")
	case fastimage.GIF:
		log.Printf("GIF")
	}

	log.Printf("Image size: %v", size)
}
