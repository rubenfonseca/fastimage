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
	buffer := bytes.Buffer{}

	logger.Printf("Opening HTTP stream")
	resp, err := http.Get(uri)
	defer closeHTTPStream(resp, &buffer)

	if err != nil {
		return Unknown, nil, err
	}

	logger.Printf("Starting operation")

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
				size, _ := ImageTypeParser.GetSize(buffer.Bytes())

				if size != nil {
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

func closeHTTPStream(http *http.Response, buffer *bytes.Buffer) {
	logger.Printf("Closing HTTP Stream")
	http.Body.Close()
	logger.Printf("Closed after reading just %v bytes out of %v bytes", buffer.Len(), http.ContentLength)
}
