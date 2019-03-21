package workable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultDomain    = "workable"
	sandboxDomain    = "workablesandbox"
	defaultSubdomain = "www"

	defaultBaseURL = "https://{subdomain}.{domain}.com"
)

// Client manages communication with the Workable API.
type Client struct {
	// httpcClient is the HTTP Client used to communicate with the API.
	httpClient *http.Client

	// OAuth, The access token you received once the OAuth process is complete and the user grants the partner permission to access their data
	token Token

	baseURL string

	// Sandbox or Live
	domain string
}

// NewClient returns a new instance of *Client.
func NewClient(token Token, httpClient *http.Client) *Client {
	return newClient(defaultBaseURL, defaultDomain, token, httpClient)
}

// NewSandboxClient returns a new instance of *Client that connects to the sandbox version of workable.
func NewSandboxClient(token Token, httpClient *http.Client) *Client {
	return newClient(defaultBaseURL, sandboxDomain, token, httpClient)
}

func (c *Client) Accounts() accountService {
	return &accountServiceImpl{
		client: c,
	}
}

func (c *Client) OAuth(info OAuthServiceInput) oauthService {
	return &oauthServiceImpl{
		client: c,
		info:   info,
	}
}

func newClient(baseURL, domain string, token Token, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		httpClient: httpClient,
		token:      token,
		domain:     domain,
		baseURL:    baseURL,
	}

	return client
}

// ReadJSON reads the json value into the v param. Can only read once!
func readJSON(r io.ReadCloser, v interface{}) error {
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&v)
	return err
}

// Params are used to send parameters with the request.
type Params map[string]interface{}

// newRequest creates an authenticated API request that is ready to send.
func (c *Client) newRequest(method string, endpoint string, params Params, body interface{}) (*http.Request, error) {
	method = strings.ToUpper(method)
	requestURL := c.baseURL + "/spi/v3/{endpoint}"
	requestURL = strings.Replace(requestURL, "{subdomain}", "www", -1)
	requestURL = strings.Replace(requestURL, "{domain}", c.domain, -1)
	requestURL = strings.Replace(requestURL, "{endpoint}", endpoint, -1)

	// Query String
	qs := url.Values{}
	for k, v := range params {
		qs.Add(k, fmt.Sprintf("%v", v))
	}

	if len(qs) > 0 {
		requestURL += "?" + qs.Encode()
	}

	// Request body
	var buf bytes.Buffer
	if body != nil {
		err := json.NewEncoder(&buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, requestURL, &buf)
	if err != nil {
		return req, err
	}

	if c.token.AccessToken != "" {
		print(c.token.AccessToken)
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token.AccessToken))
	}

	print(req.URL.String())
	if req.Method == "POST" || req.Method == "PUT" {
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
	}

	return req, err
}

// do takes a prepared API request and makes the API call to Workable.
// It will decode the JSON into a destination struct you provide as well
// as parse any validation errors that may have occurred.
// It returns a Response object that provides a wrapper around http.Response
// with some convenience methods.
func (c *Client) do(req *http.Request, v interface{}) error {
	return do(c.httpClient, req, v)
}

func do(client *http.Client, req *http.Request, v interface{}) error {
	req.Close = true
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp == nil {
		return ErrShouldNotBeNil
	}
	defer resp.Body.Close()

	logBody(resp)
	if r, err := isError(resp); r && err == nil {
		workableError := Error{}
		err = readJSON(resp.Body, &workableError)
		if err != nil {
			return err
		}
		if r, err = isClientError(resp); r && err == nil {
			clientError := ClientError{
				StatusCode:   resp.StatusCode,
				ErrorMessage: workableError,
			}
			return clientError
		}
		if r, err = isServerError(resp); r && err == nil {
			serverError := ServerError{
				StatusCode:   resp.StatusCode,
				ErrorMessage: workableError,
			}
			return serverError
		}
	} else if err != nil {
		return err
	}

	err = readJSON(resp.Body, &v)
	return err
}

func interfaceToCSV(a interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), ","), "[]")
}
func spaceDelimit(a interface{}) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), " "), "[]")
}

func formatReadCloser(r *io.ReadCloser) string {
	if r == nil {
		return ""
	}
	body, err := ioutil.ReadAll(*r)
	if err != nil {
		return ""
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	*r = rdr1 // restore body

	return string(body)
}

func logBody(r *http.Response) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("ERR:", err.Error())
		return
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(body))
	r.Body.Close()
	r.Close = true
	r.Body = rdr1 // restore body
	log.Println("BODY:", string(body))
}
