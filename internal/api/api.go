package api

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"squarecloud/internal/config"
)

var UserAgent = "SquareGo (Square Cloud CLI)"

type ApiResponse[T any] struct {
	RawResponse *http.Response

	Status   string `json:"status"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Response T      `json:"response"`
}

func request[T any](method, path string, body io.Reader) (*ApiResponse[T], error) {
	config, _ := config.GetConfig()
	req, err := http.NewRequest(method, APIURL+path, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", config.Token)
	req.Header.Set("User-Agent", UserAgent)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	decoded, err := decodeResponse[T](response)
	if err != nil {
		return nil, err
	}

	return decoded, nil
}

func sendFile(path string, filep string) (response *http.Response, err error) {
	config, _ := config.GetConfig()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(filep)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))

	if _, err = io.Copy(part, file); err != nil {
		return
	}

	if err = writer.Close(); err != nil {
		return
	}

	req, err := http.NewRequest(http.MethodPost, APIURL+path, body)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", config.Token)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("User-Agent", UserAgent)

	client := &http.Client{}

	if response, err = client.Do(req); err != nil {
		return
	}

	return
}

func decodeResponse[T any](response *http.Response) (apiResponse *ApiResponse[T], err error) {
	decoder := json.NewDecoder(response.Body)
	if err = decoder.Decode(&apiResponse); err != nil {
		return
	}

	apiResponse.RawResponse = response
	return
}
