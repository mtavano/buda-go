package buda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func (b *Buda) makeRequest(method, path string, body io.Reader, private bool) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", b.baseURL, path)
	log.Println("[MAKE REQUEST] url:", url)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "buda: Buda.makeRequest http.NewRequest error")
	}
	req.Header.Set("Content-Type", "application/json")

	if private {
		err = b.authenticate(req)
		if err != nil {
			return nil, errors.Wrap(err, "buda: authenticateRequest error")
		}
	}

	response, err := b.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "buda: httpClient.Do error")
	}

	fmt.Println("client.Do error", err)
	fmt.Println("client.Do response status", response.Status)
	readAndPreserveBody(response)

	return response, nil
}

func (b *Buda) scanBody(res *http.Response, scanner interface{}) error {
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, scanner)
}

// MarshallBody ...
func (b *Buda) MarshallBody(v interface{}) (io.Reader, error) {
	if v == nil {
		return nil, errors.New("buda: MarshallBody cannot marshal a null interface")
	}

	slice, err := json.Marshal(v)
	if err != nil {
		return nil, errors.Wrap(err, "buda: MarshallBody error")
	}

	return bytes.NewBuffer(slice), nil
}

func readAndPreserveBody(resp *http.Response) ([]byte, error) {
	// Lee el contenido del body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("response -->", string(bodyBytes))

	// Importante: cerramos el body original
	resp.Body.Close()

	// Restauramos el body para que pueda volver a ser leído
	resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return bodyBytes, nil
}
