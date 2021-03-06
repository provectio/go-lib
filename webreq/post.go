package webreq

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// POST:
// Web request with method POST into Json.
// Not usefull to set "Content-Type" : "applciation/json" in headers
func POST(url string, headers HeadersKey, postObj interface{}) (result []byte, statusCode int, err error) {

	var jsonBytes []byte
	jsonBytes, err = json.Marshal(postObj)
	if err != nil {
		return
	}

	client := &http.Client{
		CheckRedirect: http.DefaultClient.CheckRedirect,
		Timeout:       3 * time.Second,
	}

	var body io.Reader
	if GzipSupport {
		bodyWriter := &bytes.Buffer{}
		writer := gzip.NewWriter(bodyWriter)
		writer.Write(jsonBytes)
		writer.Close()
		body = bytes.NewBuffer(bodyWriter.Bytes())
	} else {
		body = bytes.NewBuffer(jsonBytes)
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return
	}

	for keyName, keyValue := range headers {
		req.Header.Set(keyName, keyValue)
	}

	req.Header.Set("Content-Type", `application/json ; charset=utf8`)
	if GzipSupport {
		req.Header.Set("Content-Encoding", `gzip`)
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
	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return
		}
		reader.Close()
	} else {
		reader = resp.Body
	}

	result, err = io.ReadAll(reader)

	return
}
