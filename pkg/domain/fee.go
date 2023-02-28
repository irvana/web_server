package domain

type Currency struct {
	ID      string `json:"id"`
	Label   string `json:"label"`
	Name    string `json:"name"`
	Decimal string `json:"decimal"`
}

type Pair struct {
	ID         string `json:"id"`
	Label      string `json:"label"`
	BaseCcyID  string `json:"baseCcyId"`
	QuoteCcyID string `json:"quoteCcyId"`
	Decimal    string `json:"decimal"`
	Tradeable  string `json:"tradeable"`
}

type News struct {
	ID        string `json:"id"`
	NewsID    string `json:"newsId"`
	PairID    string `json:"pairId"`
	Stamp     string `json:"stamp"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}
