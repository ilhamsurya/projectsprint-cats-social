package utils

import (
	"bytes"
	"io"
	"net/http"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strings"

	"github.com/pkg/errors"
)

func ValidateImageFile(file io.ReadSeeker, allowedFormat []string) error {
	var fileTypeRaw = new(bytes.Buffer)

	// 512 bytes indicate the file type
	// see: https://pkg.go.dev/net/http#DetectContentType
	_, err := io.CopyN(fileTypeRaw, file, 512)
	if err != nil {
		return errors.Wrap(err, "when reading the uploaded file")
	}

	// reset the pointer to the start of the file
	file.Seek(0, io.SeekStart)

	fileType := http.DetectContentType(fileTypeRaw.Bytes())

	if !strings.Contains(fileType, "image") {
		return errors.New(msg.ErrInvalidFormatFile)
	}

	for _, validFormat := range allowedFormat {
		if fileType == validFormat {
			return nil
		}
	}

	return errors.New(msg.ErrUnsupportedImgFormat)
}
