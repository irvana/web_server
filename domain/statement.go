package domain

import "context"

type StatementResponse struct {
	StatementID  string `json:"statementId,omitempty"`
	RefNo        string `json:"refNo,omitempty"`
	DebitAmount  string `json:"debitAmount,omitempty"`
	CreditAmount string `json:"creditAmount,omitempty"`
	Balance      string `json:"balance,omitempty"`
	PostDate     string `json:"postDate,omitempty"`
	ValueDate    string `json:"valueDate,omitempty"`
}

type StatementRepository interface {
	List(ctx context.Context, req *BaseRequest) ([]StatementResponse, error)
}
