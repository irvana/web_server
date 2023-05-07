package domain

import "context"

type (
	AccountResponse struct {
		Acc         string `json:"acc,omitempty"`
		Name        string `json:"name,omitempty"`
		AccountType string `json:"accountType,omitempty"`
		ProductType string `json:"productType,omitempty"`
		Bank        string `json:"bank,omitempty"`
		BankBranch  string `json:"bankBranch,omitempty"`
		BankCode    string `json:"bankCode,omitempty"`
		Ccy         string `json:"ccy,omitempty"`
		Balance     string `json:"balance,omitempty"`
	}

	AccountOpenResponse struct {
		Status  string `json:"status,omitempty"`
		Message string `json:"message,omitempty"`
	}
)

type (
	AccountUsecase interface {
		GetList(ctx context.Context, req *BaseRequest) ([]AccountResponse, error)
		Open(ctx context.Context, req *BaseRequest) (*AccountOpenResponse, error)
	}

	AccountRepository interface {
		GetList(ctx context.Context, req *BaseRequest) ([]AccountResponse, error)
		Open(ctx context.Context, req *BaseRequest) (*AccountOpenResponse, error)
	}
)
