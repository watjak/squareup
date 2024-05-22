package squareup

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	listPaymentsResponse = `
{
  "payments": [
    {
      "id": "lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
      "created_at": "2024-05-20T12:07:07.941Z",
      "updated_at": "2024-05-20T12:07:08.091Z",
      "amount_money": {
        "amount": 29,
        "currency": "USD"
      },
      "status": "COMPLETED",
      "source_type": "EXTERNAL",
      "location_id": "LMPCZVC1C3FHM",
      "order_id": "oYnxgIOneziRXSWQcwHEv39pod4F",
      "note": "my-note",
      "total_money": {
        "amount": 29,
        "currency": "USD"
      },
      "capabilities": [
        "EDIT_AMOUNT_UP",
        "EDIT_AMOUNT_DOWN",
        "EDIT_TIP_AMOUNT_UP",
        "EDIT_TIP_AMOUNT_DOWN"
      ],
      "external_details": {
        "type": "CARD",
        "source": "Developer Control Panel"
      },
      "receipt_number": "lyQe",
      "receipt_url": "https://squareupsandbox.com/receipt/preview/lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
      "application_details": {
        "square_product": "ECOMMERCE_API",
        "application_id": "sandbox-sq0idb-lky4CaPAWmDnHY3YtYxINg"
      },
      "version_token": "JvrByjSdgc1rllhW9yYXKSfvoLKUjm64Ga3ilf1tLVS6o"
    }
  ]
}`

	createPaymentResponseJSONBody = `
{
  "payment": {
    "id": "R2B3Z8WMVt3EAmzYWLZvz7Y69EbZY",
    "created_at": "2021-10-13T21:14:29.577Z",
    "updated_at": "2021-10-13T21:14:30.504Z",
    "amount_money": {
      "amount": 1000,
      "currency": "USD"
    },
    "app_fee_money": {
      "amount": 10,
      "currency": "USD"
    },
    "status": "COMPLETED",
    "delay_duration": "PT168H",
    "source_type": "CARD",
    "card_details": {
      "status": "CAPTURED",
      "card": {
        "card_brand": "VISA",
        "last_4": "1111",
        "exp_month": 11,
        "exp_year": 2022,
        "fingerprint": "sq-1-Hxim77tbdcbGejOejnoAklBVJed2YFLTmirfl8Q5XZzObTc8qY_U8RkwzoNL8dCEcQ",
        "card_type": "DEBIT",
        "prepaid_type": "NOT_PREPAID",
        "bin": "411111"
      },
      "entry_method": "ON_FILE",
      "cvv_status": "CVV_ACCEPTED",
      "avs_status": "AVS_ACCEPTED",
      "auth_result_code": "vNEn2f",
      "statement_description": "SQ *EXAMPLE TEST GOSQ.C",
      "card_payment_timeline": {
        "authorized_at": "2021-10-13T21:14:29.732Z",
        "captured_at": "2021-10-13T21:14:30.504Z"
      }
    },
    "location_id": "L88917AVBK2S5",
    "order_id": "pRsjRTgFWATl7so6DxdKBJa7ssbZY",
    "reference_id": "123456",
    "risk_evaluation": {
      "created_at": "2021-10-13T21:14:30.423Z",
      "risk_level": "NORMAL"
    },
    "note": "Brief Description",
    "customer_id": "W92WH6P11H4Z77CTET0RNTGFW8",
    "total_money": {
      "amount": 1000,
      "currency": "USD"
    },
    "approved_money": {
      "amount": 1000,
      "currency": "USD"
    },
    "receipt_number": "R2B3",
    "receipt_url": "https://squareup.com/receipt/preview/EXAMPLE_RECEIPT_ID",
    "delay_action": "CANCEL",
    "delayed_until": "2021-10-20T21:14:29.577Z",
    "application_details": {
      "square_product": "ECOMMERCE_API",
      "application_id": "sq0ids-TcgftTEtKxJTRF1lCFJ9TA"
    },
    "version_token": "TPtNEOBOa6Qq6E3C3IjckSVOM6b3hMbfhjvTxHBQUsB6o"
  }
}`

	paymentResponse = `
{
  "payments": 
    {
      "id": "lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
      "created_at": "2024-05-20T12:07:07.941Z",
      "updated_at": "2024-05-20T12:07:08.091Z",
      "amount_money": {
        "amount": 29,
        "currency": "USD"
      },
      "status": "COMPLETED",
      "source_type": "EXTERNAL",
      "location_id": "LMPCZVC1C3FHM",
      "order_id": "oYnxgIOneziRXSWQcwHEv39pod4F",
      "note": "my-note",
      "total_money": {
        "amount": 29,
        "currency": "USD"
      },
      "capabilities": [
        "EDIT_AMOUNT_UP",
        "EDIT_AMOUNT_DOWN",
        "EDIT_TIP_AMOUNT_UP",
        "EDIT_TIP_AMOUNT_DOWN"
      ],
      "external_details": {
        "type": "CARD",
        "source": "Developer Control Panel"
      },
      "receipt_number": "lyQe",
      "receipt_url": "https://squareupsandbox.com/receipt/preview/lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
      "application_details": {
        "square_product": "ECOMMERCE_API",
        "application_id": "sandbox-sq0idb-lky4CaPAWmDnHY3YtYxINg"
      },
      "version_token": "JvrByjSdgc1rllhW9yYXKSfvoLKUjm64Ga3ilf1tLVS6o"
    }
}`
)

func TestPaymentServiceOp_CreatePayment(t *testing.T) {
	setup()
	defer teardown()

	expectedPaymentCreateRequest := &CreatePayment{
		IdempotencyKey: "7b0f3ec5-086a-4871-8f13-3c81b3875218",
		AmountMoney: &AmountMoney{
			Amount:   1000,
			Currency: "USD",
		},
		SourceId:     "ccof:GaJGNaZa8x4OgDJn4GB",
		Autocomplete: true,
		CustomerId:   "W92WH6P11H4Z77CTET0RNTGFW8",
		LocationId:   "L88917AVBK2S5",
		ReferenceId:  "123456",
		Note:         "Brief description",
		AppFeeMoney: &AmountMoney{
			Amount:   10,
			Currency: "USD",
		},
	}

	mux.HandleFunc("/v2/payments", func(w http.ResponseWriter, r *http.Request) {
		v := new(CreatePayment)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}

		testMethod(t, r, http.MethodPost)
		if !reflect.DeepEqual(v, expectedPaymentCreateRequest) {
			t.Errorf("Request body = %+v, expected %+v", v, expectedPaymentCreateRequest)
		}

		fmt.Fprint(w, createPaymentResponseJSONBody)
	})

	actualPaymentResponse, _, err := client.Payment.CreatePayment(ctx, expectedPaymentCreateRequest)
	if err != nil {
		t.Errorf("Payment.CreatePayment returned error: %v", err)
	}

	expectedPaymentResponse := &Payment{
		Payment: &PaymentEntry{
			Id:        "R2B3Z8WMVt3EAmzYWLZvz7Y69EbZY",
			CreatedAt: time.Date(2021, 10, 13, 21, 14, 29, 577000000, time.UTC),
			UpdatedAt: time.Date(2021, 10, 13, 21, 14, 30, 504000000, time.UTC),
			AmountMoney: &AmountMoney{
				Amount:   1000,
				Currency: "USD",
			},
			AppFeeMoney: &AmountMoney{
				Amount:   10,
				Currency: "USD",
			},
			Status:        "COMPLETED",
			DelayDuration: "PT168H",
			SourceType:    "CARD",
			CardDetails: &CardDetails{
				Status: "CAPTURED",
				Card: &Card{
					CardBrand:   "VISA",
					Last4:       "1111",
					ExpMonth:    11,
					ExpYear:     2022,
					Fingerprint: "sq-1-Hxim77tbdcbGejOejnoAklBVJed2YFLTmirfl8Q5XZzObTc8qY_U8RkwzoNL8dCEcQ",
					CardType:    "DEBIT",
					PrepaidType: "NOT_PREPAID",
					Bin:         "411111",
				},
				EntryMethod:          "ON_FILE",
				CvvStatus:            "CVV_ACCEPTED",
				AvsStatus:            "AVS_ACCEPTED",
				AuthResultCode:       "vNEn2f",
				StatementDescription: "SQ *EXAMPLE TEST GOSQ.C",
				CardPaymentTimeline: &CardPaymentTimeline{
					AuthorizedAt: time.Date(2021, 10, 13, 21, 14, 29, 732000000, time.UTC),
					CapturedAt:   time.Date(2021, 10, 13, 21, 14, 30, 504000000, time.UTC),
				},
			},
			LocationId:  "L88917AVBK2S5",
			OrderId:     "pRsjRTgFWATl7so6DxdKBJa7ssbZY",
			ReferenceId: "123456",
			RiskEvaluation: &RiskEvaluation{
				CreatedAt: time.Date(2021, 10, 13, 21, 14, 30, 423000000, time.UTC),
				RiskLevel: "NORMAL",
			},
			Note:       "Brief Description",
			CustomerId: "W92WH6P11H4Z77CTET0RNTGFW8",
			TotalMoney: &AmountMoney{
				Amount:   1000,
				Currency: "USD",
			},
			ApprovedMoney:      &AmountMoney{Amount: 1000, Currency: "USD"},
			ReceiptNumber:      "R2B3",
			ReceiptUrl:         "https://squareup.com/receipt/preview/EXAMPLE_RECEIPT_ID",
			DelayAction:        "CANCEL",
			DelayedUntil:       time.Date(2021, 10, 20, 21, 14, 29, 577000000, time.UTC),
			ApplicationDetails: &ApplicationDetails{SquareProduct: "ECOMMERCE_API", ApplicationId: "sq0ids-TcgftTEtKxJTRF1lCFJ9TA"},
			VersionToken:       "TPtNEOBOa6Qq6E3C3IjckSVOM6b3hMbfhjvTxHBQUsB6o",
		},
	}

	if !reflect.DeepEqual(actualPaymentResponse, expectedPaymentResponse) {
		t.Errorf("Payment.CreatePayment returned %+v, expected %+v", actualPaymentResponse, expectedPaymentResponse)
	}
}

func TestPaymentServiceOp_ListPayment(t *testing.T) {
	var paymentResponseCase1 = PaymentEntry{
		Id:        "lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
		CreatedAt: time.Date(2024, 5, 20, 12, 7, 7, 941000000, time.UTC),
		UpdatedAt: time.Date(2024, 5, 20, 12, 7, 8, 91000000, time.UTC),
		AmountMoney: &AmountMoney{
			Amount:   29,
			Currency: "USD",
		},
		Status:     "COMPLETED",
		SourceType: "EXTERNAL",
		LocationId: "LMPCZVC1C3FHM",
		OrderId:    "oYnxgIOneziRXSWQcwHEv39pod4F",
		Note:       "my-note",
		TotalMoney: &AmountMoney{
			Amount:   29,
			Currency: "USD",
		},
		Capabilities: []string{
			"EDIT_AMOUNT_UP",
			"EDIT_AMOUNT_DOWN",
			"EDIT_TIP_AMOUNT_UP",
			"EDIT_TIP_AMOUNT_DOWN",
		},
		ExternalDetails: &ExternalDetails{
			Type:   "CARD",
			Source: "Developer Control Panel",
		},
		ReceiptNumber: "lyQe",
		ReceiptUrl:    "https://squareupsandbox.com/receipt/preview/lyQeJ5EYqpWTbFgGXGKi3itW6PPZY",
		ApplicationDetails: &ApplicationDetails{
			SquareProduct: "ECOMMERCE_API",
			ApplicationId: "sandbox-sq0idb-lky4CaPAWmDnHY3YtYxINg",
		},
		VersionToken: "JvrByjSdgc1rllhW9yYXKSfvoLKUjm64Ga3ilf1tLVS6o",
	}

	var expectedPaymentResponse1 []PaymentEntry
	expectedPaymentResponse1 = append(expectedPaymentResponse1, paymentResponseCase1)

	type args struct {
		ctx     context.Context
		options *ListOptions
	}
	var tests = []struct {
		name    string
		args    args
		want    *ListPayments
		wantErr bool
	}{
		{
			name: "ListPayments",
			args: args{
				ctx:     ctx,
				options: &ListOptions{},
			},
			want: &ListPayments{
				Payment: expectedPaymentResponse1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			defer teardown()
			mux.HandleFunc("/v2/payments", func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, http.MethodGet)
				fmt.Fprint(w, listPaymentsResponse)
			})
			got, _, err := client.Payment.ListPayment(tt.args.ctx, tt.args.options)
			if (err != nil) != tt.wantErr {
				t.Errorf("PaymentServiceOp.ListPayments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PaymentServiceOp.ListPayments() = %v, want %v", got, tt.want)
			}
		})
	}
}
