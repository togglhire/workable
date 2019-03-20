package workable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	defaultBaseURL = "https://www.workable.com/"
	accessTokenURL = "https://www.workable.com/oauth/token"
	authorizeURL   = "https://www.workable.com/oauth/authorize"
)

// Client manages communication with the Workable API.
type Client struct {
	// httpcClient is the HTTP Client used to communicate with the API.
	httpClient *http.Client

	// OAuth, The access token you received once the OAuth process is complete and the user grants the partner permission to access their data
	accessToken  *AccessTokenOutput
	clientID     string
	clientSecret string
	redirectURI  string

	// BaseURL is the base url for api requests.
	baseURL string

	// Services used for talking with different parts of the Workable API
	OAuth oauthService
}

// NewClient returns a new instance of *Client.
func NewClient(clientID, clientSecret, redirectURI string, accessToken *AccessTokenOutput, httpClient *http.Client) *Client {
	return newClient(clientID, clientSecret, redirectURI, accessToken, httpClient)
}

// SetAccessToken updates the access token used for accessing API endpoints
func (c *Client) SetAccessToken(accessToken *AccessTokenOutput) {
	c.accessToken = accessToken
}

func newClient(clientID, clientSecret, redirectURI string, accessToken *AccessTokenOutput, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		httpClient:   httpClient,
		accessToken:  accessToken,
		baseURL:      defaultBaseURL,
		clientID:     clientID,
		clientSecret: clientSecret,
		redirectURI:  redirectURI,
	}

	//Services
	client.OAuth = &oauthServiceImpl{
		client: client,
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
	requestURL := fmt.Sprintf("%sv1/%s", c.baseURL, endpoint)

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

	if c.accessToken != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.accessToken.AccessToken))
	}

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

type workableErrorResponse struct {
	Error struct {
		Error Error `json:"error`
	} `json:"error"`
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

	if r, err := isError(resp); r && err == nil {
		workableError := workableErrorResponse{}
		err = readJSON(resp.Body, &workableError)
		if err != nil {
			return err
		}
		if r, err = isClientError(resp); r && err == nil {
			clientError := ClientError{
				StatusCode:   resp.StatusCode,
				ErrorMessage: workableError.Error.Error,
			}
			return clientError
		}
		if r, err = isServerError(resp); r && err == nil {
			serverError := ServerError{
				StatusCode:   resp.StatusCode,
				ErrorMessage: workableError.Error.Error,
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
