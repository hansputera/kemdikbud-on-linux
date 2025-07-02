package utils

import (
	"bytes"
	"errors"
	"io"
	"os"

	"github.com/hansputera/kemdikbud-on-linux/constants"
)

func DetectFileSig(p string) (*string, error) {
	file, err := os.Open(p)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	head := data
	if len(data) > 8192 {
		head = data[:8192]
	}

	for sig, kind := range constants.FILE_SIGNATURES {
		if bytes.Contains(head, []byte(sig)) {
			return &kind, nil
		}
	}

	return nil, errors.New("unknown sig")
}
