package mixpanel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "http://api.mixpanel.com"

// Mixpanel is the main configuration struct
type Mixpanel struct {
	token      string
	apiURL     string
	httpClient HTTPClient
	people     People
}

// People implements the /engage endpoint functionalities
type People struct {
	token      string
	httpClient HTTPClient
}

// Props represents a key value pair used on the endpoints
type Props map[string]interface{}

// HTTPClient must be implemented to send HTTP requests
type HTTPClient interface {
	SendRequest(method string, endpoint string, data Props) (*http.Response, error)
}

// HTTPHandler is a Factory for HTTPClient implementation
type HTTPHandler struct {
	apiURL string
}

// NewMixpanel returns a new Mixpanel configuration
func NewMixpanel(token string, httpClient HTTPClient) *Mixpanel {
	if httpClient == nil {
		httpClient = &HTTPHandler{
			apiURL: apiURL,
		}
	}
	return &Mixpanel{
		token:      token,
		apiURL:     apiURL,
		httpClient: httpClient,
		people: People{
			token:      token,
			httpClient: httpClient,
		},
	}
}

// Track sends the event with distinct ID and event properties
func (mixpanel *Mixpanel) Track(event string, distinctID string, props Props) (*http.Response, error) {
	if distinctID != "" {
		props["distinct_id"] = distinctID
	}
	props["token"] = mixpanel.token
	data := map[string]interface{}{"event": event, "properties": props}
	_, err := mixpanel.httpClient.SendRequest("GET", "track", data)
	return nil, err
}

// Set sets existing or new user properties
func (people *People) Set(distinctID string, props Props) (*http.Response, error) {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$set"] = props
	_, err := people.httpClient.SendRequest("GET", "engage", keys)
	return nil, err
}

// Unset removes existing props from existing users
func (people *People) Unset(distinctID string, props []interface{}) (*http.Response, error) {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$unset"] = props
	_, err := people.httpClient.SendRequest("GET", "engage", keys)
	return nil, err
}

// Increment adds the specified value to an existing user property
func (people *People) Increment(distinctID string, props map[string]int) (*http.Response, error) {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$add"] = props
	res, err := people.httpClient.SendRequest("GET", "engage", keys)
	return res, err
}

// EncodeData is a helper function to encode Props in base64
func EncodeData(data Props) (string, error) {
	jsonEncoded, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	base64Data := base64.StdEncoding.EncodeToString(jsonEncoded)
	return base64Data, nil
}

// SendRequest sends the data to the specified Mixpanel endpoint
func (httpHandler *HTTPHandler) SendRequest(method string, endpoint string, data Props) (*http.Response, error) {
	base64Data, err := EncodeData(data)
	if err != nil {
		return nil, err
	}

	endpoint = fmt.Sprintf("%s/%s", httpHandler.apiURL, endpoint)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("data", base64Data)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}
	return res, nil
}
