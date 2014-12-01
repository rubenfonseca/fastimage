package fastimage

// ImageType represents the type of the image detected, or `Unknown`.
type ImageType uint

const (
	// GIF represents a GIF image
	GIF ImageType = iota
	// PNG represents a PNG image
	PNG
	// JPEG represents a JPEG image
	JPEG
	// BMP represents a BMP image
	BMP
	// Unknown represents an unknown image type
	Unknown
)

// ImageSize holds the width and height of an image
type ImageSize struct {
	Width  uint32
	Height uint32
}
