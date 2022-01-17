package webreq

import (
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
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
	if statusCode > 305 || statusCode == 204 {
		return
	}

	var reader io.ReadCloser
	if contentJSON := strings.Contains(resp.Header.Get("Content-Type"), "application/json"); contentJSON && resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		reader.Close()
	} else if contentJSON {
		reader = resp.Body
	} else {
		err = errors.New("bad Content-Type return")
		if err != nil {
			return
		}
	}

	result, err = ioutil.ReadAll(reader)

	return
}
