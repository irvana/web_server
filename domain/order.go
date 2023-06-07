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
	Instrument  string `json:"instrument,omitempty"`
	Status      string `json:"status,omitempty"`
	Reason      string `json:"reason,omitempty"`
}

type OrderRequest struct {
	Cif        string `json:"cif"`
	OrderID    string `json:"orderId,omitempty"`
	Name       string `json:"name,omitempty"`
	PairID     string `json:"pairId,omitempty"`
	BaseQuote  string `json:"baseQuote,omitempty"`
	BidAsk     string `json:"bidAsk,omitempty"`
	Amt        string `json:"amt,omitempty"`
	Rate       string `json:"rate,omitempty"`
	DebitAcc   string `json:"debitAcc,omitempty"`
	CreditAcc  string `json:"creditAcc,omitempty"`
	TakeProfit string `json:"takeProfit,omitempty"`
	StopLoss   string `json:"stopLoss,omitempty"`
	ExpiryDate string `json:"expiryDate,omitempty"`
	UUIDStamp  string `json:"uuidStamp,omitempty"`
}

type StatementRequest struct {
	Cif         string `json:"cif"`
	StatementID string `json:"statementId,omitempty"`
	RefNo       string `json:"refNo,omitempty"`
	Name        string `json:"name,omitempty"`
	PairID      string `json:"pairId,omitempty"`
	BaseQuote   string `json:"baseQuote,omitempty"`
	BidAsk      string `json:"bidAsk,omitempty"`
	Amt         string `json:"amt,omitempty"`
	Rate        string `json:"rate,omitempty"`
	DebitAcc    string `json:"debitAcc,omitempty"`
	CreditAcc   string `json:"creditAcc,omitempty"`
	TakeProfit  string `json:"takeProfit,omitempty"`
	StopLoss    string `json:"stopLoss,omitempty"`
	UUIDStamp   string `json:"uuidStamp,omitempty"`
}

type OrderUsecase interface {
	GetStatus(ctx context.Context, req *OrderRequest) ([]OrderResponse, error)
	Cancel(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	Amend(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	GetDetail(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	Create(ctx context.Context, req *OrderRequest) (*OrderResponse, error)

	GetStatementList(ctx context.Context, req *StatementRequest) ([]StatementResponse, error)

	GetTransactionDetail(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error)
	Deal(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error)
}
type OrderRepository interface {
	GetStatus(ctx context.Context, req *OrderRequest) ([]OrderResponse, error)
	Cancel(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	Amend(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	GetDetail(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
	Create(ctx context.Context, req *OrderRequest) (*OrderResponse, error)
}
