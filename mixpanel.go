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
	token  string
	apiURL string
	people People
}

// People implements the /engage endpoint functionalities
type People struct {
	token string
}

// Props represents a key value pair used on the endpoints
type Props map[string]interface{}

// NewMixpanel returns a new Mixpanel configuration
func NewMixpanel(token string) *Mixpanel {
	people := People{
		token: token,
	}
	return &Mixpanel{
		token:  token,
		apiURL: apiURL,
		people: people,
	}
}

// Track sends the event with distinct ID and event properties
func (mixpanel *Mixpanel) Track(event string, distinctID string, props Props) error {
	if distinctID != "" {
		props["distinct_id"] = distinctID
	}
	props["token"] = mixpanel.token
	data := map[string]interface{}{"event": event, "properties": props}
	err := SendRequest("GET", "track", data)
	return err
}

// Set sets existing or new user properties
func (people *People) Set(distinctID string, props Props) error {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$set"] = props
	err := SendRequest("GET", "engage", keys)
	return err
}

// Unset removes existing props from existing users
func (people *People) Unset(distinctID string, props []interface{}) error {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$unset"] = props
	err := SendRequest("GET", "engage", keys)
	return err
}

// Increment adds the specified value to an existing user property
func (people *People) Increment(distinctID string, props map[string]int) error {
	keys := Props{}
	if distinctID != "" {
		keys["$distinct_id"] = distinctID
	}
	keys["$token"] = people.token
	keys["$add"] = props
	err := SendRequest("GET", "engage", keys)
	return err
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
func SendRequest(method string, endpoint string, data Props) error {
	base64Data, err := EncodeData(data)
	if err != nil {
		return err
	}

	endpoint = fmt.Sprintf("%s/%s", apiURL, endpoint)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("data", base64Data)
	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	res.Body.Close()
	return nil
}
