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
	CreatePayment(ctx context.Context, payment *CreatePayment) (*Payment, *Response, error)
	CancelByIdempotencyKey(ctx context.Context, id string) (*Payment, *Response, error)
	GetPayment(ctx context.Context, paymentId string) (*Payment, *Response, error)
	UpdatePayment(ctx context.Context, paymentId string, payment *Payment) (*Payment, *Response, error)
	CancelPayment(ctx context.Context, paymentId string) (*Payment, *Response, error)
	CompletePayment(ctx context.Context, paymentId, versionToken string) (*Payment, *Response, error)
}

var _ PaymentService = &PaymentServiceOp{}

type PaymentServiceOp struct {
	client *Client
}

type ListPayments struct {
	Payment []PaymentEntry `json:"payments"`
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
	Type           string       `json:"type"`
	Source         string       `json:"source"`
	SourceFeeMoney *AmountMoney `json:"source_fee_money,omitempty"`
	SourceId       string       `json:"source_id,omitempty"`
}

type CreatePayment struct {
	IdempotencyKey                 string          `json:"idempotency_key"`
	SourceId                       string          `json:"source_id"`
	AmountMoney                    *AmountMoney    `json:"amount_money,omitempty"`
	AppFeeMoney                    *AmountMoney    `json:"app_fee_money,omitempty"`
	BillingAddress                 *BillingAddress `json:"billing_address,omitempty"`
	CashDetails                    *CashDetails    `json:"cash_details,omitempty"`
	TipMoney                       *AmountMoney    `json:"tip_money,omitempty"`
	ShippingAddress                *BillingAddress `json:"shipping_address,omitempty"`
	AcceptPartialAuthorization     bool            `json:"accept_partial_authorization,omitempty"`
	Autocomplete                   bool            `json:"autocomplete,omitempty"`
	BuyerEmailAddress              string          `json:"buyer_email_address,omitempty"`
	CustomerDetails                *CustomerDetail `json:"customer_details,omitempty"`
	CustomerId                     string          `json:"customer_id,omitempty"`
	DelayAction                    string          `json:"delay_action,omitempty"`
	DelayDuration                  string          `json:"delay_duration,omitempty"`
	ExternalDetails                *ExternalDetail `json:"external_details,omitempty"`
	LocationId                     string          `json:"location_id,omitempty"`
	Note                           string          `json:"note,omitempty"`
	OrderId                        string          `json:"order_id,omitempty"`
	ReferenceId                    string          `json:"reference_id,omitempty"`
	StatementDescriptionIdentifier string          `json:"statement_description_identifier,omitempty"`
	TeamMemberId                   string          `json:"team_member_id,omitempty"`
	VerificationToken              string          `json:"verification_token,omitempty"`
}

type UpdatePayment struct {
	Payment        *UpdatePaymentDetails `json:"payment,omitempty"`
	IdempotencyKey string                `json:"idempotency_key"`
}

type UpdatePaymentDetails struct {
	AmountMoney   *AmountMoney `json:"amount_money,omitempty"`
	AppFeeMoney   *AmountMoney `json:"app_fee_money,omitempty"`
	ApprovedMoney *AmountMoney `json:"approved_money,omitempty"`
	CashDetails   *CashDetails `json:"cash_details,omitempty"`
	DelayAction   string       `json:"delay_action,omitempty"`
	TipMoney      *AmountMoney `json:"tip_money,omitempty"`
	VersionToken  string       `json:"version_token,omitempty"`
}

type CashDetails struct {
	BuyerSuppliedMoney *AmountMoney `json:"buyer_supplied_money"`
	ChangeBackMoney    *AmountMoney `json:"change_back_money,omitempty"`
}

type CancelPaymentByIdempotencyKey struct {
	IdempotencyKey string `json:"idempotency_key"`
}

type completePayment struct {
	VersionToken string `json:"version_token"`
}

// ListPayment returns a list of payments taken by the account making the request.
func (s *PaymentServiceOp) ListPayment(ctx context.Context, options *ListOptions) ([]ListPayments, *Response, error) {
	path := PaymentBasePath
	path, err := addOptions(path, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var payments []ListPayments
	resp, err := s.client.Do(ctx, req, &payments)
	if err != nil {
		return nil, resp, err
	}

	return payments, resp, nil
}

// CreatePayment creates a payment.
func (s *PaymentServiceOp) CreatePayment(ctx context.Context, payment *CreatePayment) (*Payment, *Response, error) {
	req, err := s.client.NewRequest(ctx, http.MethodPost, PaymentBasePath, payment)
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// CancelByIdempotencyKey cancels a payment by idempotency key.
func (s *PaymentServiceOp) CancelByIdempotencyKey(ctx context.Context, idempotencyKey string) (*Payment, *Response, error) {
	path := PaymentBasePath + "/cancel"
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, CancelPaymentByIdempotencyKey{IdempotencyKey: idempotencyKey})
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// GetPayment returns a payment by ID.
func (s *PaymentServiceOp) GetPayment(ctx context.Context, paymentId string) (*Payment, *Response, error) {
	path := PaymentBasePath + "/" + paymentId
	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// UpdatePayment updates a payment.
func (s *PaymentServiceOp) UpdatePayment(ctx context.Context, paymentId string, payment *Payment) (*Payment, *Response, error) {
	path := PaymentBasePath + "/" + paymentId
	req, err := s.client.NewRequest(ctx, http.MethodPut, path, payment)
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// CancelPayment cancels a payment.
func (s *PaymentServiceOp) CancelPayment(ctx context.Context, paymentId string) (*Payment, *Response, error) {
	path := PaymentBasePath + "/" + paymentId + "/cancel"
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// CompletePayment completes a payment.
func (s *PaymentServiceOp) CompletePayment(ctx context.Context, paymentId, versionToken string) (*Payment, *Response, error) {
	path := PaymentBasePath + "/" + paymentId + "/complete"

	var cpp *completePayment
	if versionToken != "" {
		cpp.VersionToken = versionToken
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, cpp)
	if err != nil {
		return nil, nil, err
	}

	root := new(Payment)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}
