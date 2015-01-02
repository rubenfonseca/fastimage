package fastimage

import (
	"bytes"
	"io"
	"net/http"
)

// DetectImageType is the main function used to detect the type and size
// of a remote image represented by the url.
//
// Only check ImageType and ImageSize if error is not nil.
func DetectImageType(uri string) (ImageType, *ImageSize, error) {

	logger.Printf("Opening HTTP stream")
	resp, err := http.Get(uri)

	defer resp.Body.Close()

	if err != nil {
		return Unknown, nil, err
	}

	return DetectImageTypeFromResponse(resp)
}

// DetectImageTypeFromResponse is a secondary function used to detect the type and size
// of a remote image represented by the resp.
// This way you can create your own request and then pass it here.
// Check examples from http://golang.org/pkg/net/http/
//
// Only check ImageType and ImageSize if error is not nil.
func DetectImageTypeFromResponse(resp *http.Response) (ImageType, *ImageSize, error) {
	logger.Printf("Starting operation")

	buffer := bytes.Buffer{}
	defer logger.Printf("Ended after reading %v bytes out of %v bytes", buffer.Len(), resp.ContentLength)

	for {
		err := readToBuffer(resp.Body, &buffer)
		if buffer.Len() < 2 {
			continue
		}

		if err != nil {
			logger.Printf("Bailing out because of err %v", err)
			return Unknown, nil, err
		}

		for _, ImageTypeParser := range imageTypeParsers {
			if ImageTypeParser.Detect(buffer.Bytes()) {
				t := ImageTypeParser.Type()
				size, err := ImageTypeParser.GetSize(buffer.Bytes())

				if err == nil {
					logger.Printf("Found image type %v with size %v", t, size)
					return t, size, nil
				}
				break
			}
		}
	}
}

func readToBuffer(body io.Reader, buffer *bytes.Buffer) error {
	chunk := make([]byte, 8)
	count, err := body.Read(chunk)

	logger.Printf("Read %v bytes", count)
	buffer.Write(chunk[:count])

	return err
}
