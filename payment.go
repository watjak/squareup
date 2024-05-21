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
	ListPayment(ctx context.Context, options *ListOptions) ([]ListPayments, *Response, error)
	CreatePayment(ctx context.Context, payment *Payment) (*Payment, *Response, error)
	CancelByIdempotencyKey(ctx context.Context, id string) (*Payment, *Response, error)
	GetPayment(ctx context.Context, paymentId string) (*Payment, *Response, error)
	UpdatePayment(ctx context.Context, paymentId string, payment *Payment) (*Payment, *Response, error)
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
	Type           string      `json:"type"`
	Source         string      `json:"source"`
	SourceFeeMoney AmountMoney `json:"source_fee_money"`
	SourceId       string      `json:"source_id"`
}

func (p *PaymentServiceOp) ListPayment(ctx context.Context, options *ListOptions) ([]ListPayments, *Response, error) {
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

type CreatePayment struct {
	IdempotencyKey                 string         `json:"idempotency_key"`
	SourceId                       string         `json:"source_id"`
	AmountMoney                    AmountMoney    `json:"amount_money"`
	AppFeeMoney                    AmountMoney    `json:"app_fee_money"`
	BillingAddress                 BillingAddress `json:"billing_address"`
	CashDetails                    CashDetails    `json:"cash_details"`
	TipMoney                       AmountMoney    `json:"tip_money"`
	ShippingAddress                BillingAddress `json:"shipping_address"`
	AcceptPartialAuthorization     bool           `json:"accept_partial_authorization"`
	Autocomplete                   bool           `json:"autocomplete"`
	BuyerEmailAddress              string         `json:"buyer_email_address"`
	CustomerDetails                CustomerDetail `json:"customer_details"`
	CustomerId                     string         `json:"customer_id"`
	DelayAction                    string         `json:"delay_action"`
	DelayDuration                  string         `json:"delay_duration"`
	ExternalDetails                ExternalDetail `json:"external_details"`
	LocationId                     string         `json:"location_id"`
	Note                           string         `json:"note"`
	OrderId                        string         `json:"order_id"`
	ReferenceId                    string         `json:"reference_id"`
	StatementDescriptionIdentifier string         `json:"statement_description_identifier"`
	TeamMemberId                   string         `json:"team_member_id"`
	VerificationToken              string         `json:"verification_token"`
}

type CashDetails struct {
	BuyerSuppliedMoney AmountMoney `json:"buyer_supplied_money"`
	ChangeBackMoney    AmountMoney `json:"change_back_money"`
}

func (p *PaymentServiceOp) CreatePayment(ctx context.Context, payment *Payment) (*Payment, *Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PaymentServiceOp) GetPayment(ctx context.Context, paymentId string) (*Payment, *Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PaymentServiceOp) CancelByIdempotencyKey(ctx context.Context, id string) (*Payment, *Response, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PaymentServiceOp) UpdatePayment(ctx context.Context, paymentId string, payment *Payment) (*Payment, *Response, error) {
	//TODO implement me
	panic("implement me")
}
