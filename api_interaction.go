package buda

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (b *Buda) makeRequest(method, path string, body io.Reader, private bool) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", b.baseURL, path)
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
