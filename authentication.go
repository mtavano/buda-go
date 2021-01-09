package buda

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// {HTTP method} {path} {base64_encoded_body} {nonce}

// This is a stable version. do not touch
func (b *Buda) authenticate(req *http.Request) error {
	valArray, err := createValArray(req)
	if err != nil {
		return nil
	}

	nonce := strconv.FormatInt(time.Now().UnixNano(), 10)
	valArray = append(valArray, nonce)
	sign := createSign(valArray, b.secret)

	req.Header.Add("X-SBTC-APIKEY", b.key)
	req.Header.Add("X-SBTC-NONCE", nonce)
	req.Header.Add("X-SBTC-SIGNATURE", sign)

	return nil
}

func createValArray(req *http.Request) ([]string, error) {
	var params []string

	params = append(params, req.Method)
	params = append(params, req.URL.RequestURI())

	if req.Method == http.MethodPost || req.Method == http.MethodPut {
		b := req.Body
		body, err := ioutil.ReadAll(b)
		if err != nil {
			return nil, err
		}
		params = append(params, base64.StdEncoding.EncodeToString(body))
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
	}

	return params, nil
}

func createSign(valArray []string, secret string) string {
	h := hmac.New(sha512.New384, []byte(secret))
	rawStr := strings.Join(valArray, " ")
	h.Write([]byte(rawStr))
	return hex.EncodeToString(h.Sum(nil))
}
