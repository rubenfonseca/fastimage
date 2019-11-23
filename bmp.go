package fastimage

import (
	"bytes"
	"encoding/binary"
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

func (b imageBMP) GetSize(byteArray []byte) (*ImageSize, error) {
	if len(byteArray) < 28 {
		return nil, errors.New("Insufficient data")
	}

	buffer := bytes.NewBuffer(byteArray)
	buffer.Next(14) // skip BITMAPFILEHEADER

	var headerType uint32
	binary.Read(buffer, binary.LittleEndian, &headerType)

	var imageSize ImageSize
	if headerType == 12 {
		// handle Windows BMP v2 or OS/2 BMP v1
		imageSize.Width = uint32(readULint16(buffer.Next(2)))
		imageSize.Height = uint32(readULint16(buffer.Next(2)))
	} else {
		binary.Read(buffer, binary.LittleEndian, &imageSize.Width)
		binary.Read(buffer, binary.LittleEndian, &imageSize.Height)
	}

	return &imageSize, nil
}

func init() {
	register(&imageBMP{})
}
