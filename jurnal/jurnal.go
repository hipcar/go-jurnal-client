package jurnal

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	libraryVersion      = "0.1"
	userAgent           = "jurnal-api/" + libraryVersion
	jurnalAPIURL        = "http://api.jurnal.id/core/api/v1/"
	jurnalSandBoxAPIURL = "http://sandbox-api.jurnal.id/core/api/v1/"
)

type JurnalEnvironment string

const (
	Sandbox    JurnalEnvironment = "sandbox"
	Production JurnalEnvironment = "production"
)

var (
	// ErrUnauthorized can be returned on any call on response status code 401.
	ErrUnauthorized = errors.New("jurnal-api: unauthorized")
)

type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	doer      Doer
	baseURL   *url.URL
	UserAgent string
	EnableLog bool
	APIKey    string

	JournalEntry *JournalEntry

	httpClient *http.Client
}

type errorResponse struct {
	Errors            string   `json:"errors"`
	ErrorFullMessages []string `json:"error_full_messages"`
}

type DoerFunc func(req *http.Request) (resp *http.Response, err error)

// NewClient created new jurnal api client with doer.
// If doer is nil then http.DefaultClient used instead.
func NewClient(env JurnalEnvironment) *Client {
	var baseUrl string

	if env == Sandbox {
		baseUrl = jurnalSandBoxAPIURL
	} else {
		baseUrl = jurnalAPIURL
	}

	baseURL, _ := url.Parse(baseUrl)
	client := &Client{
		doer:      http.DefaultClient,
		baseURL:   baseURL,
		UserAgent: userAgent,
	}

	client.JournalEntry = &JournalEntry{client}

	return client
}

func (c *Client) Request(method string, path string, data interface{}, v interface{}) error {

	urlStr := path

	rel, err := url.Parse(urlStr)
	if err != nil {
		return err
	}
	u := c.baseURL.ResolveReference(rel)
	var body io.Reader

	if data != nil {
		b, err := json.Marshal(data)

		if err != nil {
			return err
		}
		body = bytes.NewReader(b)

		if c.EnableLog {
			fmt.Printf("Request %s to %s with data: %s \n", method, u.String(), string(b))
		}
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", c.APIKey)

	resp, err := c.doer.Do(req.WithContext(context.Background()))
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return ErrUnauthorized
	}

	// Return error from jurnal API
	if resp.StatusCode >= 300 {
		rb := new(errorResponse)

		err = json.NewDecoder(resp.Body).Decode(rb)

		errMsg := ""

		if rb.Errors != "" {
			errMsg += rb.Errors + ". "
		}

		for _, errorFullMessage := range rb.ErrorFullMessages {
			errMsg += errorFullMessage + ". "
		}

		if errMsg == "" {
			errMsg = fmt.Sprintf("general error with status code %d", resp.StatusCode)
		}

		return errors.New(errMsg)
	}

	// Decode to interface
	res := v
	err = json.NewDecoder(resp.Body).Decode(res)

	by, _ := json.Marshal(res)
	if c.EnableLog {
		fmt.Printf("Response %s from %s : %s \n", method, u.String(), string(by))
	}

	return err
}
