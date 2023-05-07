package domain

import "context"

type OrderResponse struct {
	Cif         string `json:"cif,omitempty"`
	Pair        string `json:"pair,omitempty"`
	BaseQuote   string `json:"baseQuote,omitempty"`
	BidAsk      string `json:"bidAsk,omitempty"`
	BaseAmt     string `json:"baseAmt,omitempty"`
	CtrAmt      string `json:"ctrAmt,omitempty"`
	Rate        string `json:"rate,omitempty"`
	DebitAcc    string `json:"debitAcc,omitempty"`
	CreditAcc   string `json:"creditAcc,omitempty"`
	TakeProfit  string `json:"takeProfit,omitempty"`
	StopLoss    string `json:"stopLoss,omitempty"`
	OrderID     string `json:"orderId,omitempty"`
	OrderDate   string `json:"orderDate,omitempty"`
	OrderType   string `json:"orderType,omitempty"`
	OrderStatus string `json:"orderStatus,omitempty"`
	ExpiryDate  string `json:"expiryDate,omitempty"`
	Instrument  string `json:"instrument,omitempty,omitempty"`
	Status      string `json:"status,omitempty,omitempty"`
	Reason      string `json:"reason,omitempty,omitempty"`
}

type OrderUsecase interface {
	GetStatus(ctx context.Context, req *BaseRequest) ([]OrderResponse, error)
	Cancel(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	Amend(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	GetDetail(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	Create(ctx context.Context, req *BaseRequest) (*OrderResponse, error)

	GetStatementList(ctx context.Context, req *BaseRequest) ([]StatementResponse, error)

	GetTransactionDetail(ctx context.Context, req *BaseRequest) (*TransactionResponse, error)
	Deal(ctx context.Context, req *BaseRequest) (*TransactionResponse, error)
}
type OrderRepository interface {
	GetStatus(ctx context.Context, req *BaseRequest) ([]OrderResponse, error)
	Cancel(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	Amend(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	GetDetail(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
	Create(ctx context.Context, req *BaseRequest) (*OrderResponse, error)
}
