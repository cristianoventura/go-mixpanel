# go-mixpanel
Go implementation of the Mixpanel API

## Usage

```go
// Initialize
mixpanel := NewMixpanel("your_token_here")

// Track Events
eventName := "My Event"
distinctID := "123"
props := map[string]interface{}{"property one": "value one", "property N": 20}
err := mixpanel.Track(eventName, distinctID, props)

// Set User properties
props := map[string]interface{}{"Name": "User Name", "Age": 26}
err := mixpanel.people.Set(distinctID, props)

// Unset User properties
propsToUnset := []interface{}{"Name",}
err := mixpanel.people.Unset(distinctID, propsToUnset)

// Incremnet User properties
err := mixpanel.people.Increment(distinctID, map[string]int{"Age": 1})
```
