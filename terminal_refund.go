package squareup

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

const (
	terminalRefundBasePath = "v2/terminals/refunds"

	terminalRefundSearchPath = "search"
)

type TerminalRefundService interface {
	CreateTerminalRefund(ctx context.Context, refund *CreateTerminalRefundEntry) (*GetTerminalRefund, *Response, error)
	SearchTerminalRefund(ctx context.Context, options *ListOptions, query *TerminalRefundQuery) ([]SearchTerminalRefund, *Response, error)
	GetTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error)
	CancelTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error)
	DismissTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error)
}

var _ TerminalRefundService = &TerminalRefundServiceOp{}

type TerminalRefundServiceOp struct {
	client *Client
}
type SearchTerminalRefund struct {
	Refunds []TerminalRefundEntry `json:"refunds"`
}
type GetTerminalRefund struct {
	Refund *TerminalRefundEntry `json:"refund"`
}
type TerminalRefundEntry struct {
	Id               string       `json:"id"`
	PaymentId        string       `json:"payment_id"`
	AmountMoney      *AmountMoney `json:"amount_money"`
	Reason           string       `json:"reason"`
	DeviceId         string       `json:"device_id"`
	DeadlineDuration string       `json:"deadline_duration"`
	Status           string       `json:"status,omitempty"`
	CancelReason     string       `json:"cancel_reason"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
	AppId            string       `json:"app_id"`
	Card             *Card        `json:"card"`
	OrderId          string       `json:"order_id"`
	LocationId       string       `json:"location_id"`
}
type CreateTerminalRefundEntry struct {
	Refund         *TerminalRefund `json:"refund"`
	IdempotencyKey string          `json:"idempotency_key"`
}

type TerminalRefund struct {
	AmountMoney      *AmountMoney `json:"amount_money"`
	DeviceId         string       `json:"device_id"`
	PaymentId        string       `json:"payment_id"`
	Reason           string       `json:"reason"`
	DeadlineDuration string       `json:"deadline_duration,omitempty"`
}
type TerminalRefundQuery struct {
	Sort struct {
		SortOrder string `json:"sort_order,omitempty"`
	} `json:"sort"`
	Filter struct {
		CreatedAt struct {
			EndAt   string `json:"end_at,omitempty"`
			StartAt string `json:"start_at,omitempty"`
		} `json:"created_at"`
	} `json:"filter"`
}

func (t TerminalRefundServiceOp) CreateTerminalRefund(ctx context.Context, refund *CreateTerminalRefundEntry) (*GetTerminalRefund, *Response, error) {
	req, err := t.client.NewRequest(ctx, http.MethodPost, terminalActionBasePath, refund)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalRefund)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t TerminalRefundServiceOp) SearchTerminalRefund(ctx context.Context, options *ListOptions, query *TerminalRefundQuery) ([]SearchTerminalRefund, *Response, error) {

	path := fmt.Sprintf("%s/%s", terminalRefundBasePath, terminalRefundSearchPath)
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new([]SearchTerminalRefund)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return *root, resp, err
}

func (t TerminalRefundServiceOp) GetTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error) {
	if len(refundId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s", terminalRefundBasePath, refundId)

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalRefund)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t TerminalRefundServiceOp) CancelTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error) {
	if len(refundId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/cancel", terminalRefundBasePath, refundId)

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalRefund)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

func (t TerminalRefundServiceOp) DismissTerminalRefund(ctx context.Context, refundId string) (*GetTerminalRefund, *Response, error) {
	if len(refundId) == 0 {
		return nil, nil, NewArgError("actionId", "cannot be an empty string")
	}

	path := fmt.Sprintf("%s/%s/dismiss", terminalRefundBasePath, refundId)

	req, err := t.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GetTerminalRefund)
	resp, err := t.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
