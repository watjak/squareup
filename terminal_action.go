package squareup

import (
	"context"
	"fmt"
	"time"
)

const (
	terminalBasePath = "v2/terminals/actions"

	searchPath = "search"
)

// TerminalActionService is an interface for interfacing with the Square Terminal Action API
type TerminalActionService interface {
	Search(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]TerminalAction, *Response, error)
	Get(ctx context.Context, actionId string) (*TerminalAction, *Response, error)
}

var _ TerminalActionService = &TerminalActionServiceOp{}

type TerminalActionServiceOp struct {
	client *Client
}

type TerminalAction struct {
	Id              string          `json:"id"`
	DeviceId        string          `json:"device_id"`
	Status          string          `json:"status"`
	CancelReason    string          `json:"cancel_reason"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	LocationId      string          `json:"location_id"`
	Type            string          `json:"type"`
	AppId           string          `json:"app_id"`
	CheckoutOptions CheckoutOptions `json:"checkout_options"`
}

type AmountMoney struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type DeviceOptions struct {
	TipSettings       TipSettings `json:"tip_settings"`
	SkipReceiptScreen bool        `json:"skip_receipt_screen"`
}

type TipSettings struct {
	SeparateTipScreen bool `json:"separate_tip_screen"`
	CustomTipField    bool `json:"custom_tip_field"`
	AllowTipping      bool `json:"allow_tipping"`
}

type PaymentOptions struct {
	Autocomplete bool `json:"autocomplete"`
}

type CheckoutOptions struct {
	AmountMoney    `json:"amount_money"`
	ReferenceId    string `json:"reference_id"`
	Note           string `json:"note"`
	DeviceOptions  `json:"device_options"`
	PaymentType    string `json:"payment_type"`
	PaymentOptions `json:"payment_options"`
}

// TerminalActionQuery represents the query parameters for the Search method
type TerminalActionQuery struct {
	Filter struct {
		DeviceId string `json:"device_id"`
		Status   string `json:"status"`
		Type     string `json:"type"`
	} `json:"filter"`
	Sort struct {
		SortOrder string `json:"sort_order"`
	} `json:"sort"`
}

func (t TerminalActionServiceOp) Search(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]TerminalAction, *Response, error) {
	path := fmt.Sprintf("%s/%s", terminalBasePath, searchPath)
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}
	req, err := t.client.NewRequest(ctx, "POST", path, query)
	if err != nil {
		return nil, nil, err
	}

	root := new([]TerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, err
}

func (t TerminalActionServiceOp) Get(ctx context.Context, actionId string) (*TerminalAction, *Response, error) {
	if len(actionId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s", terminalBasePath, actionId)

	req, err := t.client.NewRequest(ctx, "GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(TerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}