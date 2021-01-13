package extparty

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type Breeds struct {
	Message map[string][]string
}

type ImegeBread struct {
	Message []string
}

func GetDogBreeds() (breeds *Breeds, err error) {
	path := "https://dog.ceo/api/breeds/list/all"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}

	var res []byte
	if res, err = submit(req); err != nil {
		return
	}

	if err = json.Unmarshal(res, &breeds); err != nil {
		return
	}

	return
}

func GetEmagesDogBreeds(name string) (image *ImegeBread, err error) {
	path := "https://dog.ceo/api/breed/" + name + "/images"

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return
	}

	var res []byte
	if res, err = submit(req); err != nil {
		return
	}

	if err = json.Unmarshal(res, &image); err != nil {
		return
	}

	return
}

func submit(req *http.Request) (response []byte, err error) {
	client := &http.Client{
		Timeout: time.Duration(20) * time.Second,
	}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	if requestDump, e := httputil.DumpRequest(req, true); e == nil {
		fmt.Println(string(requestDump))
	}

	var res *http.Response
	if res, err = client.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	if responseDump, e := httputil.DumpResponse(res, true); e == nil {
		fmt.Println(string(responseDump))
	}

	if response, err = ioutil.ReadAll(res.Body); err != nil {
		return
	}

	return
}
