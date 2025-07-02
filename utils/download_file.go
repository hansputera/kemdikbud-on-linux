package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func DownloadFile(url string, path string) error {
	out, err := os.Create(path)
	if err != nil {
		return err
	}

	defer out.Close()

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %d", response.StatusCode)
	}

	bar := progressbar.DefaultBytes(response.ContentLength, fmt.Sprintf("Downloading %s", url))
	_, err = io.Copy(io.MultiWriter(out, bar), response.Body)
	if err != nil {
		return err
	}

	return nil
}
