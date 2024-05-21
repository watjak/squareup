package squareup

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type EventType string

const (
	EventTypeOauthAuthorizeRevoked EventType = "oauth.authorization.revoked"
	EventTypeBackAccountCreated    EventType = "bank_account.created"
	EventTypeBackAccountDisabled   EventType = "bank_account.disabled"
	EventTypeBackAccountVerified   EventType = "bank_account.verified"

	EventTypePaymentCreated EventType = "payment.created"
	EventTypePaymentUpdated EventType = "payment.updated"
)

type EventData struct {
	// Object is a raw mapping of the API resource contained in the event.
	// Although marked with json:"-", it's still populated independently by
	// a custom UnmarshalJSON implementation.
	// Object containing the API resource relevant to the event. For example, an `invoice.created` event will have a full [invoice object](https://stripe.com/docs/api#invoice_object) as the value of the object key.
	Object map[string]interface{} `json:"-"`
	// Object containing the names of the updated attributes and their values prior to the event (only included in events of type `*.updated`). If an array attribute has any updated elements, this object contains the entire array. In Stripe API versions 2017-04-06 or earlier, an updated array attribute in this object includes only the updated array elements.
	PreviousAttributes map[string]interface{} `json:"previous_attributes"`
	Raw                json.RawMessage        `json:"object"`
}

type Event struct {
	MerchantID string     `json:"merchant_id"`
	Type       EventType  `json:"type"`
	EventID    string     `json:"event_id"`
	Data       *EventData `json:"data"`
}

// GetObjectValue returns the value from the e.Data.Object bag based on the keys hierarchy.
func (e *Event) GetObjectValue(keys ...string) string {
	return getValue(e.Data.Object, keys)
}

// GetPreviousValue returns the value from the e.Data.Prev bag based on the keys hierarchy.
func (e *Event) GetPreviousValue(keys ...string) string {
	return getValue(e.Data.PreviousAttributes, keys)
}

// UnmarshalJSON handles deserialization of the EventData.
// This custom unmarshaling exists so that we can keep both the map and raw data.
func (e *EventData) UnmarshalJSON(data []byte) error {
	type eventdata EventData
	var ee eventdata
	err := json.Unmarshal(data, &ee)
	if err != nil {
		return err
	}

	*e = EventData(ee)
	return json.Unmarshal(e.Raw, &e.Object)
}

// getValue returns the value from the m map based on the keys.
func getValue(m map[string]interface{}, keys []string) string {
	node := m[keys[0]]

	for i := 1; i < len(keys); i++ {
		key := keys[i]

		sliceNode, ok := node.([]interface{})
		if ok {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				panic(fmt.Sprintf(
					"Cannot access nested slice element with non-integer key: %s",
					key))
			}
			node = sliceNode[intKey]
			continue
		}

		mapNode, ok := node.(map[string]interface{})
		if ok {
			node = mapNode[key]
			continue
		}

		panic(fmt.Sprintf(
			"Cannot descend into non-map non-slice object with key: %s", key))
	}

	if node == nil {
		return ""
	}

	return fmt.Sprintf("%v", node)
}
