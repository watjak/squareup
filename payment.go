package squareup

import (
	"context"
	"net/http"
	"path"
	"time"
)

const (
	PaymentBasePath = "v2/payments"
)

type PaymentService interface {
	ListPayment(ctx context.Context, options *ListOptions) (*ListPayments, *Response, error)
	CreatePayment(ctx context.Context, payment *CreatePayment) (*Payment, *Response, error)
	CancelByIdempotencyKey(ctx context.Context, id string) (*Payment, *Response, error)
	GetPayment(ctx context.Context, paymentId string) (*Payment, *Response, error)
	UpdatePayment(ctx context.Context, paymentId string, payment *Payment) (*Payment, *Response, error)
	CancelPayment(ctx context.Context, paymentId string) (*Payment, *Response, error)
	CompletePayment(ctx context.Context, paymentId, versionToken string) (*Payment, *Response, error)
}

var _ PaymentService = &PaymentServiceOp{}

// PaymentServiceOp handles communication with the payment related methods of the Square API.
type PaymentServiceOp struct {
	client *Client
}

// ListPayments represents a list of payments.
type ListPayments struct {
	Payment []PaymentEntry `json:"payments"`
}

// Payment represents a payment.
type Payment struct {
	Payment *PaymentEntry `json:"payment"`
}

// PaymentEntry represents a payment entry.
type PaymentEntry struct {
	Id                 string              `json:"id"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	AmountMoney        *AmountMoney        `json:"amount_money,omitempty"`
	AppFeeMoney        *AmountMoney        `json:"app_fee_money,omitempty"`
	Status             string              `json:"status,omitempty"`
	DelayDuration      string              `json:"delay_duration,omitempty"`
	SourceType         string              `json:"source_type,omitempty"`
	CardDetails        *CardDetails        `json:"card_details,omitempty"`
	LocationId         string              `json:"location_id,omitempty"`
	OrderId            string              `json:"order_id,omitempty"`
	ReferenceId        string              `json:"reference_id,omitempty"`
	RiskEvaluation     *RiskEvaluation     `json:"risk_evaluation,omitempty"`
	Note               string              `json:"note,omitempty"`
	CustomerId         string              `json:"customer_id,omitempty"`
	TotalMoney         *AmountMoney        `json:"total_money,omitempty"`
	ApprovedMoney      *AmountMoney        `json:"approved_money,omitempty"`
	Capabilities       []string            `json:"capabilities,omitempty"`
	ExternalDetails    *ExternalDetails    `json:"external_details,omitempty"`
	ReceiptNumber      string              `json:"receipt_number,omitempty"`
	ReceiptUrl         string              `json:"receipt_url,omitempty"`
	DelayAction        string              `json:"delay_action,omitempty"`
	DelayedUntil       time.Time           `json:"delayed_until,omitempty"`
	ApplicationDetails *ApplicationDetails `json:"application_details,omitempty"`
	VersionToken       string              `json:"version_token,omitempty"`
}

type CardDetails struct {
	Status               string               `json:"status,omitempty"`
	Card                 *Card                `json:"card,omitempty"`
	EntryMethod          string               `json:"entry_method,omitempty"`
	CvvStatus            string               `json:"cvv_status,omitempty"`
	AvsStatus            string               `json:"avs_status,omitempty"`
	AuthResultCode       string               `json:"auth_result_code,omitempty"`
	StatementDescription string               `json:"statement_description,omitempty"`
	CardPaymentTimeline  *CardPaymentTimeline `json:"card_payment_timeline,omitempty"`
}

type Card struct {
	CardBrand   string `json:"card_brand"`
	Last4       string `json:"last_4"`
	ExpMonth    int    `json:"exp_month"`
	ExpYear     int    `json:"exp_year"`
	Fingerprint string `json:"fingerprint"`
	CardType    string `json:"card_type"`
	PrepaidType string `json:"prepaid_type"`
	Bin         string `json:"bin"`
}

type CardPaymentTimeline struct {
	AuthorizedAt time.Time `json:"authorized_at"`
	CapturedAt   time.Time `json:"captured_at"`
}

type RiskEvaluation struct {
	CreatedAt time.Time `json:"created_at"`
	RiskLevel string    `json:"risk_level"`
}

// ApplicationDetails represents an application detail.
type ApplicationDetails struct {
	SquareProduct string `json:"square_product"`
	ApplicationId string `json:"application_id"`
}

// ExternalDetails represents an external detail.
type ExternalDetails struct {
	Type           string       `json:"type"`
	Source         string       `json:"source"`
	SourceFeeMoney *AmountMoney `json:"source_fee_money,omitempty"`
	SourceId       string       `json:"source_id,omitempty"`
}

// CreatePayment represents a payment to be created.
type CreatePayment struct {
	IdempotencyKey                 string           `json:"idempotency_key"`
	SourceId                       string           `json:"source_id"`
	AmountMoney                    *AmountMoney     `json:"amount_money,omitempty"`
	AppFeeMoney                    *AmountMoney     `json:"app_fee_money,omitempty"`
	BillingAddress                 *BillingAddress  `json:"billing_address,omitempty"`
	CashDetails                    *CashDetails     `json:"cash_details,omitempty"`
	TipMoney                       *AmountMoney     `json:"tip_money,omitempty"`
	ShippingAddress                *BillingAddress  `json:"shipping_address,omitempty"`
	AcceptPartialAuthorization     bool             `json:"accept_partial_authorization,omitempty"`
	Autocomplete                   bool             `json:"autocomplete,omitempty"`
	BuyerEmailAddress              string           `json:"buyer_email_address,omitempty"`
	CustomerDetails                *CustomerDetail  `json:"customer_details,omitempty"`
	CustomerId                     string           `json:"customer_id,omitempty"`
	DelayAction                    string           `json:"delay_action,omitempty"`
	DelayDuration                  string           `json:"delay_duration,omitempty"`
	ExternalDetails                *ExternalDetails `json:"external_details,omitempty"`
	LocationId                     string           `json:"location_id,omitempty"`
	Note                           string           `json:"note,omitempty"`
	OrderId                        string           `json:"order_id,omitempty"`
	ReferenceId                    string           `json:"reference_id,omitempty"`
	StatementDescriptionIdentifier string           `json:"statement_description_identifier,omitempty"`
	TeamMemberId                   string           `json:"team_member_id,omitempty"`
	VerificationToken              string           `json:"verification_token,omitempty"`
}

// UpdatePayment represents a payment to be updated.
type UpdatePayment struct {
	Payment        *UpdatePaymentDetails `json:"payment,omitempty"`
	IdempotencyKey string                `json:"idempotency_key"`
}

// UpdatePaymentDetails represents the details of a payment to be updated.
type UpdatePaymentDetails struct {
	AmountMoney   *AmountMoney `json:"amount_money,omitempty"`
	AppFeeMoney   *AmountMoney `json:"app_fee_money,omitempty"`
	ApprovedMoney *AmountMoney `json:"approved_money,omitempty"`
	CashDetails   *CashDetails `json:"cash_details,omitempty"`
	DelayAction   string       `json:"delay_action,omitempty"`
	TipMoney      *AmountMoney `json:"tip_money,omitempty"`
	VersionToken  string       `json:"version_token,omitempty"`
}

// CashDetails represents the cash details of a payment.
type CashDetails struct {
	BuyerSuppliedMoney *AmountMoney `json:"buyer_supplied_money"`
	ChangeBackMoney    *AmountMoney `json:"change_back_money,omitempty"`
}

// CancelPaymentByIdempotencyKey represents a payment to be canceled by idempotency key.
type CancelPaymentByIdempotencyKey struct {
	IdempotencyKey string `json:"idempotency_key"`
}

// completePayment represents a payment to be completed.
type completePayment struct {
	VersionToken string `json:"version_token"`
}

// ListPayment returns a list of payments taken by the account making the request.
func (s *PaymentServiceOp) ListPayment(ctx context.Context, options *ListOptions) (*ListPayments, *Response, error) {
	p := PaymentBasePath
	p, err := addOptions(p, options)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(ListPayments)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
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
	p := path.Join(PaymentBasePath, "cancel")
	req, err := s.client.NewRequest(ctx, http.MethodPost, p, CancelPaymentByIdempotencyKey{IdempotencyKey: idempotencyKey})
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
	p := path.Join(PaymentBasePath, paymentId)
	req, err := s.client.NewRequest(ctx, http.MethodGet, p, nil)
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
	p := path.Join(PaymentBasePath, paymentId)
	req, err := s.client.NewRequest(ctx, http.MethodPut, p, payment)
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
	p := path.Join(PaymentBasePath, paymentId, "cancel")
	req, err := s.client.NewRequest(ctx, http.MethodPost, p, nil)
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
	p := PaymentBasePath + "/" + paymentId + "/complete"

	var cpp *completePayment
	if versionToken != "" {
		cpp.VersionToken = versionToken
	}

	req, err := s.client.NewRequest(ctx, http.MethodPost, p, cpp)
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
