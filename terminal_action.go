package squareup

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	terminalBasePath = "v2/terminals/actions"

	searchPath = "search"
)

// TerminalActionService is an interface for interfacing with the Square Terminal Action API
type TerminalActionService interface {
	Create(ctx context.Context, action *CreateTerminalActionEntry) (*GetTerminalAction, *Response, error)
	Search(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]SearchTerminalAction, *Response, error)
	Get(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error)
	Cancel(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error)
	Dismiss(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error)
}

var _ TerminalActionService = &TerminalActionServiceOp{}

type TerminalActionServiceOp struct {
	client *Client
}

type SearchTerminalAction struct {
	Action []TerminalActionEntry `json:"action"`
}

type GetTerminalAction struct {
	Action TerminalActionEntry `json:"action"`
}

type TerminalActionEntry struct {
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

type CreateTerminalActionEntry struct {
	Action         TerminalAction `json:"action"`
	IdempotencyKey string         `json:"idempotency_key"`
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

// TerminalActionConfirmationOptions represents the confirmation options for a terminal action
type TerminalActionConfirmationOptions struct {
	AgreeButtonText    string `json:"agree_button_text"`
	Body               string `json:"body"`
	Title              string `json:"title"`
	DisagreeButtonText string `json:"disagree_button_text"`
}

// TerminalActionDataCollectionOptions represents the data collection options for a terminal action
type TerminalActionDataCollectionOptions struct {
	Body      string `json:"body"`
	InputType string `json:"input_type"`
	Title     string `json:"title"`
}

type TerminalActionQrCodeOptions struct {
	BarcodeContents string `json:"barcode_contents"`
	Body            string `json:"body"`
	Title           string `json:"title"`
}

type TerminalActionReceiptOptions struct {
	PaymentId   string `json:"payment_id"`
	IsDuplicate bool   `json:"is_duplicate"`
	PrintOnly   bool   `json:"print_only"`
}

type TerminalActionSaveCardOptions struct {
	CustomerId  string `json:"customer_id"`
	ReferenceId string `json:"reference_id"`
}

type TerminalActionSignatureOptions struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

type TerminalActionOptions struct {
	ReferenceId string `json:"reference_id,omitempty"`
	Title       string `json:"title,omitempty"`
}

type TerminalActionSelectOptions struct {
	Options []TerminalActionOptions `json:"options"`
	Body    string                  `json:"body"`
}

type TerminalAction struct {
	ConfirmationOptions     TerminalActionConfirmationOptions   `json:"confirmation_options"`
	DataCollectionOptions   TerminalActionDataCollectionOptions `json:"data_collection_options"`
	QrCodeOptions           TerminalActionQrCodeOptions         `json:"qr_code_options"`
	ReceiptOptions          TerminalActionReceiptOptions        `json:"receipt_options"`
	SaveCardOptions         TerminalActionSaveCardOptions       `json:"save_card_options"`
	SignatureOptions        TerminalActionSignatureOptions      `json:"signature_options"`
	SelectOptions           TerminalActionSelectOptions         `json:"select_options"`
	AwaitNextAction         bool                                `json:"await_next_action"`
	AwaitNextActionDuration string                              `json:"await_next_action_duration"`
	DeadlineDuration        string                              `json:"deadline_duration"`
	DeviceId                string                              `json:"device_id"`
	Type                    string                              `json:"type"`
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

func (t *TerminalActionServiceOp) Search(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]SearchTerminalAction, *Response, error) {
	path := fmt.Sprintf("%s/%s", terminalBasePath, searchPath)
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, query)
	if err != nil {
		return nil, nil, err
	}

	root := new([]SearchTerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, err
}

func (t *TerminalActionServiceOp) Get(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error) {
	if len(actionId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s", terminalBasePath, actionId)

	req, err := t.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t *TerminalActionServiceOp) Create(ctx context.Context, action *CreateTerminalActionEntry) (*GetTerminalAction, *Response, error) {
	req, err := t.client.NewRequest(ctx, http.MethodPost, terminalBasePath, action)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t *TerminalActionServiceOp) Cancel(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error) {
	if len(actionId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/cancel", terminalBasePath, actionId)

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t *TerminalActionServiceOp) Dismiss(ctx context.Context, actionId string) (*GetTerminalAction, *Response, error) {
	if len(actionId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/dismiss", terminalBasePath, actionId)

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalAction)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
