package desk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	apiURLPrefix       = "https://"
	apiURLSuffix       = ".desk.com/api/v2"
	clientVersion      = "1.0.0"
	defaultHTTPTimeout = 80 * time.Second
	debug              = false
)

var Username string
var Password string
var Subdomain string

var backend *Backend
var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func GetBackend() *Backend {
	if backend == nil {
		backend = &Backend{httpClient}
	}
	log.Printf("Backend: %+v", backend)
	return backend
}

type Backend struct {
	HTTPClient *http.Client
}

func (bc Backend) Call(method, path, subdomain, username, password string, form *url.Values, v interface{}) error {
	var body io.Reader
	if form != nil && len(*form) > 0 {
		data := form.Encode()
		if strings.ToUpper(method) == "GET" {
			path += "?" + data
		} else {
			body = bytes.NewBufferString(data)
		}
	}

	req, err := bc.NewRequest(method, path, subdomain, username, password, body)
	if err != nil {
		return err
	}

	if err := bc.Do(req, v); err != nil {
		return err
	}

	return nil
}

func (bc Backend) Do(req *http.Request, v interface{}) error {
	log.Printf("Requesting %v %v%v\n", req.Method, req.URL.Host, req.URL.Path)
	start := time.Now()

	res, err := bc.HTTPClient.Do(req)

	if debug {
		log.Printf("Completed in %v\n", time.Since(start))
	}

	if err != nil {
		log.Printf("Request to Desk failed: %v\n", err)
		return err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("Cannot parse Desk response: %v\n", err)
		return err
	}

	if res.StatusCode >= 400 {
		log.Printf("Error encountered from Desk: %s\n", resBody)
		return errors.New(string(resBody))
	}

	if debug {
		log.Printf("Desk response: %q\n", resBody)
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}
	return nil
}

func (bc Backend) NewRequest(method, path, subdomain, username, password string, body io.Reader) (*http.Request, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	url := apiURLPrefix + subdomain + apiURLSuffix + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("Cannot create Desk request: %v\n", err)
		return nil, err
	}

	req.SetBasicAuth(username, password)

	req.Header.Add("User-Agent", "Desk/v2 github.com/joncalhoun/desk-go/"+clientVersion)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	return req, nil
}
