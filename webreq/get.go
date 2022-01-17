package webreq

import (
	"compress/gzip"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// GET:
// Web request with method POST into Json.
// Not usefull to set "Content-Type" : "applciation/json" in headers
func GET(url string, headers HeadersKey, resultObj *interface{}) (statusCode int, err error) {

	client := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       3 * time.Second,
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return
	}

	for keyName, keyValue := range headers {
		req.Header.Set(keyName, keyValue)
	}

	req.Header.Set("Content-Type", `application/json`)
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

	if contentJSON := resp.Header.Get("Content-Type") == "application/json"; contentJSON && resp.Header.Get("Content-Encoding") == "gzip" {
		var reader *gzip.Reader
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		reader.Close()
		err = json.NewDecoder(reader).Decode(resultObj)
	} else if contentJSON {
		err = json.NewDecoder(resp.Body).Decode(resultObj)
	} else {
		err = errors.New("bad Content-Type return")
	}

	return
}
