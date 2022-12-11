package common

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type HttpAuth struct {
	BaseUrl  string
	Username string
	Password string
}

func (auth HttpAuth) Get(url string) (resp *http.Response) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}

	authKey := base64.StdEncoding.EncodeToString([]byte(auth.Username + ":" + auth.Password))
	req.Header.Set("Authorization", "Basic "+authKey)
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	return resp
}

func (auth HttpAuth) GetBytes(url string) (byt []byte) {
	resp := auth.Get(url)

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	return byt
}

func (auth HttpAuth) GetMarshal(url string, data interface{}) {
	resp := auth.Get(url)
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Panic(err)
	}
}

func (auth HttpAuth) GetWithBase(suffix string) (resp *http.Response) {
	return auth.Get(auth.BaseUrl + suffix)
}
