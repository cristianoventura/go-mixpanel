# go-mixpanel
Go implementation of the Mixpanel API

## Usage

```go
// Initialize
mixpanel := NewMixpanel("your_token_here", nil)

// Track Events
eventName := "My Event"
distinctID := "123"
props := map[string]interface{}{"property one": "value one", "property N": 20}
res, err := mixpanel.Track(eventName, distinctID, props)

// Set User properties
props := map[string]interface{}{"Name": "User Name", "Age": 26}
res, err := mixpanel.people.Set(distinctID, props)

// Unset User properties
propsToUnset := []interface{}{"Name",}
res, err := mixpanel.people.Unset(distinctID, propsToUnset)

// Incremnet User properties
res, err := mixpanel.people.Increment(distinctID, map[string]int{"Age": 1})
```

## Implementing your own http handler

Create your own struct and implement the `SendRequest` function.

Example:

```go
type MyHTTPHandler struct {
  apiURL string
}

func (myHTTP *MyHTTPHandler) SendRequest(method string, endpoint string, data Props) (*http.Response, error) {
  // implementation goes here
}
```

Instantiate Mixpanel passing your struct

```go
mixpanel := NewMixpanel("token_here", &MyHTTPHandler{
  apiURL: "http://my-own-endpoint",
})
```
