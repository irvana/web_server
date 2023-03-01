package handler

import "web_server/domain"

type onboardingHandler struct {
}

const (
	ACCOUNT_LIST = "/account/list"
	ACCOUNT_OPEN = "/account/open"

	STATEMENT_LIST = "/statement/list"

	ORDER_STATUS = "/order/status"
	ORDER_CANCEL = "/order/cancel"
	ORDER_AMEND  = "/order/amend"
	ORDER_DETAIL = "/order/detail"
	ORDER_CREATE = "/order/create"

	TRANSACTION_DETAIL = "/transaction/detail"
	TRANSACTION_DEAL   = "/transaction/deal"
	REF_ALL            = "/ref/all"
	REF_CCY            = "/ref/ccy"
	REF_PAIR           = "/ref/pair"
	REF_NEWS           = "/ref/news"

	WEBSOCKET_SERVER = "/wss/rates"
)

func NewOnboardingHandler(obUsecase domain.OrderUsecase) {

}
