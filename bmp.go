package fastimage

import (
	"errors"
)

type imageBMP struct{}

func (b imageBMP) Type() ImageType {
	return BMP
}

func (b imageBMP) Detect(buffer []byte) bool {
	firstTwoBytes := buffer[:2]
	return string(firstTwoBytes) == "BM"
}

func (b imageBMP) GetSize(buffer []byte) (*ImageSize, error) {
	if len(buffer) < 28 {
		return nil, errors.New("Insufficient data")
	}

	imageSize := ImageSize{}
	imageSize.Width = readUInt32(buffer[18:22])
	imageSize.Height = readUInt32(buffer[22:26])

	return &imageSize, nil
}

func init() {
	register(&imageBMP{})
}
