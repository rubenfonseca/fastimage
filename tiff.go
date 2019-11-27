package fastimage

import (
	"bytes"
	"encoding/binary"
	"errors"
)

type imageTIFF struct{}

func (t imageTIFF) Type() ImageType {
	return TIFF
}

func (t imageTIFF) Detect(buffer []byte) bool {
	firstTwoBytes := string(buffer[:2])
	return firstTwoBytes == "II" || firstTwoBytes == "MM"
}

func (t imageTIFF) GetSize(byteArray []byte) (*ImageSize, error) {
	buffer := bytes.NewBuffer(byteArray)
	byteOrder, err := t.getByteOrder(buffer)
	if err != nil {
		return nil, err
	}

	buffer.Next(2)
	t.skipOffset(buffer, byteOrder)

	return t.getSize(buffer, byteOrder), nil
}

func (t imageTIFF) getSize(buffer *bytes.Buffer, byteOrder binary.ByteOrder) *ImageSize {
	var (
		tagCount, orientation uint16
		size                  = ImageSize{}
	)
	binary.Read(buffer, byteOrder, &tagCount)

	for i := 0; i < int(tagCount); i++ {
		var key, value uint16
		binary.Read(buffer, byteOrder, &key)
		buffer.Next(6)
		binary.Read(buffer, byteOrder, &value)

		if key == 0x0100 {
			size.Width = uint32(value)
		} else if key == 0x0101 {
			size.Height = uint32(value)
		} else if key == 0x0112 {
			orientation = value
		}

		if size.Width != 0 && size.Height != 0 && orientation != 0 {
			break
		}

		buffer.Next(2)
	}

	rotated := int(orientation) >= 5
	if rotated {
		return &ImageSize{
			Width:  size.Height,
			Height: size.Width,
		}
	}

	return &size
}

func (t imageTIFF) skipOffset(buffer *bytes.Buffer, byteOrder binary.ByteOrder) {
	var offset uint32
	binary.Read(buffer, byteOrder, &offset)

	// skip offset - 8 bytes
	buffer.Next(int(offset) - 8)
}

func (t imageTIFF) getByteOrder(buffer *bytes.Buffer) (binary.ByteOrder, error) {
	byteOrder := buffer.Next(2)
	if bytes.Equal(byteOrder, []byte("II")) {
		return binary.LittleEndian, nil
	} else if bytes.Equal(byteOrder, []byte("MM")) {
		return binary.BigEndian, nil
	}

	return nil, errors.New("cannot determine TIFF image byte order")
}

func init() {
	register(&imageTIFF{})
}
