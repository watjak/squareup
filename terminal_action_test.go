package squareup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

// TestTerminalActionClient_SearchTerminalAction tests the SearchTerminalAction method.
func TestTerminalActionClient_SearchTerminalAction(t *testing.T) {
	setup()
	defer teardown()
	request := &TerminalActionQuery{}
	mux.HandleFunc("/v2/terminals/actions/search", func(w http.ResponseWriter, r *http.Request) {
		v := new(TerminalActionQuery)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}
		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, request) {
			t.Errorf("Request body = %+v, expected %+v", v, request)
		}
		response := `
		{
		  "action": [
			{
			  "id": "TrwuBUPu2IpqO",
			  "device_id": "{{device_id}}",
			  "status": "CANCELED",
			  "cancel_reason": "TIMED_OUT",
			  "location_id": "LMPCZVC1C3FHM",
			  "type": "CHECKOUT",
			  "app_id": "sandbox-sq0idb-ShZkaDWHMnXDLk4jZVskKw",
			  "checkout_options": {
				"amount_money": {
				  "amount": 100,
				  "currency": "USD"
				},
				"reference_id": "232323223",
				"note": "hamberger",
				"device_options": {
				  "tip_settings": {
					"separate_tip_screen": true,
					"custom_tip_field": true,
					"allow_tipping": true
				  },
				  "skip_receipt_screen": true
				},
				"payment_type": "CARD_PRESENT",
				"payment_options": {
				  "autocomplete": true
				}
			  }
			}
		  ]
		}`
		fmt.Fprint(w, response)
	})

	action, _, err := client.TerminalAction.Search(ctx, nil, nil)
	if err != nil {
		t.Errorf("TerminalAction.Search returned error: %v", err)
	}
	var expected = []TerminalActionEntry{
		{
			Id:           "TrwuBUPu2IpqO",
			DeviceId:     "{{device_id}}",
			Status:       "CANCELED",
			CancelReason: "TIMED_OUT",
			LocationId:   "LMPCZVC1C3FHM",
			Type:         "CHECKOUT",
			AppId:        "sandbox-sq0idb-ShZkaDWHMnXDLk4jZVskKw",
			CheckoutOptions: CheckoutOptions{
				AmountMoney: AmountMoney{
					Amount:   100,
					Currency: "USD",
				},
				ReferenceId: "232323223",
				Note:        "hamberger",
				DeviceOptions: DeviceOptions{
					TipSettings: TipSettings{
						SeparateTipScreen: true,
						CustomTipField:    true,
						AllowTipping:      true,
					},
					SkipReceiptScreen: true,
				},
				PaymentType: "CARD_PRESENT",
				PaymentOptions: PaymentOptions{
					Autocomplete: true,
				},
			},
		},
	}
	if !reflect.DeepEqual(action, expected) {
		t.Errorf("Terminal.Action.Search returned %+v, expected %+v", action, expected)
	}
}
