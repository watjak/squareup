package squareup

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	terminalCheckoutBasePath = "v2/terminals/checkouts"

	terminalCheckoutSearchPath = "search"
)

// TerminalCheckoutService is an interface for interfacing with the Square Terminal API
type TerminalCheckoutService interface {
	CreateTerminalCheckout(ctx context.Context, checkout *CreateTerminalCheckoutEntry) (*GetTerminalCheckout, *Response, error)
	SearchTerminalCheckout(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]SearchTerminalCheckout, *Response, error)
	GetTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error)
	CancelTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error)
	DismissTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error)
}

var _ TerminalCheckoutService = &TerminalCheckoutServiceOp{}

type TerminalCheckoutServiceOp struct {
	client *Client
}
type SearchTerminalCheckout struct {
	Checkouts []TerminalCheckoutEntry `json:"checkouts"`
}

type GetTerminalCheckout struct {
	Checkout *TerminalCheckoutEntry `json:"checkout"`
}

type TerminalCheckoutEntry struct {
	Id             string          `json:"id"`
	AmountMoney    *AmountMoney    `json:"amount_money"`
	ReferenceId    string          `json:"reference_id"`
	Note           string          `json:"note"`
	DeviceOptions  *DeviceOptions  `json:"device_options"`
	Status         string          `json:"status"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
	AppId          string          `json:"app_id"`
	CancelReason   string          `json:"cancel_reason"`
	LocationId     string          `json:"location_id"`
	PaymentType    string          `json:"payment_type"`
	PaymentOptions *PaymentOptions `json:"payment_options"`
}

type CreateTerminalCheckoutEntry struct {
	IdempotencyKey string            `json:"idempotency_key"`
	Checkout       *TerminalCheckout `json:"checkout"`
}
type TerminalCheckout struct {
	AmountMoney                    *AmountMoney    `json:"amount_money"`
	DeviceOptions                  *DeviceOptions  `json:"device_options"`
	AppFeeMoney                    *AmountMoney    `json:"app_fee_money"`
	TipMoney                       *AmountMoney    `json:"tip_money"`
	CustomerId                     string          `json:"customer_id"`
	Note                           string          `json:"note"`
	OrderId                        string          `json:"order_id"`
	PaymentOptions                 *PaymentOptions `json:"payment_options"`
	PaymentType                    string          `json:"payment_type"`
	ReferenceId                    string          `json:"reference_id"`
	StatementDescriptionIdentifier string          `json:"statement_description_identifier"`
	TeamMemberId                   string          `json:"team_member_id"`
}

func (s *TerminalCheckoutServiceOp) CreateTerminalCheckout(ctx context.Context, checkout *CreateTerminalCheckoutEntry) (*GetTerminalCheckout, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, terminalCheckoutBasePath, checkout)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalCheckout)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (s *TerminalCheckoutServiceOp) SearchTerminalCheckout(ctx context.Context, options *ListOptions, query *TerminalActionQuery) ([]SearchTerminalCheckout, *Response, error) {
	path := fmt.Sprintf("%s/%s", terminalCheckoutBasePath, terminalCheckoutSearchPath)
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, query)
	if err != nil {
		return nil, nil, err
	}

	root := new([]SearchTerminalCheckout)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, err
}

func (s *TerminalCheckoutServiceOp) GetTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error) {
	if len(checkoutId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s", terminalCheckoutBasePath, checkoutId)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalCheckout)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (s *TerminalCheckoutServiceOp) CancelTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error) {
	if len(checkoutId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/cancel", terminalCheckoutBasePath, checkoutId)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalCheckout)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (s *TerminalCheckoutServiceOp) DismissTerminalCheckout(ctx context.Context, checkoutId string) (*GetTerminalCheckout, *Response, error) {
	if len(checkoutId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/dismiss", terminalCheckoutBasePath, checkoutId)

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalCheckout)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
