package errors

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	ofErrors "bitbucket.org/ayopop/of-errors/v2"
)

func TestInternalErr(t *testing.T) {
	internalError := ofErrors.NewErrorFormatter(ofErrors.AyoErrorServerTimeout0202)
	badRequestError := ofErrors.NewErrorFormatter(ofErrors.AyoErrorBadRequest0204)
	preConditionFailed := ofErrors.NewErrorFormatter(ofErrors.AyoErrorBadRequest0204)
	wrongRequestContent := ofErrors.NewErrorFormatter(ofErrors.BadContentType0301)
	badRequestHeaderAcceptNotAllow := ofErrors.NewErrorFormatter(ofErrors.BadHeader0302)
	preconditionFailedReferenceNumber := ofErrors.NewErrorFormatter(ofErrors.ReferenceNumberIsInvalid0311)
	invalidTransactionId := ofErrors.NewErrorFormatter(ofErrors.TransactionIdParameterIsInvalid0310)
	preconditionFailedEmail := ofErrors.NewErrorFormatter(ofErrors.EmailAddressIsInvalid0309)
	preconditionFailedPhoneNumber := ofErrors.NewErrorFormatter(ofErrors.BFFPhoneNumberInvalid0601)
	preconditionFailedMerchantCode := ofErrors.NewErrorFormatter(ofErrors.MerchantCodeMustHaveSixCharacters0308)
	preconditionFailedCardNumber := ofErrors.NewErrorFormatter(ofErrors.CardNumberIsInvalid0107)
	preconditionFailedCustomerID := ofErrors.NewErrorFormatter(ofErrors.CustomerIdIsInvalid0305)
	preconditionFailedOTP := ofErrors.NewErrorFormatter(ofErrors.OtpInvalid0011)
	preconditionFailedMerchantCodeInactiveAndNotInDatabase := ofErrors.NewErrorFormatter(ofErrors.MerchantInactive0005)
	duplicateCorrelationId := ofErrors.NewErrorFormatter(ofErrors.XCorrelationIdAlreadyUsed0316)
	ayoErrorCustomerAndMerchantMismatching := ofErrors.NewErrorFormatter(ofErrors.AyoErrorCustomerAndMerchantMismatching0203)

	type args struct {
		ctx        context.Context
		errContext string
		errMsg     string
	}

	tests := []struct {
		name string
		args args
		want struct {
			Code    int
			Message string
			Errors  ErrorItem
		}
	}{

		{
			name: "InternalErr",
			args: args{
				ctx:        context.TODO(),
				errContext: AyoErrorServerTimeout0202,
				errMsg:     ServiceUnavailableErrMsg,
			},
			want: WantError{
				Code: ProductCodeInternal, Message: ServiceUnavailableErrMsg, Errors: ErrorItem{
					Code: internalError.GetCode(),
				},
			},
		},
		{
			name: "BadRequest",
			args: args{
				ctx:        context.TODO(),
				errContext: AyoErrorBadRequest0204,
				errMsg:     BadRequestErrMsg,
			},
			want: WantError{
				Code: ProductCodeBadRequest, Message: BadRequestErrMsg, Errors: ErrorItem{
					Code: badRequestError.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailed",
			args: args{
				ctx:        context.TODO(),
				errContext: AyoErrorBadRequest0204,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preConditionFailed.GetCode(),
				},
			},
		},

		{
			name: "BadRequestContentTypeNotAllow",
			args: args{
				ctx:        context.TODO(),
				errContext: BadContentType0301,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: wrongRequestContent.GetCode(),
				},
			},
		},
		{
			name: "BadRequestHeaderAcceptNotAllow",
			args: args{
				ctx:        context.TODO(),
				errContext: BadHeader0302,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: badRequestHeaderAcceptNotAllow.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedTransactionID",
			args: args{
				ctx:        context.TODO(),
				errContext: TransactionIdParameterIsInvalid0310,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: invalidTransactionId.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedReferenceNumber",
			args: args{
				ctx:        context.TODO(),
				errContext: ReferenceNumberIsInvalid0311,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedReferenceNumber.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedCardNumber",
			args: args{
				ctx:        context.TODO(),
				errContext: CardNumberIsInvalid0107,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedCardNumber.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedCustomerID",
			args: args{
				ctx:        context.TODO(),
				errContext: CustomerIdIsInvalid0305,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedCustomerID.GetCode(),
				},
			},
		},

		{
			name: "PreconditionFailedOTP",
			args: args{
				ctx:        context.TODO(),
				errContext: OtpInvalid0011,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedOTP.GetCode(),
				},
			},
		},

		{
			name: "PreconditionFailedMerchantCodeInactiveAndNotInDatabase",
			args: args{
				ctx:        context.TODO(),
				errContext: MerchantInactive0005,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedMerchantCodeInactiveAndNotInDatabase.GetCode(),
				},
			},
		},
		{
			name: "DuplicateCorrelationId",
			args: args{
				ctx:        context.TODO(),
				errContext: XCorrelationIdAlreadyUsed0316,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: duplicateCorrelationId.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedEmail",
			args: args{
				ctx:        context.TODO(),
				errContext: EmailAddressIsInvalid0309,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedEmail.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedPhoneNumber",
			args: args{
				ctx:        context.TODO(),
				errContext: BFFPhoneNumberInvalid0601,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedPhoneNumber.GetCode(),
				},
			},
		},
		{
			name: "PreconditionFailedMerchantCode",
			args: args{
				ctx:        context.TODO(),
				errContext: MerchantCodeMustHaveSixCharacters0308,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: preconditionFailedMerchantCode.GetCode(),
				},
			},
		},
		{
			name: "AyoErrorCustomerAndMerchantMismatching",
			args: args{
				ctx:        context.TODO(),
				errContext: AyoErrorCustomerAndMerchantMismatching0203,
				errMsg:     PreconditionErrMsg,
			},
			want: WantError{
				Code: ProductCodePreconditionFailed, Message: PreconditionErrMsg, Errors: ErrorItem{
					Code: ayoErrorCustomerAndMerchantMismatching.GetCode(),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				err := GetError(tt.args.ctx, tt.args.errContext, tt.args.errMsg)

				if got, want := err.Message, tt.want.Message; got != want {
					t.Fatalf("err.Error() = %q; want %q", got, want)
				}
				if got, want := err.Code, tt.want.Code; got != want {
					t.Fatalf("err.Code = %q; want %q", got, want)
				}
				if got, want := err.Errors[0].Code, tt.want.Errors.Code; got != want {
					t.Fatalf("len(err.Errors) = %q; want %q", got, want)
				}
			},
		)
	}
}

func TestCustomError_Error(t *testing.T) {
	type fields struct {
		Code            int
		Message         string
		ResponseTime    string
		TransactionId   string
		ReferenceNumber string
		Errors          []ErrorItem
		HTTPCode        int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				Code:            200,
				Message:         "test",
				ResponseTime:    "",
				TransactionId:   "",
				ReferenceNumber: "",
				Errors: []ErrorItem{
					{
						Code:    "200",
						Message: "test",
						Details: "test",
					},
				},
				HTTPCode: 0,
			},
			want: "CustomError code = 200 desc - test errors = [{200 test test}]",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				o := CustomError{
					Code:            tt.fields.Code,
					Message:         tt.fields.Message,
					ResponseTime:    tt.fields.ResponseTime,
					TransactionId:   tt.fields.TransactionId,
					ReferenceNumber: tt.fields.ReferenceNumber,
					Errors:          tt.fields.Errors,
					HTTPCode:        tt.fields.HTTPCode,
				}
				if got := o.Error(); got != tt.want {
					t.Errorf("Error() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestCustomError_GetHTTPCode(t *testing.T) {
	type fields struct {
		Code            int
		Message         string
		ResponseTime    string
		TransactionId   string
		ReferenceNumber string
		Errors          []ErrorItem
		HTTPCode        int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "Success", fields: fields{
				Code:            503,
				Message:         "",
				ResponseTime:    "",
				TransactionId:   "",
				ReferenceNumber: "",
				Errors:          nil,
				HTTPCode:        503,
			}, want: 503,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				o := CustomError{
					Code:            tt.fields.Code,
					Message:         tt.fields.Message,
					ResponseTime:    tt.fields.ResponseTime,
					TransactionId:   tt.fields.TransactionId,
					ReferenceNumber: tt.fields.ReferenceNumber,
					Errors:          tt.fields.Errors,
					HTTPCode:        tt.fields.HTTPCode,
				}
				if got := o.GetHTTPCode(); got != tt.want {
					t.Errorf("GetHTTPCode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestGetError(t *testing.T) {
	type args struct {
		ctx        context.Context
		errContext string
		errMsg     string
	}
	tests := []struct {
		name string
		args args
		want *CustomError
	}{
		{
			name: "Success",
			args: args{ctx: context.Background(), errContext: "precondition.failed", errMsg: PreconditionErrMsg},
			want: GetError(context.Background(), "precondition.failed", PreconditionErrMsg),
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := GetError(tt.args.ctx, tt.args.errContext, tt.args.errMsg); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("GetError() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_getHttpCode(t *testing.T) {
	type args struct {
		errContext string
	}
	tests := []struct {
		name string
		args args
		want ErrorVal
	}{
		{
			name: "Success",
			args: struct{ errContext string }{errContext: "precondition.failed"},
			want: ErrorVal{
				Code:     ProductCodePreconditionFailed,
				Message:  PreconditionErrMsg,
				HttpCode: http.StatusPreconditionFailed,
			},
		},
		{
			name: "Invalid error-code",
			args: struct{ errContext string }{errContext: "precondition1.failed"},
			want: ErrorVal{
				Code:     ProductCodeInternal,
				Message:  ServiceUnavailableErrMsg,
				HttpCode: http.StatusServiceUnavailable,
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := getHttpCode(tt.args.errContext); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("getHttpCode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestCustomError_GetErrCode(t *testing.T) {
	type fields struct {
		Code            int
		Message         string
		ResponseTime    string
		TransactionId   string
		ReferenceNumber string
		Errors          []ErrorItem
		HTTPCode        int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "OK",
			fields: fields{
				Code:            0,
				Message:         "",
				ResponseTime:    "",
				TransactionId:   "",
				ReferenceNumber: "",
				Errors: []ErrorItem{
					{
						Code:    "200",
						Message: "test",
						Details: "test",
					},
				},
				HTTPCode: 0,
			},
			want: "200",
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				o := CustomError{
					Code:            tt.fields.Code,
					Message:         tt.fields.Message,
					ResponseTime:    tt.fields.ResponseTime,
					TransactionId:   tt.fields.TransactionId,
					ReferenceNumber: tt.fields.ReferenceNumber,
					Errors:          tt.fields.Errors,
					HTTPCode:        tt.fields.HTTPCode,
				}
				if got := o.GetErrCode(); got != tt.want {
					t.Errorf("GetErrCode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
