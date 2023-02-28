package domain

type OrderResponse struct {
	Cif         string `json:"cif"`
	Pair        string `json:"pair"`
	BaseQuote   string `json:"baseQuote"`
	BidAsk      string `json:"bidAsk"`
	BaseAmt     string `json:"baseAmt"`
	CtrAmt      string `json:"ctrAmt"`
	Rate        string `json:"rate"`
	DebitAcc    string `json:"debitAcc"`
	CreditAcc   string `json:"creditAcc"`
	TakeProfit  string `json:"takeProfit"`
	StopLoss    string `json:"stopLoss"`
	OrderID     string `json:"orderId"`
	OrderDate   string `json:"orderDate"`
	OrderType   string `json:"orderType"`
	OrderStatus string `json:"orderStatus"`
	ExpiryDate  string `json:"expiryDate"`
	Instrument  string `json:"instrument,omitempty"`
	Status      string `json:"status,omitempty"`
	Reason      string `json:"reason,omitempty"`
}
