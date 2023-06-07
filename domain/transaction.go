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

type TransactionRequest struct {
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

type TransactionRepository interface {
	GetDetail(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error)
	Deal(ctx context.Context, req *TransactionRequest) (*TransactionResponse, error)
}
