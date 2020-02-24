package mixpanel

import (
	"net/http"
	"testing"
	"time"
)

type mockHTTPHandler struct {
	apiURL string
}

func (mock *mockHTTPHandler) SendRequest(method string, endpoint string, data Props) (*http.Response, error) {
	return nil, &http.Response{
		Body:   "Success",
		Status: 200,
	}
}

func TestCustomHTTPClient(t *testing.T) {
	mixpanel := NewMixpanel("token_here", &mockHTTPHandler{
		apiURL: "http://example.com",
	})

	got, err := mixpanel.Track("event", "123", map[string]interface{}{"time": time.Now()})
	want := &http.Response{
		Body:   "Success",
		Status: 200,
	}

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("\n got: %s\nwant: %s\n", got, want)
	}
}

func TestEncodeData(t *testing.T) {
	props := map[string]interface{}{"prop1": "value1", "props2": "value2"}

	got, err := EncodeData(props)
	want := "eyJwcm9wMSI6InZhbHVlMSIsInByb3BzMiI6InZhbHVlMiJ9"

	if err != nil {
		t.Fatal(err)
	}

	if got != want {
		t.Errorf("\n got: %s\nwant: %s\n", got, want)
	}
}
