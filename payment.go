package squareup

import (
	"context"
	"net/http"
	"time"
)

const (
	PaymentBasePath = "v2/payments"
)

type PaymentService interface {
	List(ctx context.Context, options *ListOptions) ([]ListPayments, *Response, error)
}

var _ PaymentService = &PaymentServiceOp{}

type PaymentServiceOp struct {
	client *Client
}

type ListPayments struct {
	Payment []PaymentEntry `json:"payment"`
}

type Payment struct {
	Payment PaymentEntry `json:"payment"`
}

type PaymentEntry struct {
	Id                 string            `json:"id"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
	AmountMoney        AmountMoney       `json:"amount_money"`
	Status             string            `json:"status"`
	SourceType         string            `json:"source_type"`
	LocationId         string            `json:"location_id"`
	OrderId            string            `json:"order_id"`
	Note               string            `json:"note"`
	TotalMoney         AmountMoney       `json:"total_money"`
	Capabilities       []string          `json:"capabilities"`
	ExternalDetails    ExternalDetail    `json:"external_details"`
	ReceiptNumber      string            `json:"receipt_number"`
	ReceiptUrl         string            `json:"receipt_url"`
	ApplicationDetails ApplicationDetail `json:"application_details"`
	VersionToken       string            `json:"version_token"`
}

type ApplicationDetail struct {
	SquareProduct string `json:"square_product"`
	ApplicationId string `json:"application_id"`
}

type ExternalDetail struct {
	Type   string `json:"type"`
	Source string `json:"source"`
}

func (p *PaymentServiceOp) List(ctx context.Context, options *ListOptions) ([]ListPayments, *Response, error) {
	path := PaymentBasePath
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := p.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var payments []ListPayments
	resp, err := p.client.Do(ctx, req, &payments)
	if err != nil {
		return nil, resp, err
	}

	return payments, resp, nil
}
