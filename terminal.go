package squareup

import "time"

const (
	terminalBasePath = "v2/terminals"

	terminalSearchPath = "search"
)

// TerminalService is an interface for interfacing with the Square Terminal API
type TerminalService interface {
}

var _ TerminalService = &TerminalServiceOp{}

type TerminalServiceOp struct {
	client *Client
}

type SearchTerminalCheckout struct {
	Checkouts []TerminalCheckoutEntry `json:"checkouts"`
}

type GetTerminalCheckout struct {
	Checkout TerminalCheckoutEntry `json:"checkout"`
}

type TerminalCheckoutEntry struct {
	Id             string         `json:"id"`
	AmountMoney    AmountMoney    `json:"amount_money"`
	ReferenceId    string         `json:"reference_id"`
	Note           string         `json:"note"`
	DeviceOptions  DeviceOptions  `json:"device_options"`
	Status         string         `json:"status"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	AppId          string         `json:"app_id"`
	CancelReason   string         `json:"cancel_reason"`
	LocationId     string         `json:"location_id"`
	PaymentType    string         `json:"payment_type"`
	PaymentOptions PaymentOptions `json:"payment_options"`
}

type CreateTerminalCheckout struct {
	IdempotencyKey string `json:"idempotency_key"`
	Checkout       struct {
		AmountMoney struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"amount_money"`
		DeviceOptions struct {
			DeviceId          string `json:"device_id"`
			CollectSignature  bool   `json:"collect_signature"`
			ShowItemizedCart  bool   `json:"show_itemized_cart"`
			SkipReceiptScreen bool   `json:"skip_receipt_screen"`
			TipSettings       struct {
				AllowTipping      bool  `json:"allow_tipping"`
				CustomTipField    bool  `json:"custom_tip_field"`
				SeparateTipScreen bool  `json:"separate_tip_screen"`
				SmartTipping      bool  `json:"smart_tipping"`
				TipPercentages    []int `json:"tip_percentages"`
			} `json:"tip_settings"`
		} `json:"device_options"`
		AppFeeMoney struct {
			Currency string `json:"currency"`
			Amount   int    `json:"amount"`
		} `json:"app_fee_money"`
		TipMoney struct {
			Amount   int    `json:"amount"`
			Currency string `json:"currency"`
		} `json:"tip_money"`
		CustomerId     string `json:"customer_id"`
		Note           string `json:"note"`
		OrderId        string `json:"order_id"`
		PaymentOptions struct {
			AcceptPartialAuthorization bool `json:"accept_partial_authorization"`
			Autocomplete               bool `json:"autocomplete"`
		} `json:"payment_options"`
		PaymentType                    string `json:"payment_type"`
		ReferenceId                    string `json:"reference_id"`
		StatementDescriptionIdentifier string `json:"statement_description_identifier"`
		TeamMemberId                   string `json:"team_member_id"`
	} `json:"checkout"`
}
