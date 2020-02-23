package mixpanel

import (
	"testing"
)

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
