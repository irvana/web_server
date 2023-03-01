package domain

import "context"

type StatementResponse []StatementDetails

type StatementDetails struct {
	Cif       string `json:"cif"`
	TicketID  string `json:"ticketId"`
	DealDate  string `json:"dealDate"`
	ValueDate string `json:"valueDate"`
	Pair      string `json:"pair"`
	BaseQuote string `json:"baseQuote"`
	BidAsk    string `json:"bidAsk"`
	BaseAmt   string `json:"baseAmt"`
	CtrAmt    string `json:"ctrAmt"`
	Rate      string `json:"rate"`
	DebitAcc  string `json:"debitAcc"`
	CreditAcc string `json:"creditAcc"`
}

type StatementUsecase interface {
	List(ctx context.Context, req BaseRequest) (StatementResponse, error)
}
type StatementRepository interface {
	List(ctx context.Context, req BaseRequest) (StatementResponse, error)
}
