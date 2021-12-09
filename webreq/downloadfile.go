package webreq

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if status := strconv.Itoa(resp.StatusCode); !strings.HasPrefix(status, "2") && !strings.HasPrefix(status, "3") {
		return errors.New(url + " return bad status code: " + status)
	}

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
