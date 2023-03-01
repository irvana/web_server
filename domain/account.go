package domain

import "context"

type (
	AccountResponse struct {
		List   []Details `json:"list"`
		Status string    `json:"status"`
	}

	Details struct {
		Acc     string `json:"acc"`
		CcyID   string `json:"ccyId"`
		Balance string `json:"balance"`
	}
)

type (
	AccountUsecase interface {
		GetList(ctx context.Context, req BaseRequest) (AccountResponse, error)
		Open(ctx context.Context, req BaseRequest) (AccountResponse, error)
	}

	AccountRepository interface {
		GetList(ctx context.Context, req BaseRequest) (AccountResponse, error)
		Open(ctx context.Context, req BaseRequest) (AccountResponse, error)
	}
)
