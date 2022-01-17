package webreq

import (
	"bytes"
	"compress/gzip"
	"errors"
	"net/http"
	"strings"
	"time"
)

// GET:
// Web request with method GET.
// Not usefull to set "Content-Type" : "applciation/json" in headers
func GET(url string, headers HeadersKey) (result []byte, statusCode int, err error) {

	client := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       3 * time.Second,
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	for keyName, keyValue := range headers {
		req.Header.Set(keyName, keyValue)
	}

	req.Header.Set("Content-Type", `application/json; charset=utf8`)
	if GzipSupport {
		req.Header.Set("Accept-Encoding", `gzip`)
	}

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	statusCode = resp.StatusCode

	if contentJSON := strings.Contains(resp.Header.Get("Content-Type"), "application/json"); contentJSON && resp.Header.Get("Content-Encoding") == "gzip" {
		var reader *gzip.Reader
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		reader.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(reader)
		result = buf.Bytes()
	} else if contentJSON {
		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		result = buf.Bytes()
	} else {
		err = errors.New("bad Content-Type return")
	}

	return
}
