package domain

import "context"

type TransactionResponse struct {
	Cif        string `json:"cif"`
	TicketID   string `json:"ticketId"`
	DealDate   string `json:"dealDate"`
	ValueDate  string `json:"valueDate"`
	Pair       string `json:"pair"`
	BaseQuote  string `json:"baseQuote"`
	BidAsk     string `json:"bidAsk"`
	BaseAmt    string `json:"baseAmt"`
	CtrAmt     string `json:"ctrAmt"`
	Rate       string `json:"rate"`
	DebitAcc   string `json:"debitAcc"`
	CreditAcc  string `json:"creditAcc"`
	TakeProfit string `json:"takeProfit"`
	StopLoss   string `json:"stopLoss"`
	Settle     string `json:"settle"`
	DealStatus string `json:"dealStatus"`
	Status     string `json:"status,omitempty"`
	Reason     string `json:"reason,omitempty"`
}
type TransactionUsecase interface {
	GetDetail(ctx context.Context, req BaseRequest) (TransactionResponse, error)
	Deal(ctx context.Context, req BaseRequest) (TransactionResponse, error)
}
type TransactionRepository interface {
	GetDetail(ctx context.Context, req BaseRequest) (TransactionResponse, error)
	Deal(ctx context.Context, req BaseRequest) (TransactionResponse, error)
}
